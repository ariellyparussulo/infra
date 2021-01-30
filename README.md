# Introduction
Creating a basic infra to lean some concepts of infrastructure.

## Repository
1. Create a minikube node using `minikube start --insecure-registries=localhost`. The insecurte registries options allows you to use Harbor inside your node.
2. Install git and ansible with apt package manager: `sudo apt-get update && sudo apt-get install git ansible`.
3. Access the node using `minikube ssh` and clone it inside it.
4. Run ansible recipe inside the node by using `ansible-playbook harbor-playbook.yml` inside **registry** folder.
5. In your host, access the harbor web interface in your browser. To find out the node ip address, use `minikube ip` command.
6. To be able to upload your Docker images to the harbor cluster, use `docker login localhost`and log in on it by using the credentials that you can find out in the repository folder.
## Backend (and testing it)
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

## Creating initial infrastructure
1. Create a config map with the following configurations:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-configuration
  labels:
    app: postgres
  namespace: infra
data:
  POSTGRES_DB: backend
  POSTGRES_USER: postgres
  POSTGRES_PASSWORD: postgres
```

2. Run `kubectl create -f YOUR_CONFIG_MAP.yaml`.
3. Access **the cluster configuration folder** and run [postgres](./cluster/postgres.yml) with `kubectl create -f cluster/postgres.yml`

It will create a replica of postgress database pod and expose its port inside the Kubernetes cluster.