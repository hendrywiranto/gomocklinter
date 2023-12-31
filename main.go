package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/hendrywiranto/gomocklinter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.New())
}
