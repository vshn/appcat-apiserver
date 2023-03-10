//go:build tools
// +build tools

// Package tools is a place to put any tooling dependencies as imports.
// Go modules will be forced to download and install them.
package tools

import (
	// This is basically KubeBuilder
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
	// To generate mocks
	_ "github.com/golang/mock/mockgen"
	// To have protoc generator
	_ "github.com/golang/protobuf/protoc-gen-go"
	// To have Kind updated via Renovate
	_ "sigs.k8s.io/kind"
	// To have protobuf generator
	_ "k8s.io/code-generator"
)
