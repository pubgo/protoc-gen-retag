// Note: 本项目主要思路和代码来源于protoc-gen-go-tag
// https://github.com/searKing/golang/tree/master/tools/protoc-gen-go-tag

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pubgo/protoc-gen-retag/ast"
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

// Build info, injected at build time via -ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	showVersion := flag.Bool("version", false, "print version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("protoc-gen-retag %s (commit: %s, built at: %s)\n", version, commit, date)
		os.Exit(0)
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.
		Run(func(gen *protogen.Plugin) error {
			gen.SupportedFeatures = gengo.SupportedFeatures
			var originFiles []*protogen.GeneratedFile

			for _, f := range gen.Files {
				if !f.Generate {
					continue
				}

				originFiles = append(originFiles, gengo.GenerateFile(gen, f))
			}

			ast.Rewrite(gen)

			for _, f := range originFiles {
				f.Skip()
			}
			return nil
		})
}
