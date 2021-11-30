package options

import (
	"testing"
)

func TestValidateFileNameShouldReturnErrorWhenFileNameIsNil(t *testing.T) {
	configurationOptions, err := New("", "", "")
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFileMustBeSpecifiedError {
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}

func TestValidateFileTypeShouldReturnErrorWhenFileTypeIsNil(t *testing.T) {
	configurationOptions, err := New("sample", "", "")
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFileTypeMustBeSpecifiedError {
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}


func TestValidateFileTypeShouldReturnErrorWhenFileTypeCouldNotRecognized(t *testing.T) {
	configurationOptions, err := New("sample", "txt", "")
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFileTypeCouldNotBeRecognizedError {
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}



func TestValidateFilePathShouldReturnErrorWhenFilePathIsNil(t *testing.T) {
	configurationOptions, err := New("sample", "yaml", "")
	if err == nil{
		t.Error("an error expected but does not exist")
	}
	if err != configurationFilePathMustBeSpecifiedError {
		t.Error("an error exists but it does not match the expected type")
	}
	if configurationOptions != nil{
		t.Error("configuration source must be nil")
	}
}

func TestNewConfigurationSourceOptionsShouldReturnConfigurationSourceOptions(t *testing.T) {
	configurationOptions, err := New("sample", "yaml", "sample")
	if err != nil{
		t.Error("an error does not be expected")
	}

	if configurationOptions == nil{
		t.Error("configuration source must be specified")
	}

	fileName := configurationOptions.GetFileName()
	if fileName != "sample"{
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
}




