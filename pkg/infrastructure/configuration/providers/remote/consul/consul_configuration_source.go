package consul

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"jarjarbinks/pkg/infrastructure/configuration/interfaces"
	"jarjarbinks/pkg/infrastructure/configuration/providers/remote/options"
	"time"
)

type consulConfigurationSource struct {
	remoteConfig *viper.Viper
}

func (c consulConfigurationSource) GetValueByKey(key string) interface{} {
	return c.remoteConfig.Get(key)
}

func (c consulConfigurationSource) GetIntArrayValueByKey(key string) []int {
	return c.remoteConfig.GetIntSlice(key)
}

func (c consulConfigurationSource) GetDurationValueByKey(key string) time.Duration {
	return c.remoteConfig.GetDuration(key)
}

func (c consulConfigurationSource) GetStringArrayValueByKey(key string) []string {
	return c.remoteConfig.GetStringSlice(key)
}

func (c consulConfigurationSource) GetStringValueByKey(key string) string {
	return c.remoteConfig.GetString(key)
}

func (c consulConfigurationSource) GetIntValueByKey(key string) int {
	return c.remoteConfig.GetInt(key)
}

func (c consulConfigurationSource) GetInt64ValueByKey(key string) int64 {
	return c.remoteConfig.GetInt64(key)
}

func (c consulConfigurationSource) GetFloatValueByKey(key string) float64 {
	return c.remoteConfig.GetFloat64(key)
}

func (c consulConfigurationSource) GetBooleanValueByKey(key string) bool {
	return c.remoteConfig.GetBool(key)
}

func (c consulConfigurationSource) GetTimeValueByKey(key string) time.Time {
	return c.remoteConfig.GetTime(key)
}

func New(endpoint string, fileType string, filePath string, durationOfRemoteConfigChangeWatcher time.Duration) (interfaces.ConfigurationSource, error) {
	if option, err := options.New(endpoint, fileType, filePath, durationOfRemoteConfigChangeWatcher); err != nil {
		return nil, err
	} else {
		remoteConfig := viper.New()
		err := remoteConfig.AddRemoteProvider("consul", option.GetEndpoint(), option.GetFilePath())
		remoteConfig.SetConfigType(option.GetFileType())
		if err != nil {
			return nil, err
		}
		err = remoteConfig.ReadRemoteConfig()
		if err != nil {
			return nil, err
		}

		go func() {
			for {
				time.Sleep(option.GetDurationOfRemoteConfigChangeWatcher())
				err = remoteConfig.WatchRemoteConfig()
				if err != nil {
					continue
				}
			}
		}()
		return &consulConfigurationSource{remoteConfig: remoteConfig}, nil
	}

}
