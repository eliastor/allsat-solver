package clause

const (
	tNone uint = iota
	tConjunction
	tDisjunction
	tSDNF
)

type Tag uint32
type ID uint32

type Clause interface {
	String() string
	Ints() []int
	Merge(Clause) Clause
	Type() int
	Convert2SDNF() Clause
	Tag(Tag) Tag
}
