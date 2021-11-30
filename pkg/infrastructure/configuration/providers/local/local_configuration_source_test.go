package local

import (
	"testing"
)

func TestNewWithOptionsShouldReturnErrorWhenFileDoesNotExistsInDefinedPath(t *testing.T) {
	configurationSource, err := New("sample", "yaml", ".")
	if err == nil{
		t.Error("an error expected but does not exist")
	}

	if configurationSource != nil {
		t.Error("configuration must be nil")
	}

}
