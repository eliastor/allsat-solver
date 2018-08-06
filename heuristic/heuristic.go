package heuristic

import (
	"github.com/eliastor/allsat-solver/clause"
)

var heuristics map[string]Heuristic

/*
Hooks of heurisitic are used by solver to call heuristic logic at specific steps of solver.
Note that heuristic hooks can be called concurrently
*/
type Hooks interface {
	Add(c clause.Clause)
	Delete(c clause.Clause)
	Mark(c clause.Clause)
	Merge(a clause.Clause, b clause.Clause, result clause.Clause)
	Select(c clause.Clause)
}

type Config struct {
	Name   string
	Config interface{}
}

type Heuristic interface {
	Reset()
	Hooks
	IterateTag(f func(c clause.Clause) bool) clause.Clause
	IterateID(f func(c clause.Clause) int) clause.Clause
}

func init() {
	heuristics = make(map[string]Heuristic, 3)
}

func RegisterHeuristic(name string, h Heuristic) {
	heuristics[name] = h
}

/* heuristic tag and id usage:

 */
