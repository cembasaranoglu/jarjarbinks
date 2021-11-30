package options

import (
	"errors"
	"fmt"
	"jarjarbinks/pkg/infrastructure/configuration/utilities"
	"strings"
	"time"
)

var (
	validFileTypes                                 = []string{"yaml", "toml", "json", "ini", "env"}
	configurationFileEndpointMustBeSpecifiedError  = errors.New("configuration file endpoint must be specified")
	configurationFileTypeMustBeSpecifiedError      = errors.New("configuration file type must be specified")
	configurationFileTypeCouldNotBeRecognizedError = errors.New(fmt.Sprintf("configuration file type could not be recognized. your file type must be one of %s", strings.Join(validFileTypes, ",")))
	configurationFilePathMustBeSpecifiedError      = errors.New("configuration file path must be specified")
)

type RemoteConfigurationSourceOptions struct {
	durationOfRemoteConfigChangeWatcher time.Duration
	endpoint                            string
	fileType                            string
	filePath                            string
}

func New(endpoint string, fileType string, filePath string, durationOfRemoteConfigChangeWatcher time.Duration) (*RemoteConfigurationSourceOptions, error) {
	configurationOptions := &RemoteConfigurationSourceOptions{
		endpoint:                            endpoint,
		fileType:                            fileType,
		filePath:                            filePath,
		durationOfRemoteConfigChangeWatcher: durationOfRemoteConfigChangeWatcher,
	}
	if err := configurationOptions.validateEndpoint(); err != nil {
		return nil, err
	}
	if err := configurationOptions.validateFileType(); err != nil {
		return nil, err
	}
	if err := configurationOptions.validateFilePath(); err != nil {
		return nil, err
	}
	return configurationOptions, nil
}

func (c *RemoteConfigurationSourceOptions) GetEndpoint() string {
	return c.endpoint
}

func (c *RemoteConfigurationSourceOptions) GetFileType() string {
	return c.fileType
}

func (c *RemoteConfigurationSourceOptions) GetFilePath() string {
	return c.filePath
}

func (c *RemoteConfigurationSourceOptions) GetDurationOfRemoteConfigChangeWatcher() time.Duration {
	return c.durationOfRemoteConfigChangeWatcher
}

func (c *RemoteConfigurationSourceOptions) validateEndpoint() error {
	if len(c.endpoint) == 0 {
		return configurationFileEndpointMustBeSpecifiedError
	}
	return nil
}

func (c *RemoteConfigurationSourceOptions) validateFileType() error {
	if len(c.fileType) == 0 {
		return configurationFileTypeMustBeSpecifiedError
	}

	if !utilities.ExistsInSlice(validFileTypes, c.fileType) {
		return configurationFileTypeCouldNotBeRecognizedError
	}

	return nil
}

func (c *RemoteConfigurationSourceOptions) validateFilePath() error {
	if len(c.filePath) == 0 {
		return configurationFilePathMustBeSpecifiedError
	}

	return nil
}
