package interfaces
type ControllerBase interface{
	Name() string
	Prefix() string
	Version() string
}
