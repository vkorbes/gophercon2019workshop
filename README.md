# Workshop

[Go & Kubernetes Sitting In A Tree](https://www.gophercon.com/agenda/session/70232)!

# TODO

- This doc.
- Slides.
- Cloud set-up.
- E-mail attendees installation instructions.

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

```dockerfile
FROM golang:1.11.9-alpine3.9
WORKDIR /go/src/app
RUN apk add --no-cache git gcc musl-dev
ENV GO111MODULE=on
EXPOSE 8080
COPY main.go go.mod ./
RUN go build -gcflags "all=-N -l" main.go
CMD ["./main"]
```

```yml
apiVersion: v1
kind: Pod
metadata:
name: mycontainer
spec:
containers:
    - image: registry.com/user/repo:1.0
    name: mycontainer
    ports:
        - containerPort: 8080
        name: http
        protocol: TCP
```

Ideally you would set up a combination of deployment/service/ingress:

```yml
 apiVersion: extensions/v1beta1
 kind: Deployment
 metadata:
   name: my-deployment
 spec:
   replicas: 1
   selector:
     matchLabels:
       app: my-app
   template:
     metadata:
       labels:
         app: my-app
     spec:
       containers:
       - name: my-app
         image: registry.com/user/repo:1.0
         imagePullPolicy: Always
         ports:
         - containerPort: 5000

 apiVersion: v1
 kind: Service
 metadata:
   name: my-deployment
 spec:
   ports:
   - port: 5000
     targetPort: 5000
   selector:
     app: my-app

apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: test-ingress
spec:
  backend:
    serviceName: my-app
    servicePort: 80
```

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