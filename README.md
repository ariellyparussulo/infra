# Introduction
Creating a basic infra to lean some concepts of infrastructure.

## Repository
Run `minikube start --insecure-registries="localhost"` to start your minikube with this insecure register that you will create listed. After that, garantee that you have ansible and git installed in your machine.
Then, clone this repository and run `ansible-playbook harbor-playbook.yml` to create Harbor registry in your machine. Every image that we will create in this project will be stored there. To create a new project in harbor, access it using a browser by its IP. To get minibube's IP, use `minikube ip` command. Creata a project called `infra` inside Harbor.

To create template in Harbor registry, you need to be logged in. Use the following command to login in. The credentials can be fount in the inventory configurations.

```sh
docker login localhost
```

## Backend
Small go server with a POST and GET requests with its respective Dockerfile to generate an image to be used by our infrastructure (it will be done later). To test this image, use:

```sh
cd backend/
docker build -t localhost/infra/backend:1.0 .
docker run -p 8000:8000 CONTAINER_ID
docker push localhost/infra/backend:1.0
```

And for now, it accepts the following requests:

```sh
curl http://localhost:8000/v1/users/123
curl -X POST -H 'Content-Type: application/json' -d '{ "firstName": "name", "lastName": "name", "cpf": "000.000.000-00", "rg": "00.000.000-00" }' http://localhost:8000/v1/users/
```

kubectl create secret generic regcred --from-file=.dockerconfigjson=/home/USER/.docker/config.json --type=kubernetes.io/dockerconfigjson --namespace=infra