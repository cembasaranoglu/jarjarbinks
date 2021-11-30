package fifo

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"jarjarbinks/pkg/domain"
	"testing"
)

func TestCollection(t *testing.T) {
	entries := []*domain.Entry{}
	entries = append(entries, &domain.Entry{Key: 1})
	entries = append(entries, &domain.Entry{Key: 2})
	entries = append(entries, &domain.Entry{Key: 3})

	c := &collection{ll: list.New()}
	c.Init()

	for _, e := range entries {
		c.Add(e)
	}

	for _, e := range entries {
		for i := 0; i < e.Key.(int); i++ {
			c.Move(e)
		}
	}

	oldest := c.Discard()
	c.Remove(entries[2])
	back := c.ll.Back().Value.(*domain.Entry)

	assert.Equal(t, 1, oldest.Key)
	assert.Equal(t, 1, c.Len())
	assert.Equal(t, 2, back.Key)
}
