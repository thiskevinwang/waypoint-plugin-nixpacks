PLUGIN_NAME=nixpacks

all: protos build install

protos:
	@echo ""
	@echo "==> Build Protos"

	@protoc -I . --go-grpc_out=. --go_out=. builder/output.proto

build:
	@echo ""
	@echo "==> Compile Plugin"

	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/waypoint-plugin-${PLUGIN_NAME} ./main.go

install:
	@echo ""
	@echo "==> Installing Plugin"

	@rm -f examples/node-express/waypoint-plugin-nixpacks
	@cp -r bin/waypoint-plugin-nixpacks examples/node-express