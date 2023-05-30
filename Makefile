.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = v10.0.0

APP 			= ginadmin
SERVER_BIN  	= ${APP}
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

CONFIG_DIR      = ./configs
CONFIG_FILE     = dev
STATIC_DIR      = ./build/dist

all: start

start:
	@go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" main.go start --configdir $(CONFIG_DIR) --config $(CONFIG_FILE) --staticdir $(STATIC_DIR)

build:
	@go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN)

# go install github.com/google/wire/cmd/wire@latest
wire:
	@wire gen ./internal/wirex

# go install github.com/swaggo/swag/cmd/swag@latest
swagger:
	@swag init --parseDependency --generalInfo ./main.go --output ./internal/swagger

# https://github.com/OpenAPITools/openapi-generator
openapi:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/internal/swagger/swagger.yaml -g openapi -o /local/internal/swagger/v3

clean:
	rm -rf data $(SERVER_BIN)

serve: build
	./$(SERVER_BIN) start --configdir $(CONFIG_DIR) --config $(CONFIG_FILE) --staticdir $(STATIC_DIR)

serve-d: build
	./$(SERVER_BIN) start --configdir $(CONFIG_DIR) --config $(CONFIG_FILE) --staticdir $(STATIC_DIR) -d

stop:
	./$(SERVER_BIN) stop