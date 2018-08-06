package selective

import (
	"runtime"
	"sync"

	"github.com/eliastor/allsat-solver/clause"
	"github.com/eliastor/allsat-solver/heuristic"
)

const (
	name = "selective"
)

type uid uintptr

type selectiveH struct {
	once sync.Once
	heuristic.Heuristic
	tags         map[clause.Tag]uid
	ids          map[clause.ID]uid
	clauses      map[uid]clause.Clause
	varRating    map[int]uint
	absVarRating map[uint]uint
	clauseRating map[clause.Clause]uint
	varTop       []int
}

func init() {
	heuristic.RegisterHeuristic(name, &selectiveH{})
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (s *selectiveH) Reset() {
	s.tags = make(map[clause.Tag]uid, 100)
	s.ids = make(map[clause.ID]uid, 100)
	s.varRating = make(map[int]uint, 100)
	s.absVarRating = make(map[uint]uint, 50)
	s.clauseRating = make(map[clause.Clause]uint, 100)
	runtime.GC()
}

func (s *selectiveH) calcRatings() {
	for k := range s.varRating {
		s.varRating[k] = 0
	}
	for k := range s.absVarRating {
		s.absVarRating[k] = 0
	}
	for _, c := range s.clauses {
		for _, v := range c.Ints() {
			s.varRating[v]++
			s.absVarRating[uint(abs(v))]++
		}
	}
	s.fillTop()
}

func (s *selectiveH) fillTop() {
	//select top 32 literals

}

func inSlice(needle interface{}, haystack []interface{}) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func (s *selectiveH) calcIDs() {
	for _, c := range s.clauses {
		//c.Mark(clause.ZeroMask)
		//c.Mark()
		var sign int
		switch c.Type() {
		case clause.Conjunction:
			sign = -1
		case clause.SDNF:
			sign = 1
		}
		for _, v := range c.Ints() {
			if inSlice(v, s.varTop) {
				c.Zero()
				c.Vary(1)
			}
		}

	}
}

//Hooks

func (s *selectiveH) Add(c clause.Clause) {

}
func (s *selectiveH) Delete(c clause.Clause) {

}
func (s *selectiveH) Mark(c clause.Clause) {

}
func (s *selectiveH) Merge(a clause.Clause, b clause.Clause, result clause.Clause) {

}
func (s *selectiveH) Select(c clause.Clause) {

}
