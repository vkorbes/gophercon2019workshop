# Workshop

This is the outline for the GopherCon 2019 workshop [Go & Kubernetes Sitting In A Tree](https://www.gophercon.com/agenda/session/70232).

## Code

The code for the app used in the workshop lives here: https://github.com/campoy/links

The application is a URL shortener.

## Tools

- kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
- kubefwd: https://github.com/txn2/kubefwd/releases
- Helm: https://github.com/helm/helm#Install
- Telepresence: https://www.telepresence.io/reference/install
- Garden: https://docs.garden.io/basics/installation
- Skaffold: https://skaffold.dev/docs/getting-started/#installing-skaffold
- Tilt: https://docs.tilt.dev/install.html
- stern: https://github.com/wercker/stern#Installation
- Squash: https://github.com/solo-io/squash/releases

## Lessons

- Dockerize an application: https://github.com/campoy/links/tree/master/step1
- Run it on Kubernetes: https://github.com/campoy/links/tree/master/step2
- Split it into microservices: https://github.com/campoy/links/tree/master/step3
- Dockerize your microservices: https://github.com/campoy/links/tree/master/step4
- Can you run it on Kubernetes? https://github.com/campoy/links/tree/master/step5
- Implementing gRPC: https://github.com/campoy/links/tree/master/step6
- Implementing a gRPC gateway: https://github.com/campoy/links/tree/master/step7
- The complete application: https://github.com/campoy/links/tree/master/step8

## Notes

### Kubernetes Tooling Talk

The talk presented early in the morning was slightly adapted from this one:
- Video: https://www.youtube.com/watch?v=dIs8FwJzVI8
- Slides https://garden.slides.com/ellenkorbes/k8sdevtools?token=t3egVfZS#/

### Creating an optimized Dockerfile for Go

The main things to keep in mind are:
- *Use vendoring*, so are you continously re-build your application during development, dependencies don't have to be downloaded every single time.
- *Use multi-stage builds*, so your images are small and you don't have to wait for them to be moved around.
- Example: 
https://github.com/campoy/links/blob/master/step5/links/web.Dockerfile

### Development workflow with only kubectl&docker

The main workflow consists of:

- Make changes to your code.
- Build a new image: `docker build -t user/repo:1.0 .`
- Push the image to a registry: `docker push registry.com/user/repo:1.0`
- Update your cluster: `kubectl apply -f obj.yaml`

To do that, you need to have set up in advance:
- docker login (for the registry)
- A Dockerfile
- A Kubernetes manifest

### kubectl debugging

These basic commands might prove useful:

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

Do peruse this document so that when you need something, you know where to find it: https://kubernetes.io/docs/reference/kubectl/cheatsheet/

## Tooling overview

### kubefwd

kubefwd port forwards every single service to your local host, so you can access any endpoint as if you were in the cluster, e.g. with a hostname like `http://database/`. 

It's useful to develop a new service from scratch on your machine and have it talk to the services in your cluster, and also for debugging where you can e.g. run a bunch of weird queries on Postman locally.

Run it with `sudo kubefwd svc`, or `sudo kubefwd svc -n namespace` to target a specific namespace.

### Telepresence

Telepresence let's you switch a deployment that's running in a cluster with a local, development version of it that you have running locally. It does so via a bunch of network proxy witchcraft trickery.

To run the development version of your service where all the network activity gets proxied to/from the cluster, do `telepresence --swap-deployment <deployment-name> --expose <port> --run <./my-app>`

### skaffold

skaffold monitors the source code of your application and re-builds and re-deploys whenever part of it changes. It is very quick and simple to set up.

It can be configured to reload files within live containers so they don't have to be re-built every time. This is great for changing static assets or to work with dynamic languages like Python and Ruby.

To use skaffold you must already have your Kubernetes manifests and Dockerfiles set up.

To use skaffold with the example application, observe that you have a skaffold.yml file and run `skaffold dev`. Config: https://github.com/campoy/links/blob/master/microservices-rest/skaffold.yaml

### Tilt

Tilt is functionally equivalent to skaffold, but it's pretty.

It has a really nice dashboard that opens automatically when you start Tilt. If it doesn't—it never does for me—you can open it manually at `http://localhost:10350/`.

To run it with the example app, observe that you have a Tiltfile present and run `tilt up`. Config: https://github.com/campoy/links/blob/master/step6/links/Tiltfile

### Garden

Garden does everything skaffold and Tilt do, with a ton of extra features on top. Some of them are:
- It can run tests
- It can run in CI
- It keeps a dependency graph of your whole application
- Builds can happen in-cluster (so you don't need minikube/docker on your machine)

It's a lot of stuff, so visit https://docs.garden.io/ and check it out. Because of all the functionality that Garden has, it has a higher cognitive cost (read: learning curve) than skaffold/Tilt. Thus the trade-off to be observed when choosing a tool for the job is one of simplicity versus extra functionality.

Running the example application with Garden requires observing a few details, see this commit: https://github.com/campoy/links/pull/1/commits/6e64d7dc33dc4dda74225888d2bfd4b1beef547f

Garden then requires one project-level config file, like this: https://github.com/eysi09/links/blob/tools/microservices-rest/garden.yml

And one module-level config file for each service, like this: https://github.com/eysi09/links/blob/tools/microservices-rest/web/garden.yml

### Squash

Squash let's you attach a debugger like Delve onto running containers. You can do it on multiple containers at a time and watch them communicate line by line.

For more information about Squash, see: https://squash.solo.io/overview/

And for a primer on using Delve (a debugger made specifically for go) see: https://medium.com/@ellenkorbes/debug-your-go-code-without-100-extra-printlns-384a86437f2b
