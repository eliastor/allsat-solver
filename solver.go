package solver

import (
	"errors"
	"runtime"
	//"github.com/irifrance/gini/dimacs"
	"github.com/eliastor/allsat-solver/clause"
	"github.com/eliastor/allsat-solver/heuristic"
	"github.com/eliastor/allsat-solver/store"
)

type Config struct {
	Input         string
	Workers       int
	SaveSolutions bool
	ProfileMemory bool
	Heuristics    []heuristic.Config
	Storage       store.Config
}

var config Config

var maxWorkers int = (3 * runtime.NumCPU() / 2)

func init() {

}

func checkConfig() (err error) {
	if config.Input == "" {
		return errors.New("No input file specified")
	}
	if config.Workers < 1 {
		config.Workers = 1
	}
	if config.Workers > maxWorkers {
		config.Workers = maxWorkers
	}
	//heuristics
	//storage
	return
}

func parseProblem(p [][]int) (err error) {
	for _, c := range p {
		cl := store.Int2Clause(c)
		
		store.AddClause(, clause.EmptyTag)
	}
	return
}

func Run(c Config, p [][]int) (err error) {
	config = c
	err = checkConfig()
	if err != nil {
		return
	}
	err = parseProblem(p)
	if err != nil {
		return
	}
	return
}
