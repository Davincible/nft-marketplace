
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
	@mkdir -p bsc socket
	@docker compose -f docker-compose.bsc-bootstrap.yaml up 

.PHONY: docker-compose
docker-compose:
	@docker compose up --remove-orphans 

.PHONY: init 
init:
	@npm install

.PHONY: compile-contracts
compile-contracts:
	@solc --abi --bin contracts/*.sol --overwrite --optimize -o build @openzeppelin=node_modules/@openzeppelin

.PHONY: gen-bindings-contract
gen-bindings-contract:
	@abigen --abi build/$(contract).abi --pkg contracts --type $(contract) --out contracts/$(contract).go --bin build/$(contract).bin

.PHONY: gen-bindings
gen-bindings:
	@make --no-print-directory gen-bindings-contract contract=Greeter
	@make --no-print-directory gen-bindings-contract contract=KuiperNFT
	@make --no-print-directory gen-bindings-contract contract=KuiperNFTFactory
	@make --no-print-directory gen-bindings-contract contract=KuiperNFTMarketplace

