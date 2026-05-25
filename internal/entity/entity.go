package entity

import "sync/atomic"

type Entity interface {
	ID() uint64
}

type BaseEntity struct {
	id uint64
}

var entitySeq uint64

func New() *BaseEntity {
	return &BaseEntity{id: atomic.AddUint64(&entitySeq, 1)}
}

func (e *BaseEntity) ID() uint64 {
	return e.id
}
