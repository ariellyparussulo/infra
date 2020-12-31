# Introduction
Creating a basic infra to lean some concepts of infrastructure.

## Backend
Small go server with a POST and GET requests with its respective Dockerfile to generate an image to be used by our infrastructure (it will be done later). To test this image, use:

```sh
cd backend/
docker build -t TAG_HERE .
docker run -p 8000:8000 CONTAINER_ID
```

And for now, it accepts the following requests:

```sh
curl http://localhost:8000/v1/users/123
curl -X POST -H 'Content-Type: application/json' -d '{ "id": 123, "name": "arielly" }' http://localhost:8000/v1/users/
```