package store

import (
	"github.com/eliastor/allsat-solver/clause"
)

//Config struct for generic store engine.
type Config struct {
	// Engine name. It'll be key for accessing it by solver
	Name string
	// If engine can't handle concurrent requests, solver will do it by himself (with some lack of speed)
	IsConcurrent bool
	// Engine specific config
	Config interface{}
}

//Engine interface must be implemented by all store engines
type Engine interface {
	open(interface{}) error
	close()
	reload(interface{}) error
	AddClause(c clause.Clause) (uid uintptr, cl clause.Clause)
	GetClauseByID(ID clause.ID) (c clause.Clause)
	GetClausesByTag(tag clause.Tag) (c []clause.Clause, ID []clause.ID)
	DelClause(ID clause.ID)
	DelClausesByTag(tag clause.Tag)
	Int2Clause(c []int) clause.Clause
}

type storeError string

func (s storeError) Error() string {
	return string(s)
}

//Errors
const (
	ErrNoSuchEngine      = storeError("No such store engine")
	ErrTryLater          = storeError("Engine is not responding. Try later")
	ErrWrongEngineConfig = storeError("Wrong engine config")
)

var eng map[string]Engine
var adp Engine

func init() {
	eng = make(map[string]Engine, 3)
}

func checkAndGet(name string) (engine Engine, err error) {
	engine, ok := eng[name]
	if !ok {
		return nil, ErrNoSuchEngine
	}
	return engine, nil
}

func open(name string, c interface{}) error {
	engine, err := checkAndGet(name)
	if err != nil {
		return err
	}
	return engine.open(c)
}

func close(name string) error {
	engine, err := checkAndGet(name)
	if err != nil {
		return err
	}
	engine.close()
	return nil
}

//Reload configuration of storage.
func Reload(name string, c interface{}) error {
	engine, err := checkAndGet(name)
	if err != nil {
		return err
	}
	return engine.reload(c)
}

//RegisterEngine is called by init() function of implemented store engines
func RegisterEngine(name string, engine Engine) {
	eng[name] = engine
}

func GetEngine(name string) (engine Engine, err error) {
	engine, err = checkAndGet(name)
	if err != nil {
		return
	}
	return
}

func SetDefaultEngine(name string) (err error) {
	engine, err := checkAndGet(name)
	if err != nil {
		return
	}
	adp = engine
	return
}

func Default() Engine {
	return adp
}

func AddClause(c clause.Clause) (uintptr, clause.Clause) {
	return adp.AddClause(c)
}
func GetClauseByID(id clause.ID) (c clause.Clause) {
	return adp.GetClauseByID(id)
}

func GetClausesByTag(tag clause.Tag) (c []clause.Clause, ID []clause.ID) {
	return adp.GetClausesByTag(tag)
}

func Int2Clause(c []int) clause.Clause {
	return adp.Int2Clause(c)
}
