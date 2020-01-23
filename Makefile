
PROGRAM := pg_tileserv
CONTAINER := pramsey/pg_tileserv

.PHONY: all check clean test docker

GOFILES := $(wildcard *.go)

all: $(PROGRAM)

check:
	@go version

clean:
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -f $(CONTAINER)

$(PROGRAM): $(GOFILES)
	go build -v

docker: $(PROGRAM) Dockerfile
	docker build -f Dockerfile --build-arg VERSION=`./$(PROGRAM) --version | cut -f2 -d' '` -t pramsey/pg_tileserv:latest .

test:
	go test -v



# .PHONY: docker-clean
# docker-clean: ## Remove any Docker volumes created during development
# docker-clean: | docker-clean-volumes

# .PHONY: docker-clean-volumes
# docker-clean-volumes:
# 	docker-compose -f docker/docker-compose.build.yaml down --volumes

# .PHONY: docker-images
# docker-images: apiserver-image keycloak-image operator-events-agent-image

# .PHONY: help
# help: ALIGN=14
# help: ## Print this message
# 	@awk -F ': ## ' -- "/^[^':]+: ## /"' { printf "'$$(tput bold)'%-$(ALIGN)s'$$(tput sgr0)' %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# .PHONY: apiserver-image
# apiserver-image:
# 	@docker build -t crunchydata/crunchy-cloud-apiserver:$(VERSION) \
# 			--build-arg VERSION=$(VERSION) \
# 			-f $(DOCKER_DIR)/Dockerfile.apiserver .

# .PHONY: keycloak-image
# keycloak-image: ensure-keycloak-user-storage-provider
# 	@echo "Building keycloak image..."
# 	@docker build \
# 			--build-arg USER_PROVIDER_VERSION=$(KEYCLOAK_USER_STORE_VERSION) \
# 			-t crunchydata/crunchy-cloud-keycloak:$(KEYCLOAK_IMAGE_VERSION) \
# 			-f $(DOCKER_DIR)/Dockerfile.keycloak .

# .PHONY: operator-events-agent-image
# operator-events-agent-image:
# 	@$(info Building operator-events-agent image...)
# 	@docker build -t crunchydata/operator-events-agent:$(VERSION) \
# 			--build-arg VERSION=$(VERSION) \
# 			-f $(DOCKER_DIR)/Dockerfile.operator-events-agent .

# .PHONY: ensure-keycloak-user-storage-provider
# ensure-keycloak-user-storage-provider:
# ifeq ("$(wildcard $(BUILD_DIR)/$(KEYCLOAK_USER_STORE_FILE))","")
# 	$(info Artifact $(BUILD_DIR)/$(KEYCLOAK_USER_STORE_FILE) does not exist.)
# 	$(info Downloading: $(KEYCLOAK_USER_STORE_FILE))

# 	@$(TOOLS_DIR)/gh-dl-release \
# 		$(KEYCLOAK_USER_STORE_REPO) \
# 		$(KEYCLOAK_USER_STORE_FILE) \
# 		$(KEYCLOAK_USER_STORE_VERSION)\
# 		$(BUILD_DIR)/$(KEYCLOAK_USER_STORE_FILE)
# endif

