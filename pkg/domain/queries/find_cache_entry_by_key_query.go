package queries

type FindCacheEntryByKeyQuery struct {
	EntryKey   string      `json:"key"`
}

func (*FindCacheEntryByKeyQuery) Key() string {
	return "FindCacheEntryByKeyQuery"
}

