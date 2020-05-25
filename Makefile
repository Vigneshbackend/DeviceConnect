NAME := device-connect
DOCKER_REPO := gcr.io/okcredit-42/$(NAME)
docker:
	docker build -t $(NAME) .
docker-run: docker
	docker run --env-file config.env -p "80:80" $(NAME)
docker-push: TAG ?= latest
docker-push: docker
	docker tag $(NAME) $(DOCKER_REPO):$(TAG)
	docker push $(DOCKER_REPO):$(TAG)