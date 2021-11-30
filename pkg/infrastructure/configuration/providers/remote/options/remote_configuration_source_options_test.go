package options

import (
	"testing"
	"time"
)
var defaultDuration, _ = time.ParseDuration("10sec")
func TestValidateEndpointShouldReturnErrorWhenEndpointIsNil(t *testing.T) {
	configurationOptions, err := New("", "", "", defaultDuration)
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFileEndpointMustBeSpecifiedError{
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}

func TestValidateFileTypeShouldReturnErrorWhenFileTypeNil(t *testing.T) {
	configurationOptions, err := New("http://localhost:1122", "", "", defaultDuration)
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFileTypeMustBeSpecifiedError{
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}



func TestValidateFileTypeShouldReturnErrorWhenFileTypeCouldNotBeRecognized(t *testing.T) {
	configurationOptions, err := New("http://localhost:1122", "xml", "", defaultDuration)
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFileTypeCouldNotBeRecognizedError{
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}


func TestValidateFilePathShouldReturnErrorWhenFilePathIsNil(t *testing.T) {
	configurationOptions, err := New("http://localhost:1122", "yaml", "", defaultDuration)
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFilePathMustBeSpecifiedError{
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}


func TestNewConfigurationSourceOptionsShouldReturnConfigurationSourceOptions(t *testing.T) {
	configurationOptions, err := New("sample", "yaml", "sample", defaultDuration)
	if err != nil{
		t.Error("an error does not be expected")
	}

	if configurationOptions == nil{
		t.Error("configuration source must be specified")
	}

	endpoint := configurationOptions.GetEndpoint()
	if endpoint != "sample"{
		t.Error("fileName does not match with specified")
	}

	fileType := configurationOptions.GetFileType()
	if fileType != "yaml"{
		t.Error("fileType does not match with specified")
	}

	filePath := configurationOptions.GetFilePath()
	if filePath != "sample"{
		t.Error("filePath does not match with specified")
	}

	duration := configurationOptions.GetDurationOfRemoteConfigChangeWatcher()
	if duration != defaultDuration{
		t.Error("duration does not match with specified")
	}
}




