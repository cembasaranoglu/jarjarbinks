package enumeration

type EncodeCaller int

var (
	DefaultCallerEncoder EncodeCaller = 0
	ShortestFunctionName EncodeCaller = 1
	LongestFunctionName EncodeCaller = 2
)
