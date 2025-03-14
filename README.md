# API Book-Server Deployment and Services in Kubernetes Cluster using HELM

### Install the Helm Chart
```shell
$ helm install my-app api-server-helm-chart
```
### Check the  installation
```shell
$ helm list -a
```
If everything is correct you will see the name `my-app` in the list.

### Port Forwarding
```bash
$ kubectl port-forward svc/<your-service-name> 8080:<targetPort>
```
If everything works perfectly the application can be accessed at: [http://localhost:8080/](http://localhost:8080/)

### Uninstall the Helm Chart
```shell
$ helm uninstall my-app
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
```

# Tutorial Link
1. [Complete Helm Chart Tutorial: From Beginner to Expert Guide](https://www.youtube.com/watch?v=DQk8HOVlumI)