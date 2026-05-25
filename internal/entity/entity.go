package entity

type Entity interface {
}

type BaseEntity struct {
}

func New() *BaseEntity {
	return &BaseEntity{}
}
