package local

import (
	"github.com/spf13/viper"
	"jarjarbinks/pkg/infrastructure/configuration/interfaces"
	"jarjarbinks/pkg/infrastructure/configuration/providers/local/options"
	"time"
)

type localConfigurationSource struct {
	config *viper.Viper
}

func (l localConfigurationSource) GetValueByKey(key string) interface{} {
	return l.config.Get(key)
}

func (l localConfigurationSource) GetIntArrayValueByKey(key string) []int {
	return l.config.GetIntSlice(key)
}

func (l localConfigurationSource) GetDurationValueByKey(key string) time.Duration {
	return l.config.GetDuration(key)
}

func (l localConfigurationSource) GetStringArrayValueByKey(key string) []string {
	return l.config.GetStringSlice(key)
}

func (l localConfigurationSource) GetStringValueByKey(key string) string {
	return l.config.GetString(key)
}

func (l localConfigurationSource) GetIntValueByKey(key string) int {
	return l.config.GetInt(key)
}

func (l localConfigurationSource) GetInt64ValueByKey(key string) int64 {
	return l.config.GetInt64(key)
}

func (l localConfigurationSource) GetFloatValueByKey(key string) float64 {
	return l.config.GetFloat64(key)
}

func (l localConfigurationSource) GetBooleanValueByKey(key string) bool {
	return l.config.GetBool(key)
}

func (l localConfigurationSource) GetTimeValueByKey(key string) time.Time {
	return l.config.GetTime(key)
}

func New(fileName string, fileType string, filePath string) (interfaces.ConfigurationSource, error) {
	if option, err := options.New(fileName, fileType, filePath); err != nil {
		return nil, err
	} else {
		config := viper.New()
		config.SetConfigName(option.GetFileName())
		config.SetConfigType(option.GetFileType())
		config.AddConfigPath(option.GetFilePath())
		err := config.ReadInConfig()
		if err != nil {
			return nil, err
		}
		return &localConfigurationSource{
			config: config,
		}, nil
	}
}
