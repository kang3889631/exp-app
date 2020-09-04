SHELL = /bin/bash -o pipefail

include .env

UBUNTU_CODENAME := $(shell lsb_release -cs)

hello:
	#@test -n "$(msg)" || (echo ">> msg flag is not set. e.g. msg=good"; exit 1)
	#echo $(WelcomeMsg) $(msg)
	#@echo $(CLI_DIR)
	#@echo $(SRV_DIR)
	#@echo $(CLI_BUILD_OUTPUT_NAME)

cli-watch:
	cd $(CLI_DIR) && ng build --aot --output-path ../$(CLI_BUILD_OUTPUT_NAME) --base-href /$(CLI_BUILD_OUTPUT_NAME)/ --watch

srv-watch:
	cd $(SRV_DIR) && air

go-build:
	cd $(SRV_DIR) && go mod vendor
	cd $(SRV_DIR) && go build -mod vendor -o ./bin/main .

go-start:
	cd $(SRV_DIR) && ./bin/main

rebuild-mysql:
	cd $(DOCKER_DIR) && docker-compose rm -sf "mysql" && docker-compose up -d --no-deps --build "mysql"

restart-mysql:
	cd $(DOCKER_DIR) && docker-compose restart "mysql"

stop-mysql:
	cd $(DOCKER_DIR) && docker-compose stop "mysql"

exec-mysql:
	cd $(DOCKER_DIR) && docker-compose exec -e TERM=$$TERM -e LINES=$$LINES -e COLUMNS=$$COLUMNS "mysql" bash

log-mysql:
	cd $(DOCKER_DIR) && docker-compose logs "mysql"

