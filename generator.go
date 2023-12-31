package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const webkitModule = "github.com/diamondburned/gotk4-webkitgtk/pkg"

var preprocessors = []types.Preprocessor{}

var postprocessors = map[string][]girgen.Postprocessor{}

var Data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: webkitModule,
		Packages: []genmain.Package{
			{
				Name:       "javascriptcoregtk-6.0",
				Namespaces: []string{"JavaScriptCore-6"},
			},
			{
				Name:       "webkitgtk-6.0",
				Namespaces: []string{"WebKit-6"},
			},
			{
				Name:       "libsoup-3.0",
				Namespaces: []string{"Soup-3"},
			},
		},
		PkgGenerated: []string{
			"javascriptcore",
			"soup",
			"webkit",
		},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
		},
		Preprocessors:  preprocessors,
		Postprocessors: postprocessors,
	},
)

func main() {
	genmain.Run(Data)
}
