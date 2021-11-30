package firestore

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"jarjarbinks/pkg/infrastructure/configuration/interfaces"
	"jarjarbinks/pkg/infrastructure/configuration/providers/remote/options"
	"time"
)
type firestoreConfigurationSource struct {
	remoteConfig *viper.Viper
}

func (c firestoreConfigurationSource) GetValueByKey(key string) interface{} {
	return c.remoteConfig.Get(key)
}

func (c firestoreConfigurationSource) GetIntArrayValueByKey(key string) []int {
	return c.remoteConfig.GetIntSlice(key)
}

func (c firestoreConfigurationSource) GetDurationValueByKey(key string) time.Duration {
	return c.remoteConfig.GetDuration(key)
}

func (c firestoreConfigurationSource) GetStringArrayValueByKey(key string) []string {
	return c.remoteConfig.GetStringSlice(key)
}

func (c firestoreConfigurationSource) GetStringValueByKey(key string) string {
	return c.remoteConfig.GetString(key)
}

func (c firestoreConfigurationSource) GetIntValueByKey(key string) int {
	return c.remoteConfig.GetInt(key)
}

func (c firestoreConfigurationSource) GetInt64ValueByKey(key string) int64 {
	return c.remoteConfig.GetInt64(key)
}

func (c firestoreConfigurationSource) GetFloatValueByKey(key string) float64 {
	return c.remoteConfig.GetFloat64(key)
}

func (c firestoreConfigurationSource) GetBooleanValueByKey(key string) bool {
	return c.remoteConfig.GetBool(key)
}

func (c firestoreConfigurationSource) GetTimeValueByKey(key string) time.Time {
	return c.remoteConfig.GetTime(key)
}

func New(endpoint string, fileType string, filePath string, durationOfRemoteConfigChangeWatcher time.Duration) (interfaces.ConfigurationSource, error) {
	if option, err := options.New(endpoint, fileType, filePath, durationOfRemoteConfigChangeWatcher); err != nil {
		return nil, err
	} else {
		remoteConfig := viper.New()
		err := remoteConfig.AddRemoteProvider("firestore", option.GetEndpoint(), option.GetFilePath())
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

		return &firestoreConfigurationSource{remoteConfig: remoteConfig}, nil
	}
}
