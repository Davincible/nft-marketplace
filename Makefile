
.PHONY: docker-build-bsc
docker-build-bsc:
	@DOCKER_BUILDKIT=1 docker build -t bsc:local -f Dockerfile.bsc .

.PHONY: docker-build-bootstrap
docker-build-bootstrap:
	@make docker-build-bsc
	@DOCKER_BUILDKIT=1 docker build -t bsc-bootstrap:local -f Dockerfile.bsc-bootstrap .

.PHONY: docker-build
docker-build:
	@make docker-build-bsc
	@DOCKER_BUILDKIT=1 docker build -t bsc-init:local -f Dockerfile.bsc-init .
	@DOCKER_BUILDKIT=1 docker build -t netstats:local -f Dockerfile.netstats .

.PHONY: bsc-bootstrap
bsc-bootstrap:
	@make docker-build-bootstrap
	@docker compose -f docker-compose.bsc-bootstrap.yaml up 

.PHONY: docker-compose
docker-compose:
	@docker compose up --remove-orphans

.PHONY: compile-contracts
compile-contracts:
	@solc --abi --bin contracts/Greeter.sol -o build

.PHONY: gen-bindings
gen-bindings:
	@abigen --abi build/Greeter.abi --pkg contracts --type Greeter --out contracts/Greeter.go --bin build/Greeter.bin

