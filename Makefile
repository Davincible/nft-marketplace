
.PHONY: docker-build
docker-build:
	@DOCKER_BUILDKIT=1 docker build -t bsc:local -f Dockerfile.bsc .
	@DOCKER_BUILDKIT=1 docker build -t bsc-init:local -f Dockerfile.bsc-init .
	@DOCKER_BUILDKIT=1 docker build -t netstats:local -f Dockerfile.netstats .

.PHONY: docker-compose
docker-compose:
	@docker compose up

