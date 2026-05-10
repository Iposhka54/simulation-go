package entity

type Entity interface {
	String() string
}

type BaseEntity struct {
}

func New() *BaseEntity {
	return &BaseEntity{}
}

func (e *BaseEntity) String() string {
	panic("need override in the subclasses")
}
