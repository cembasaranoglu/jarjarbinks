package enumeration

type EncodeLevel int

var (
	DefaultLevelEncoder EncodeLevel = 0
	Lowercase EncodeLevel = 1
	Camelcase EncodeLevel = 2
)
