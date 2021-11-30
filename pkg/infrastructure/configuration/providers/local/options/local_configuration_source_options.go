package options

import (
	"errors"
	"fmt"
	"jarjarbinks/pkg/infrastructure/configuration/utilities"
	"strings"
)

var (
	validFileTypes                                 = []string{"yaml", "toml", "json", "ini", "env"}
	configurationFileMustBeSpecifiedError          = errors.New("configuration file name must be specified")
	configurationFileTypeMustBeSpecifiedError      = errors.New("configuration file type must be specified")
	configurationFileTypeCouldNotBeRecognizedError = errors.New(fmt.Sprintf("configuration file type could not be recognized. your file type must be one of %s", strings.Join(validFileTypes, ",")))
	configurationFilePathMustBeSpecifiedError      = errors.New("configuration file path must be specified")
)

type ConfigurationSourceOptions struct {
	fileName string
	fileType string
	filePath string
}

func New(fileName string, fileType string, filePath string) (*ConfigurationSourceOptions, error) {
	options := &ConfigurationSourceOptions{
		fileName: fileName,
		fileType: fileType,
		filePath: filePath,
	}
	if err := options.validateFileName(); err != nil {
		return nil, err
	}
	if err := options.validateFileType(); err != nil {
		return nil, err
	}
	if err := options.validateFilePath(); err != nil {
		return nil, err
	}
	return options, nil
}

func (c *ConfigurationSourceOptions) GetFileName() string {
	return c.fileName
}

func (c *ConfigurationSourceOptions) GetFileType() string {
	return c.fileType
}

func (c *ConfigurationSourceOptions) GetFilePath() string {
	return c.filePath
}

func (c *ConfigurationSourceOptions) validateFileName() error {
	if len(c.fileName) == 0 {
		return configurationFileMustBeSpecifiedError
	}
	return nil
}

func (c *ConfigurationSourceOptions) validateFileType() error {
	if len(c.fileType) == 0 {
		return configurationFileTypeMustBeSpecifiedError
	}

	if !utilities.ExistsInSlice(validFileTypes, c.fileType) {
		return configurationFileTypeCouldNotBeRecognizedError
	}

	return nil
}

func (c *ConfigurationSourceOptions) validateFilePath() error {
	if len(c.filePath) == 0 {
		return configurationFilePathMustBeSpecifiedError
	}

	return nil
}
