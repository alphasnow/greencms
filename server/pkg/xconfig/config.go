// @author AlphaSnow

package xconfig

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"server/pkg/utils"
	"strings"
)

const ENV_PREFIX = "CNF"

// ReadConfig
func ReadConfig(configFile string) (cnf *viper.Viper, err error) {
	cnf = viper.New()

	cnf.SetConfigFile(configFile)
	//cnf.AddConfigPath("./configs")
	//cnf.SetConfigName("app")
	//cnf.SetConfigType("yaml")

	err = cnf.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// cnf.WatchConfig()

	return cnf, nil
}

// ReadConfigAndEnv
// env 修改后不会实时更新,需要重启应用, 建议使用ReadConfig
func ReadConfigAndEnv(configFile string, envFile string) (cnf *viper.Viper, err error) {
	if cnf, err = ReadConfig(configFile); err != nil {
		return
	}

	if utils.IsExist(envFile) {
		if err = godotenv.Load(envFile); err != nil {
			return
		}
	}

	cnf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cnf.SetEnvPrefix(ENV_PREFIX)
	cnf.AutomaticEnv()

	return
}
