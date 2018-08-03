package adapter

import "github.com/eliastor/allsat-solver/clause"

type Engine interface {
	open(...interface{}) error
	close()
	reload(...interface{}) error
	AddClause(c clause.Clause, tag clause.Tag) (ID clause.ID)
	GetClause(ID clause.ID) (c clause.Clause, tag clause.Tag)
	GetClausesByTag(tag clause.Tag) (c []clause.Clause, ID []clause.ID)
	DelClause(ID clause.ID)
	DelClausesByTag(tag clause.Tag)
	Int2Clause(c []int) clause.Clause
}
