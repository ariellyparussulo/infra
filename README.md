# Introduction
Creating a basic infra to lean some concepts of infrastructure.

## Repository
1. Create a minikube node using `minikube start --insecure-registries=localhost`. The insecurte registries options allows you to use Harbor inside your node.
2. Add Nginx Controller inside your minikube cluster by using `minikube addons enable ingress`.
3. Install git and ansible with apt package manager: `sudo apt-get update && sudo apt-get install git ansible`.
4. Access the node using `minikube ssh` and clone it inside it.
5. Run ansible recipe inside the node by using `ansible-playbook harbor-playbook.yml` inside **registry** folder.
6. In your host, access the harbor web interface in your browser. To find out the node ip address, use `minikube ip` command.
7. To be able to upload your Docker images to the harbor cluster, use `docker login localhost`and log in on it by using the credentials that you can find out in the repository folder.
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
4. Still inside cluster folder, run [application](./cluster/application.yml) with `kubectl create -f cluster/application.yml` to create deploy the Golang application.
5. Create the proxy server by running `kubectl create -f cluster/proxy.yml`.
6. Add `backend.local` inside `/etc/hosts` poiting to `minikube ip`.
7. Use `curl http://backend.local/v1/users` to test your application.