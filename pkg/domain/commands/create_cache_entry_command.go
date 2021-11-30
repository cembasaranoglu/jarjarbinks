package commands

import "time"

type CreateCacheEntryCommand struct {
	EntryKey   string        `json:"key"`
	EntryValue interface{}   `json:"value"`
	ExpireAt   time.Duration `json:"expireAt"`
}

func (*CreateCacheEntryCommand) Key() string {
	return "CreateCacheEntryCommand"
}

