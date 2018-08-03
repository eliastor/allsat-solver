package main

import (
	"context"
	"errors"
	"log"
	"os"
	"runtime"
	//"github.com/irifrance/gini/dimacs"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"github.com/mitchellh/go-sat/dimacs"
)

type storageConfig struct {
	Name         string
	IsConcurrent bool
	Config       interface{}
}

var config struct {
	Input         string
	Workers       int
	SaveSolutions bool
	ProfileMemory bool
	Heuristics    []heuristicConfig
	Storage       storageConfig
}

var maxWorkers int = (3 * runtime.NumCPU() / 2)

func getConfig() error {
	//TODO: etcd & vault support
	fileFlags := confita.NewLoader(file.NewBackend(os.Args[0] + ".yaml"))
	if fileFlags.Load(context.Background(), &config) == nil {
		return nil
	}
	confFlags := confita.NewLoader(flags.NewBackend())
	if confFlags.Load(context.Background(), &config) == nil {
		return nil
	}
	return errors.New("Unable to load config")
}

func init() {

	log.SetFlags(0) // we don't need any timestamps
}

func main() {
	if err := getConfig(); err != nil {

	}
	if err := checkConfig(); err != nil {

	}

	inputFile, err := os.OpenFile(config.Input, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("Unable to open input file")
	}
	problem, err := dimacs.Parse(inputFile)

	for _, clause := range problem.Formula.Int() {

	}

}
