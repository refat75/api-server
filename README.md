# API Book-Server Deployment and Services in Kubernetes Cluster

To continue further make sure that your api-server docker image is present on locally. You can check the image existence by using the following command.

```shell
$ docker image ls
```

Here I will be using `kind` for creating kubernetes cluster. Since Kind runs Kubernetes inside Docker containers, your local Docker images are not automatically available inside the Kind cluster. You need to load your locally built image into the Kind cluster before using it in a Deployment.

### Load Local Docker Image into Kind
```bash
$ kind load docker-image <your-image-name:tag>
```
Now provide the actual image name `spec.containers.image` section in `deployment.yaml` file. Also adjust your application port in `deployment.yaml` and `service.yaml` file.

## Apply YAML File

```bash
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
```

## Port Forwarding
```bash
$ kubectl port-forward svc/<your-service-name> 8080:<targetPort>
```

If everything works perfectly the application can be accessed at: [http://localhost:8080/](http://localhost:8080/)

## Clean Up Using YAML File
```bash
$ kubectl delete -f deployment.yaml 
$ kubectl delete -f service.yaml
```

# Book Server Description

## API Endpoints
|Method|Route Type|        URL       |       Description      |
|------|--------- |------------------|------------------------|
| `GET`| `Public` | `/`              |Hello Message From Server|
|`POST`| `Public` | `/login`         |For logging in. Necessary to provide credentials.|
|`POST`| `Public` | `/logout`        |For logging out.|
|`GET` |`Private` | `/books`         |Get all the listed books.|
|`POST`|`Private` | `/books`         |Create a new book. Necessary to provide book description.|
|`GET` |`Private` | `/books/{id}`    |Get a single book using book id.|
|`PUT` |`Private` | `/books/{id}`    |Update book using id.|
|`Delete`|`Private`|`/books/{id}`    |Delete a book using id.|

## Login Credentials
```json
{
    "username": "admin",
    "password": "123456"
}