PROJECTNAME=$(shell basename "$(PWD)")
DOCKERHUBID := refat0
MAKEFLAGS += --silent

hello: help
	echo
	echo "Current Dockerhub Username: $(DOCKERHUBID)"
	echo "Hello from Makefile. Use the above command to use the project"

##run: Running the Server on port 8080.
run:
	go run main.go

##start: Start the docker image as a Container
start: build
	echo "Running the Docker image..."
	docker run -d -p 8080:8080 --rm $(DOCKERHUBID)/$(PROJECTNAME)

##stop: Stop the docker Container
stop:
	docker stop $$(docker ps -q --filter "ancestor=$(DOCKERHUBID)/$(PROJECTNAME)")
	echo "Container Stopped"


##build: Build the Docker Image using Dockerfile.
build:
	echo "Building the Docker Image"
	docker build -t $(DOCKERHUBID)/$(PROJECTNAME) -f Dockerfile .

##push: Push the build image to dockerhub repository.
push:
	docker push $(DOCKERHUBID)/$(PROJECTNAME)

##clean: Delete the locally stored docker image
clean:
	docker rmi $(DOCKERHUBID)/$(PROJECTNAME)

help: Makefile
	echo "Choose a command to run in "$(PROJECTNAME)":"
	echo
	sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'