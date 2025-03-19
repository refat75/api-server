# API Server with Makefile Commands
```bash
➤ make 
Choose a command to run in api-server:

 run     Running the Server on port 8080.
 start   Start the docker image as a Container
 stop    Stop the docker Container
 build   Build the Docker Image using Dockerfile.
 push    Push the build image to dockerhub repository.
 clean   Delete the locally stored docker image

Hello from Makefile. Use the above command to use the project

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