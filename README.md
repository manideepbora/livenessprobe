This repo shows a pattern in golang for readiness/liveness prob for the Kubernetes deployment

**How the codes are arranged**

The project is a very simple golang project to demonstrate the concept.
- The deployment folder contains a yaml file to deploy two deployments:applicationServer and dbServer and corresponding services. It also deploys a busybox pod to validate the functionality.
- main.go file servers as an http server and has few handlers. 
- healthMonitor.go - demonstrate the basic pattern for the monitoring. 

**How to build and create docker image on the docker registry**

1. Update the docker file according to your configuration
2. Update the build file in the ./script folder with proper configuration

**How to deploy and test the functionality**

1. Update the ./deploy/deploy.yaml file to point to the right docker image
2. Run the following kubectl command to deploy needed components to the default namespace
```
kubectl apply ./deploy/deploy.yaml
```
4. Change the ready state of db service to false by running the following command
```
kubectl exec -it busybox -n default
wget wget -qO - test-svc1:9090/switchReady
```
6. clean up the system by running the following kubectl command
```
kubectl delete ./deploy/deploy.yaml
