# Workshop

[Go & Kubernetes Sitting In A Tree](https://www.gophercon.com/agenda/session/70232)!

# TODO

- This doc.
- Slides.
- Cloud set-up.
- E-mail attendees installation instructions.
- minikube installation instructions.

# Schedule


### 09:00 — Introduction:
- Overview of the Go part.
- Overview of the Kubernetes part.
- Overview of the example app. (It has 3 services.)

### 10:30 — Short break.

### 10:50 — First steps & interacting with pods in a cluster:
- Something something Go.
- Deploy service 1/3 to cluster—old school kubectl&docker.
- Local Go app talks to remote service using _kubefwd._
- Something something Go.
- Install Dgraph with _Helm._
- Deploy service 2/3—old school kubectl&docker.
- Service 3/3, still local, talks to services 1 & 2 using _Telepresence._

### 12:00 — Lunch.

### 13:00 — Advanced techniques & the optimal development workflow:
- Something something Go.
- Using _Skaffold & Tilt_ as a quick and easy solution.
- Something something Go.
- Delete 2/3 manifests.
- Using _Garden_ for complex systems. Dependencies, tests (unit & e2e), in-cluster building.
- Something something Go.

### 15:00 — Short break.

### 15:20 — Debugging:
- Something something Go.
- Basics of _kubectl debugging._
- Something something Go.
- Distributed debuggers with _Squash._

### 16:30 — Closing words:
- Conclusion for Go.
- Conclusion for Kubernetes.

### 17:00 — Beer! 🍻

Total: 6h 20m.

# Installation

## Bundle

Download here:
- Linux
- macOS
- Windows

Then add the location to your PATH with e.g. `export PATH=$PATH:~/location`

## DIY

- kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
- kubefwd: https://github.com/txn2/kubefwd/releases
- Helm: https://github.com/helm/helm#Install
- Telepresence: https://www.telepresence.io/reference/install
- Garden: https://docs.garden.io/basics/installation
- Skaffold: https://skaffold.dev/docs/getting-started/#installing-skaffold
- Tilt: https://docs.tilt.dev/install.html
- stern: https://github.com/wercker/stern#Installation
- Squash: https://github.com/solo-io/squash/releases

# Kubernetes cluster

For this workshop we'll need a Kubernetes cluster. We recommend one of three options:

1. minikube
2. Docker for Desktop
3. Digital Ocean

## minikube

- Install minikube
- Set line 51 of `apps/manifests.yml` to `NodePort`.
- `kubectl apply -f apps/manifests.yml`
- `minikube service [deployment-name] --url`
- `curl [deployment-url:port]`

## Docker for Desktop

## Digital Ocean

Digital Ocean has graciously provided us credits for us to use their managed Kubernetes cluster offering for this workshop.

### Follow these steps to set it up:
- Create a Digital Ocean account if you don't have one
- Click `API` on the dashboard, then `Generate New Token`.
- Install the `doctl` tool, either from the `bin/` folder, or using one of the package managers listed [here](https://github.com/digitalocean/doctl#installing-doctl).
- `doctl auth init -t [your-api-token]`
- Create a Kubernetes cluster using DO's dashboard. (Create → Clusters). Minimum settings should work.
- `doctl kubernetes cluster kubeconfig save [cluster-name]`kube
- Create a load balancer in the same region as your cluster. (Create → Load Balancers)
- Now wait for an external IP using `kubectl get service dep-svc1 -w`.
- You can now access it e.g. `curl [external-ip:port]`.
- Lastly, once the workshop is over, go to the Kubernetes and Load Balancer pages and delete them.

### To do this using mostly the command line:
- Create your token as explained above.
- `export TOKEN=[your-token]`
- Create a Kubernetes cluster with `doctl k8s cluster create gophercluster --count 1 --size s-1vcpu-2gb --region sfo2`.
- List your droplets and take note of the ID of the droplet you just created. The droplet will be called something like `gophercluster-default-pool-oz6v`. Droplet list: `curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" "https://api.digitalocean.com/v2/droplets?page=1&per_page=1"`
- Create a load balancer, using your droplet's ID:
```
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" -d '{"name":"gopherbalancer","region":"sfo2","droplet_ids":[<DROPLET_ID1>], "forwarding_rules": [
{
  "entry_protocol": "http",
  "entry_port": 80,
  "target_protocol": "http",
  "target_port": 80
}
]}' "https://api.digitalocean.com/v2/load_balancers"
```
- Now wait for an external IP using `kubectl get service dep-svc1 -w`, and access it `curl [external-ip:port]`.
- For cleanup, start by deleting your Kubernetes cluster: `doctl k8s cluster delete gophercluster`
- Get the ID of your load balancer: `curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" "https://api.digitalocean.com/v2/load_balancers"`
- And delete it: `curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" "https://api.digitalocean.com/v2/load_balancers/<ID>"`

Further reference: https://developers.digitalocean.com/documentation/v2/

# Tools

## No tools: docker&kubectl

The main workflow consists of:

- Make changes to your code.
- Build a new image: `docker build -t user/repo:1.0 .`
- Tag it: `docker tag image registry.com/user/repo:1.0`
- Push the image to a registry: `docker push registry.com/user/repo:1.0`
- Update your cluster: `kubectl apply -f obj.yaml`

To do that, you need to have set up in advance:
- docker login (for the registry)
- Dockerfile
- Kubernetes manifest

See `app/manifests.yml` and `app/service1/Dockerfile` for examples of Kubernetes manifests and a Dockerfile.

## Helm

- What's Helm?
- `helm install stable/wordpress`

## Connect:

### kubefwd

- What's kubefwd?
- kubefwd in action.

### Telepresence

- What's kubefwd?
- kubefwd in action.

## Build & Deploy:

### Garden

What it does:
- Build & deploy
- Stack graph/dependencies
- Tests
- In-cluster building
- Hot reload
- Dashboard
- Log streaming

What you need:
- Describe your system via config files

Best fit for:
- Complex projects with lots of moving parts

### Skaffold

What it does:
- Build & deploy
- Hot reload

What you need:
- Simple configs
- Kubernetes manifests

Best fit for:
- Quickly setting up simple projects

### Tilt

What it does:
- Everything Skaffold does
- Dashboard
- Log streaming

What you need:
- Same as Skaffold

Best fit for:
- Same as Skaffold

## Debug:

### kubectl

- `kubectl apply -f obj.yaml`
- `kubectl delete -f kuard-pod.yaml`
- `kubectl exec -it pod -- bash`
- `kubectl cp <pod-name>:/path/to/remote/file /path/to/local/file`
- `kubectl delete <resource-name> <obj-name>`
- `kubectl logs -f <pod>`
- `kubectl get pods/ns`
- `kubectl describe pods <pod>`
- `kubectl port-forward <pod> 8080:8080`
- `kubectl expose deployments <name> --port=2368`
- `kubectl run <name> --image=registry.com/user/repo:1.0`

### Squash

- Attach a debugger to a running container
- Multiple containers simultaneously
- Using an IDE

### Stern

- What it is
- How to filter log streams