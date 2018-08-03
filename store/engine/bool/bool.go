package bool

import (
	"github.com/eliastor/allsat-solver/clause"
	"github.com/eliastor/allsat-solver/store"
)

const (
	name = "bool"
)

type boolClause struct {
	data []chunk
	mask []chunk
	tag  clause.Tag
}

type pool struct {
	records []boolClause
}

type adapter struct {
	store.Engine
	pool   pool
	size   uint32
	MSMask chunk
}

type config struct {
	PoolSize int
	Size     int
}

func (a *adapter) open(c interface{}) error {
	conf, ok := c.(config)
	if !ok {
		return store.ErrWrongEngineConfig
	}

	if conf.PoolSize < 1000 {
		conf.PoolSize = 1000
	}

	if conf.PoolSize > 1*1024*1024 {
		//log.Println("Too large PoolSize")
		conf.PoolSize = 1 * 1024 * 1024
	}

	if conf.Size < 1 || conf.Size > 3000 {
		//log.Println("Are serious with size??")
		return store.ErrWrongEngineConfig
	}

	a.size = uint32(conf.Size)

	if mod := a.size % chunkSize; mod != 0 {
		a.size++
		a.MSMask = (1 << mod) - 1
	}

	a.pool.records = make([]boolClause, 0, conf.PoolSize)

	return nil
}
func (a *adapter) close() {

}

func (a *adapter) reload(...interface{}) error {
	return nil
}

func init() {
	store.RegisterEngine(name, &adapter{})
}

func (a *adapter) AddClause(c clause.Clause, tag clause.Tag) (ID clause.ID) {

	return
}
func (a *adapter) GetClause(ID clause.ID) (c clause.Clause, tag clause.Tag) {
	return
}
func (a *adapter) GetClausesByTag(tag clause.Tag) (c []clause.Clause, ID []clause.ID) {
	return
}
func (a *adapter) DelClause(ID clause.ID) {

}

func (a *adapter) DelClausesByTag(tag clause.Tag) {

}

func (a adapter) Int2Clause(c []int) clause.Clause {
	cl := boolClause{}
	cl.data = make([]chunk, a.size)
	cl.mask = make([]chunk, a.size)

	for _, e := range c {
		var d uint
		if e > 0 {
			d = 1
		} else {
			e *= -1
		}
		k := e / chunkSize
		seek := uint(e % chunkSize)

		cl.mask[k] |= chunk(1 << seek)
		cl.data[k] |= chunk(d << seek)

	}

	return &boolClause{}
}

func (c *boolClause) String() string {
	return ""
}
func (c *boolClause) Ints() []int {
	return []int{}
}
func (c *boolClause) Merge(a clause.Clause) clause.Clause {
	return &boolClause{}
}
func (c *boolClause) Type() int {
	return 0
}
func (c *boolClause) Convert2SDNF() clause.Clause {
	return &boolClause{}
}
func (c *boolClause) Tag(clause.Tag) clause.Tag {
	return clause.Tag(0)
}
