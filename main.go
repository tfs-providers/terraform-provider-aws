// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/provider"
)

// 1. Add this variable definition
var (
	// version is set to the git version via ldflags during build.
	// We default to "dev" for local runs.
	version = "dev"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/tfs-providers/aws",
	}

	// 2. Change provider.New(version) to provider.NewProvider
	// NewProvider matches the function name you defined in internal/provider/provider.go
	// and serves as the factory function required by Serve.
	err := providerserver.Serve(context.Background(), provider.NewProvider, opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
