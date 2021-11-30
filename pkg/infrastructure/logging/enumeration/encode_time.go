package enumeration

type EncodeTime int

var (
	DefaultTimeEncoder EncodeTime = 0
	RFC3339Nano EncodeTime = 1
	RFC3339 EncodeTime = 2
	ISO8601 EncodeTime = 3
	Milliseconds EncodeTime = 4
	Nanoseconds EncodeTime = 5
)