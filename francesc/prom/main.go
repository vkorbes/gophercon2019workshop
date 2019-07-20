package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	c := newCollector(stats)
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func stats() ([]CPUStat, error) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return nil, errors.Wrapf(err, "could not open /proc/stat")
	}
	defer f.Close()

	stats, err := ParseCPUStats(f)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

type collector struct {
	TimeUserHertzTotal *prometheus.Desc
	stats              func() ([]CPUStat, error)
}

func newCollector(f func() ([]CPUStat, error)) prometheus.Collector {
	return &collector{
		TimeUserHertzTotal: prometheus.NewDesc(
			"cpustat_time_user_hertz",
			"Timein USER_HZ a given CPU spent in a given mode.",
			[]string{"cpu", "mode"},
			nil,
		),
		stats: stats,
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TimeUserHertzTotal
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	stats, err := c.stats()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.TimeUserHertzTotal, err)
		return
	}

	for _, s := range stats {
		tuples := []struct {
			mode string
			v    int
		}{
			{mode: "user", v: s.User},
			{mode: "system", v: s.System},
			{mode: "idle", v: s.Idle},
		}

		for _, t := range tuples {
			ch <- prometheus.MustNewConstMetric(
				c.TimeUserHertzTotal,
				prometheus.CounterValue,
				float64(t.v),
				s.ID,
				t.mode,
			)
		}
	}
}

type CPUStat struct {
	ID string

	User, System, Idle int
}

func (s CPUStat) String() string {
	return fmt.Sprintf("%5s: user: %d, system: %d, idle: %d", s.ID, s.User, s.System, s.Idle)
}

func ParseCPUStats(r io.Reader) ([]CPUStat, error) {
	s := bufio.NewScanner(r)
	s.Scan() // Skip the first line

	var stats []CPUStat
	for l := 0; s.Scan(); l++ {
		line := s.Text()
		if !strings.HasPrefix(line, "cpu") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 11 {
			continue
		}

		id := fields[0]
		user, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse user in line %d", l)
		}
		system, err := strconv.Atoi(fields[3])
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse system in line %d", l)
		}
		idle, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse idle in line %d", l)
		}

		stats = append(stats, CPUStat{id, user, system, idle})
	}

	if err := s.Err(); err != nil {
		return nil, err
	}
	return stats, nil
}
