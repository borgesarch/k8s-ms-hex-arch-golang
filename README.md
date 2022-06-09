
<img src="https://kubernetes.io/images/nav_logo.svg" width="300">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/c8/OpenID_logo.svg/1200px-OpenID_logo.svg.png" width="180"> 


# #

## Microservice in hexagonal architecture build in go lang and deployment in Kluster Kubernetes

# #

[![License](https://img.shields.io/badge/License-MIT-silver.svg?style=flat)](https://github.com/clips/pattern/blob/master/LICENSE.txt) 
[![License](https://img.shields.io/badge/Go-v1.16.4-blue.svg?style=flat)](https://github.com/clips/pattern/blob/master/LICENSE.txt) 
[![License](https://img.shields.io/badge/Gorm-v1.9.15-green.svg?style=flat)](https://github.com/clips/pattern/blob/master/LICENSE.txt)
[![License](https://img.shields.io/badge/KeyCloak-13.3.1-silver.svg?style=flat)](https://github.com/clips/pattern/blob/master/LICENSE.txt) 
[![License](https://img.shields.io/badge/Kubernetes-1.20.2-blue.svg?style=flat)](https://github.com/clips/pattern/blob/master/LICENSE.txt) 
[![License](https://img.shields.io/badge/Uberdig-v1.10.0-silver.svg?style=flat)](https://github.com/clips/pattern/blob/master/LICENSE.txt) 


##

```sh
$> kubectl describe ingress  --kubeconfig=./k8s-deployment/config/config.yaml
```

```sh

Name:             ingress-resource-backend
Namespace:        default
Address:          138.197.229.52
Default backend:  default-http-backend:80 (:80)
Rules:
  Host        Path  Backends
  ----        ----  --------
  *
              /        users:80 (10.245.57.84:80)
              /users   users:80 (10.245.182.85:80)
              /todos   todo:80 (10.245.182.85:80)
Annotations:  kubernetes.io/ingress.class: nginx nginx.ingress.kubernetes.io/rewrite-target: /
Events:       <none>

```

### Features

* Hexagonal architecture in Go Language
* Bounded Contexts
* SSO/OPENID/OAUTH2/JWT
* Dependancy Injection
* Object-oriented programming
* Automatic Migrations
* Docker
* Domain-Driven Design
* Unit Test


### Deploy todo microservice

deployment backend todo

```sh
$ kubectl apply -f ./k8s-deployment/deployments/todo-api/backend --kubeconfig=./k8s-deployment/config/config.yaml
```

deployment database todo
```sh
$ kubectl apply -f ./k8s-deployment/deployments/todo-api/database --kubeconfig=./k8s-deployment/config/config.yaml
```

### Deploy user microservice


deployment backend user

```sh
$ kubectl apply -f ./k8s-deployment/deployments/user-api/backend --kubeconfig=./k8s-deployment/config/config.yaml
```

deployment database user
```sh
$ kubectl apply -f ./k8s-deployment/deployments/user-api/database --kubeconfig=./k8s-deployment/config/config.yaml
```


Init deployment ngnix ingress controller
```sh
$ kubectl apply -f ./k8s-deployment/deployments --kubeconfig=./k8s-deployment/config/config.yaml
```


### Use individual microservice

Run in docker
```sh
$ sudo docker-compose up --build
```


### INFO ###

* Autor: Gabriel Borges - Software architect
* Email: gabrielborges.web@gmail.com
* WhatsApp: +55 (62) 984887715
* Company: COFFEE DIGITAL 
