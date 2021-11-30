package contracts

import "time"

type CacheEntryContract struct {
	Key      string      `json:"key"`
	Value    interface{} `json:"value"`
	ExpireAt time.Time   `json:"expireAt"`
}
