package interfaces
type JsonMarshaller interface {
	Marshall(v interface{}) ([]byte, error)
	Unmarshall(data []byte, v interface{}) error
	MarshallString(v interface{}) (string, error)
	UnmarshallString(data string, v interface{}) error
}
