package main

import (
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/thiskevinwang/waypoint-plugin-nixpacks/builder"
)

var Options = []sdk.Option{
	sdk.WithComponents(&builder.Builder{}),
	sdk.WithMappers(builder.NixpacksImageMapper),
}

//go:generate protoc -I . --go-grpc_out=. --go_out=. builder/output.proto
func main() {
	sdk.Main(Options...)
}

var (
	_ component.Builder = (*builder.Builder)(nil)
)
