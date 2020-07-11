// Copyright 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

// protoc-gen-go-jsonshim is a plugin for the Google protocol buffer compiler to
// generate Go code which generates MarshalJSON() and UnmarshalJSON() functions from the Protocol Buffers definition types.
//
// Install it by building this program and making it accessible within your PATH with the name:
//	protoc-gen-go-jsonshim
//
// The 'go-jsonshim' suffix becomes part of the argument for the protocol compiler,
// such that it can be invoked as:
//	protoc --go-jsonshim_out=. path/to/file.proto
//
// This generates Go service definitions for the protocol buffer defined by file.proto.
//
// With that input, the output will be written to:
//	path/to/file_jsonshim.pb.go
package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"go.lsp.dev/tools/cmd/protoc-gen-go-jsonshim/jsonshim"
)

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			jsonshim.GenerateFile(gen, f)
		}
		gen.SupportedFeatures = jsonshim.SupportedFeatures
		return nil
	})
}
