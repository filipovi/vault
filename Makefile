# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help all build-proto build-docker

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help 

build-proto: ## Generates the protobuf files   
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) filipovi/docker-protobuf --go_out=. -I. proto/generator.proto --micro_out=.

build-docker: build-docker-vault build-docker-micro-vault ## Build the docker images

build-docker-vault: ## Build the docker image for the web server
	docker build -t vault .

build-docker-micro-vault: ## Build the docker image for the micro service
	docker build -t micro-vault server/.

all: build-proto build-docker ## build all the project

minikube-start: ## Start Minikube
	minikube start --logtostderr --vm-driver kvm2

minikube-stop: ## Stop Minikube
	minikube stop

minikube-delete: ## Delete Minikube
	minikube delete
