package clause

import (
	"math"
)

//Type of stored clase. SDNF stands for special DNF form.
const (
	None int = iota
	Conjunction
	Disjunction
	SDNF
)

/*
Tag is used to mark clause by heuristic and can be usefull for heuristics to select similar clauses
Tag should be used in sense of bit array.
*/
type Tag uint32

//Special Kinds of tags
const (
	EmptyTag = 0
)

//some useful masks for tag
const (
	ZeroMask = Tag(0)
	FullMask = Tag(math.MaxUint32)
)

//ID is same thing as tag but it can be sortable and it should be used as unsigned number
type ID uint32

//Special Kinds of IDs
const (
	EmptyID = 0
)

//Clause interface to be implemented by storage engine
type Clause interface {
	String() string
	Ints() []int
	Merge(Clause) Clause
	Type() int
	Convert2SDNF() Clause
	Tag() Tag
	Mark(mask Tag)
	Unmark(mask Tag)
	ID() ID
	//Vary increase or decrease ID by n
	Vary(n int)
	//Set ID to 0
	Zero()
}
