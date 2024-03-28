package main

import (
	"context"
	"fmt"
	"os"

	"github.com/OlivierArgentieri/NukeRlmPrometheusExporter/controllers"
	"github.com/OlivierArgentieri/NukeRlmPrometheusExporter/statics"
	"github.com/joho/godotenv"
)

type program struct {
	LogFile *os.File
	svr     *controllers.Server
	ctx     context.Context
}

func (p *program) Context() context.Context {
	return p.ctx
}

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env file not foubnd, fallback to os.Getenv()")
	}
}

func main() {
	prg := program{
		svr: &controllers.Server{},
	}

	LoadDotEnv()

	prg.svr.Run(fmt.Sprintf(":%s", os.Getenv(statics.PORT)))
}
