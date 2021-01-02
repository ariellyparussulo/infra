# Introduction
Creating a basic infra to lean some concepts of infrastructure.

## Repository
An ansible project that install a harbor locally (it will be changed later) to stores the Docker images of this project. To run it, use:

```sh
ansible-playbook harbor-playbook.yml --ask-become-pass
```

And update `/etc/docker/daemon.json` with:

```json
{
        "insecure-registries" : ["localhost"]
}
```

After the docker-compose stack is up, create a project called `infra` in Harbor.

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
curl -X POST -H 'Content-Type: application/json' -d '{ "id": 123, "name": "arielly" }' http://localhost:8000/v1/users/
```