package fifo

import (
	"container/list"
	"jarjarbinks/pkg/domain"
	"jarjarbinks/pkg/domain/repository"
)

func init() {
	domain.FIFO.Register(New)
}

func New(cap int) repository.CacheStore {
	col := &collection{list.New()}
	return domain.New(col, cap)
}

type collection struct {
	ll *list.List
}

func (c *collection) Move(e *domain.Entry) {}

func (c *collection) Add(e *domain.Entry) {
	le := c.ll.PushBack(e)
	e.Element = le
}

func (c *collection) Remove(e *domain.Entry) {
	le := e.Element.(*list.Element)
	c.ll.Remove(le)
}

func (c *collection) Discard() (e *domain.Entry) {
	if le := c.ll.Front(); le != nil {
		c.ll.Remove(le)
		e = le.Value.(*domain.Entry)
	}
	return
}

func (c *collection) Len() int {
	return c.ll.Len()
}

func (c *collection) Init() {
	c.ll.Init()
}