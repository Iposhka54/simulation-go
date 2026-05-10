package entity

type EntityType int

const (
	TypeHerbivore EntityType = iota
	TypePredator
	TypeStatic
)

type Entity interface {
	Type() EntityType
}

type BaseEntity struct {
}

func New() *BaseEntity {
	return &BaseEntity{}
}

func (e *BaseEntity) Type() EntityType {
	panic("need override in the subclasses")
}
