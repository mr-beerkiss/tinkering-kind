# TODO: Add colours!
# TODO: Possibly rewrite in Go!?
if kind get clusters | grep -q bx-cluster; then
    echo "Old cluster detected... deleting"
    kind delete cluster --name bx-cluster
fi

echo "curl Creating cluster"
kind create cluster --config config/bx-cluster.yaml

echo "Setup contour ingress"
kubectl apply -f config/contour.yaml

# TODO: I could apply the patch directly to the config instead of using 
# this command. Need to make sure it stays up to date!
echo "Patch Countour ingress for kind"
kubectl patch daemonsets -n projectcontour envoy -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/control-plane","operator":"Equal","effect":"NoSchedule"},{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'

# TODO: The image is hardcoded here, not sure if there is a better way
# to do it. Also, might need to build or pull the image first... at
# the moment it's assumed it's exists and is available on the host
echo "load application image into cluster"
kind load docker-image kubernetes-bootcamp-go:v0.7 kubernetes-bootcamp-deno:v0.4 --name bx-cluster

# TODO: This failed the first time I ran it, not sure why. Error 
# Error from server (Forbidden): error when creating "config/bx-app.yaml": pods "foo-app" is forbidden: error looking up service account default/default: serviceaccount "default" not found
#echo "Create the test service!"
#kubectl apply -f config/bx-app.yaml
