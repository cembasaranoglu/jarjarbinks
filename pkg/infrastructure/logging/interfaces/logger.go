package interfaces

type Logger interface{
	Debug(msg string, parameters ...map[string]interface{})
	Info(msg string, parameters ...map[string]interface{})
	Warn(msg string, parameters ...map[string]interface{})
	Error(msg string, err error, parameters ...map[string]interface{})
	Fatal(msg string, err error, parameters ...map[string]interface{})
}

