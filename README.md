# Kubernetes Operator

## Build

first build the docker image
```
docker build -t webui-controller .
```
then save the image and load it into minikube
```
docker image save -o webui-controller.tar webui-controller
minikube image load webui-controller.tar
rm webui-controller.tar
```
or in one command
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