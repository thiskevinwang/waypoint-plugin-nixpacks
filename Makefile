PLUGIN_NAME=nixpacks

all: protos build install

protos:
	@echo ""
	@echo "==> Build Protos"

	@protoc -I . --go-grpc_out=. --go_out=. builder/output.proto

build:
	@echo ""
	@echo "==> Compile Plugin"

	@go build -o ./bin/waypoint-plugin-${PLUGIN_NAME} ./main.go

install:
	@echo ""
	@echo "==> Installing Plugin"

	@cp bin/waypoint-plugin-nixpacks examples/node-express