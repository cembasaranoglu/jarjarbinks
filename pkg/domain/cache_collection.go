package domain

type Collection interface {
	Move(*Entry)
	Add(*Entry)
	Remove(*Entry)
	Discard() *Entry
	Len() int
	Init()
}

