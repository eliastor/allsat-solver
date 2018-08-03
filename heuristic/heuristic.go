package heuristic

/*
Hooks of heurisitic are used by solver to call heuristic logic at specific steps of solver.
Note that heuristic hooks can be called concurrently
*/
type Hooks interface {
	Add()
	Delete()
	Merge()
	Select()
}

type Config struct {
	Name   string
	Config interface{}
}

type Heuristic struct {
}
