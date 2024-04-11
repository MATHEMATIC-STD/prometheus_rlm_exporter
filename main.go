package main

import (
	"context"
	"fmt"
	"os"

	"github.com/OlivierArgentieri/PrometheusExporter/controllers"
	"github.com/OlivierArgentieri/PrometheusExporter/statics"
)

type program struct {
	LogFile *os.File
	svr     *controllers.Server
	ctx     context.Context
}

func (p *program) Context() context.Context {
	return p.ctx
}

func main() {
	prg := program{
		svr: &controllers.Server{},
	}

	prg.svr.Run(fmt.Sprintf(":%s", os.Getenv(statics.PORT)))
}
