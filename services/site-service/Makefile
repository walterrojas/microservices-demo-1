docker_tag ?= latest
docker_image ?= moorara/site-service
docker_test_image ?= site-service-test

path := $(shell pwd)
docker_dir := /tmp/workspace


clean:
	@ rm -rf coverage .nyc_output component-tests.log

docker:
	@ docker build -t $(docker_image):$(docker_tag) .

docker-test:
	@ docker build \
		--file Dockerfile.test \
		--build-arg version=$(docker_tag) \
		--tag $(docker_test_image) .

up:
	@ VERSION=$(docker_tag) \
	  docker-compose up -d site-service

down:
	@ VERSION=$(docker_tag) \
	  docker-compose down

test-docker: # docker docker-test
	@ VERSION=$(docker_tag) \
	  docker-compose run site-service-test npm run test && \
		docker-compose down

test-component-docker: # docker docker-test
	@ VERSION=$(docker_tag) \
	  docker-compose run site-service-test npm run test:component && \
		docker container logs site-service | grep '^{' | jq . > component-tests.log && \
		docker-compose down

push:
	@ docker push $(docker_image):$(docker_tag)

save-images:
	@ mkdir -p $(docker_dir)
	@ docker image save -o $(docker_dir)/docker.tar $(docker_image):$(docker_tag)
	@ docker image save -o $(docker_dir)/docker-test.tar $(docker_test_image)


.PHONY: clean
.PHONY: docker docker-test
.PHONY: up down
.PHONY: test-docker test-component-docker
.PHONY: push save-images
