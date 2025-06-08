# Platitude

A demonstration project showcasing Bazel build system usage with multi-language microservices.

## Overview

Platitude consists of two REST APIs:

- A Python-based API service
- A Golang-based API service

Each service exposes an endpoint that returns random platitudes about their respective programming languages. The services are containerized, built using Bazel, and designed to be deployed on Kubernetes.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Bazel](https://bazel.build/)
- [Minikube](https://minikube.sigs.k8s.io/)
- [Docker](https://www.docker.com/)

> On MacOS, these can be installed using homebrew.

## Project Structure

```
platitude/
├── python-api/      # Python API service
├── go-api/          # Golang API service
├── k8s/             # Kubernetes manifests
├── scripts/         # Build/deployment script
```

## Building the Project

To build all services:

```sh
make build
```

## Running Locally

### Using Minikube

1. Start Minikube:
   gs

```sh
minikube start
```

2. Deploy services:

```sh
make deploy
```

## API Endpoints

### Python Service

- `GET /health` : Returns the health status of the service
- `GET /api/platitude`: Returns a random Python-related platitude

### Golang Service

- `GET /health` : Returns the health status of the service
- `GET /api/platitude`: Returns a random Go-related platitude

### Testing

To hit the endpoints, you can use the following commands:

```sh
curl localhost:8081/python/platitude

{"message":"Don't reinvent the wheel, just pip install it."}
```

```sh
curl localhost:8082/golang/platitude

{"message":"Less is more."}
```
