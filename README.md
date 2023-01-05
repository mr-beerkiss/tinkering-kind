# Tinkering with Kind

Attempting to play with Kubernetes again. Having pretty good success 
with kind. Saving some notes down here since the number of tabs is 
growing unsustainable.

This repo contains the configuration for a kind cluster with two (very)
basic testing apps written in Go and Deno.

The hope is eventually to grow this project into something worth of
a Raspberry PI Cluster that does something kinda useful.

## Install Instructions

Run `setup.sh`

> **NOTE** I'm using Countour for Ingress since it seems to have the 
most straightforward config. Would be good to better understand the 
trade-offs between the ones on the kind website. See 
[Setting Up an Ingress Controller][kind-ingress] for details

Something interesting, instead of `kubectl proxy` or settin up `ingress`. 
You can also use [`kubectl port-forward`](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/) to expose pods to the host for 
testing. Could be useful for something connecting to a DB while you develop
the code on the host. i.e. the DB runs in the cluster so you don't have 
to install it on the host.

## Apps

Remember to build the deno docker file with `--platform arm64` otherwise
the container won't be loaded

## TODO (other ideas)

- [ ] [Install Istio](https://istio.io/latest/docs/setup/platform-setup/kind/)
- [ ] Configure kubectl autocompletion
- [ ] Use a smaller image for the deno app
- [ ] Switch Pod manifests to [deployments](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/)
- [ ] Add some basic auth to the Go app and set some stuff via configMaps & secrets
- [ ] Install [Redis](https://kubernetes.io/docs/tutorials/configuration/configure-redis-using-configmap/) into the cluster to use as a KV store
- [ ] Run local cluster under https (See mkcert below)
- [ ] Enable metrics-server in kind and use it to calculate CPU + RAM 
limits
- [ ] [Install Tekton](https://tekton.dev/docs/getting-started/) for automated builds (CI/CD)
- [ ] Try [k9s](https://github.com/derailed/k9s)
- [ ] [staticcheck](https://staticcheck.io/docs/getting-started/) for static analysis & linter for golang
- [ ] Setup Cluster with Terraform - [Example](https://nickjanetakis.com/blog/configuring-a-kind-cluster-with-nginx-ingress-using-terraform-and-helm)

## Reading List

- [ ] [Kubernetes CNI Explained](https://www.tigera.io/learn/guides/kubernetes-networking/kubernetes-cni/)
- [ ] [Effective Go](https://go.dev/doc/effective_go)
- [ ] [Go Maps in Action](https://go.dev/blog/maps)
- [ ] [Kuberbetes Configuration Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/#general-configuration-tips)
- [ ] [Connecting Applications with Services](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/)
- [ ] [Kubernetes: kubectl wait](https://enix.io/en/blog/kubernetes-tips-tricks-kubectl-wait/)
- [ ] [Managing Resources](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/)
- [ ] [Config Maps](https://kubernetes.io/docs/concepts/configuration/configmap/)
- [ ] [Redis as a Primary Database](https://redis.com/blog/redis-cache-vs-redis-primary-database-in-90-seconds/)
- [ ] [How SSDs Really Work](https://arstechnica.com/information-technology/2012/06/inside-the-ssd-revolution-how-solid-state-disks-really-work/) - more info on using Redis as primary DB
- [ ] [Minio](https://min.io) a OSS S3 type thing (i.e. kubernetes object storage)
- [ ] [Uptrace](https://get.uptrace.dev/guide/opentelemetry-tracing-tool.html#clickhouse) Distributed Tracing Tool
- [ ] [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [ ] [lo](https://github.com/samber/lo) Lodash like lib for Go.

## Watch List
- [x] [Networking in Kubernetes (7min)](https://kube.academy/courses/kubernetes-in-depth/lessons/an-introduction-to-cni)
- [x] [NAT and Firewall Explained (10min)](https://www.youtube.com/watch?v=2llWuivdS7w)

## Other Interesting Things Discovered As Part of this

### [K3s on Raspberry PI](https://bryanbende.com/development/2021/05/07/k3s-raspberry-pi-initial-setup)

A blog article about setting up a Raspberry PI cluster. Could be a fun
experiment but need a justifable application stack to run on it other
than "Hello World"

### [mkcert]

A tool for creating locally trusted development certificates.

Note the install recommends installing `nss` (if you use Firefox) but
I'm not sure what it is... this might be why

> `Warning: "certutil" is not available, so the CA can't be automatically installed in Firefox! ⚠️
Install "certutil" with "brew install nss" and re-run "mkcert -install" 👈`

A `TODO` here is figuring out of it is possible to use these certs with
kind! And like magic here is an [example](https://github.com/dgafka/local-kuberentes-cluster-over-https)
Note the example requires `ctlptl` which is _another_ orchestration 
tool that creates an abstraction over common `localhost` kubernetes 
installs. 

### [Enhanced Echo Server](https://github.com/mauilion/echo-server)

A fancy echo server that supports websockets and a bunch of other 
things. According to the ReadMe it's helpful for testing Load Balancers,
Web Servers and Browsers. It's a fork, but the owner of the fork looks
like an interesting guy: [Duffie Cooley](https://mauilion.dev). Calls
himself a field CTO. His repo for [blackhat-2019](https://github.com/mauilion/blackhat-2019) 
inspired me to go a bit further in trying to documenting my journey so
far.

### [Multipass](https://github.com/canonical/multipass)

From Canonical, for creating virual Ubuntu instances on the command line
. Kinda like vagrant but uses macOS hypervisor rather than relying on 
VirtualBox.

### [kubecolor](https://github.com/hidetatz/kubecolor)

Colorised kubectl output... but the installation methods aren't working.

### [Rancher](https://rancher.com/docs/rancher/v2.6/en/overview/architecture/)

Easy Clusters management.

### [Transparently running binaries from any architecture in Linux with QEMU and binfmt_misc](https://ownyourbits.com/2018/06/13/transparently-running-binaries-from-any-architecture-in-linux-with-qemu-and-binfmt_misc/)

A deep dive into how QEMU can run different CPU architectures 
transparently. It's how Docker can do run containers build with a 
different arch to the host without any config (although it does spit out
a warning!). It's been a journey realising how "assumed amd64" 
everything is until you switch to ARM!

## Kubernetes plugin managers

[Helm]() is the obvious one. But there is also [Krew](https://krew.sigs.k8s.io). I don't know when these
people find time to go outside or eat with the speed with which these ecosystem seems to move!
At least someone else has done some analyis on [Helm vs kubectl](https://medium.com/@RedBaronDr1/helm-vs-kubectl-5aaf2dba7d71)... oh wait that's not _`krew` vs `helm`_
### [metrics-server](https://github.com/kubernetes-sigs/metrics-server)

Would be good to figure out how to enable the `metrics-server` on kind. 
There are tools like [kube-ops-view](https://codeberg.org/hjacobs/kube-ops-view) and 
[ktop]() that can help visualise resource use in the cluster. An [example
config](https://gist.github.com/hjacobs/69b6844ba8442fcbc2007da316499eb4)
from the author of kube-ops-view. Seems complicated.

### Redis

- [Deploying Redis Cluster on Kubernetes](https://www.containiq.com/post/deploy-redis-cluster-on-kubernetes)

When you have Redis in the Cluster. You can use the [Go Redis](https://redis.uptrace.dev) client.

### [Open Application Model](https://oam.dev)

Came up while talking to Brendan today. Seems to be an attempt to build
applications in a way that is agnostic to current underlying models
(e.g containers, kubernetes) to keep things portable across emerging
technologies. Sounds a bit wishy washy but worth a look.


### [makefiles.dev](https://github.com/make-files/makefiles)

Another interesting find, opinionated makefiles you download off the internet (😱). Not really 
sure how I feel about that but I guess it's good enough to get my going. 

### [Entry Level Kubernetes Certification](https://www.cncf.io/announcements/2021/10/13/entry-level-kubernetes-certification-to-help-advance-cloud-careers/)

"Official" certification for Kubernetes. Something worth considering 
but no essential. Might help with a career transition.

[kind-ingress]: https://kind.sigs.k8s.io/docs/user/ingress/#using-ingress
