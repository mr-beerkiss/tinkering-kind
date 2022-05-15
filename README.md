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

## Apps

Remember to build the deno docker file with `--platform arm64` otherwise
the container won't be loaded

## TODO (other ideas)

- [ ] [Install Istio](https://istio.io/latest/docs/setup/platform-setup/kind/)
- [ ] Configure kubectl autocompletion
- [ ] Use a smaller image for the deno app

## Reading List

- [ ] [Kubernetes CNI Explained](https://www.tigera.io/learn/guides/kubernetes-networking/kubernetes-cni/)
- [ ] [Effective Go](https://go.dev/doc/effective_go)
- [ ] [Go Maps in Action](https://go.dev/blog/maps)
- [ ] [Kuberbetes Configuration Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/#general-configuration-tips)
- [ ] [Connecting Applications with Services](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/)

## Watch List
- [ ] [Networking in Kubernetes (7min)](https://kube.academy/courses/kubernetes-in-depth/lessons/an-introduction-to-cni)
- [ ] [NAT and Firewall Explained (10min)](https://www.youtube.com/watch?v=2llWuivdS7w)

## Other Interesting Things Discovered As Part of this

### [K3s on Raspberry PI](https://bryanbende.com/development/2021/05/07/k3s-raspberry-pi-initial-setup)

A blog article about setting up a Raspberry PI cluster. Could be a fun
experiment but need a justifable application stack to run on it other
than "Hello World"

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

[Entry Level Kubernetes Certification](https://www.cncf.io/announcements/2021/10/13/entry-level-kubernetes-certification-to-help-advance-cloud-careers/)

"Official" certification for Kubernetes. Something worth considering 
but no essential. Might help with a career transition.

[kind-ingress]: https://kind.sigs.k8s.io/docs/user/ingress/#using-ingress