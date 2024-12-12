# Kubernetes Operator

First step to write kubernetes operator without using operator-sdk or kubebuilder.

## Build

first build the docker image
```
docker build -t webui-controller .
```
then load image to minikube
```
minikube image load webui-controller
```

## Deploy

create the CRD
```
kubectl apply -f yaml/crd.yaml
```
then deploy the controller
```
kubectl apply -f yaml/deploy-controller.yaml
```
and finally deploy the example
```
kubectl apply -f yaml/example.yaml
```