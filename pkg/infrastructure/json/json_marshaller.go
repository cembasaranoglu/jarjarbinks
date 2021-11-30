package json

import (
	jsoniter "github.com/json-iterator/go"
	"jarjarbinks/pkg/infrastructure/json/interfaces"
)

type defaultJsonMarshaller struct {
	jsoniter jsoniter.API
}

func (j *defaultJsonMarshaller) Marshall(v interface{}) ([]byte, error) {
	byteArray, err := j.jsoniter.Marshal(v)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

func (j *defaultJsonMarshaller) Unmarshall(data []byte, v interface{}) error {
	err := j.jsoniter.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func (j *defaultJsonMarshaller) MarshallString(v interface{}) (string, error) {
	byteArray, err := j.jsoniter.MarshalToString(v)
	if err != nil {
		return "", err
	}
	return byteArray, nil
}

func (j *defaultJsonMarshaller) UnmarshallString(data string, v interface{}) error {
	err := j.jsoniter.UnmarshalFromString(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func New() interfaces.JsonMarshaller {
	return &defaultJsonMarshaller{
		jsoniter: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}
