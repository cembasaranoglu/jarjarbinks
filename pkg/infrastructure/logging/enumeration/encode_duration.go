package enumeration

type EncodeDuration int

var (
	DefaultDurationEncoder EncodeDuration = 0
	StringDuration EncodeDuration = 1
	MillisecondDuration EncodeDuration = 2
	NanosecondDuration EncodeDuration = 3
)
