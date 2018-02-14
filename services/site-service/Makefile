docker_image ?= moorara/site-service
docker_tag ?= latest

path := $(shell pwd)
report_dir := $(path)/coverage


clean:
	@ rm -rf $(report_dir)

docker:
	@ docker build -t $(docker_image):$(docker_tag) .

up:
	@ VERSION=$(docker_tag) \
	  docker-compose up -d site-service

down:
	@ VERSION=$(docker_tag) \
	  docker-compose down


docker-test:
	@ docker build \
		--file Dockerfile.test \
		--build-arg version=$(docker_tag) \
		--tag site-service-test .

test: docker docker-test
	@ VERSION=$(docker_tag) \
	  docker-compose run site-service-test npm run test

test-component: docker docker-test
	@ VERSION=$(docker_tag) \
	  docker-compose run site-service-test npm run test:component


.PHONY: clean
.PHONY: docker up down
.PHONY: docker-test test test-component