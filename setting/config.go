package setting

import "github.com/spf13/viper"

// Config stores all the configuration of the application
// these value can be read using Viper or environment  variables
type Config struct {
	AdressPort      string `mapstructure:"ADDRESS_PORT"`
	GoogleId        string `mapstructure:"GOOGLE_ID"`
	GoogleSecret    string `mapstructure:"GOOGLE_SECRET"`
	FacebookId      string `mapstructure:"FACEBOOK_ID"`
	FacebookSecret  string `mapstructure:"FACEBOOK_SECRET"`
	TwitterId       string `mapstructure:"TWITTER_ID"`
	TwitterSecret   string `mapstructure:"TWITTER_SECRET"`
	Dialect         string `mapstructure:"DIALECT"`
	ConnInfo        string `mapstructure:"CONN_INFO"`
	Sessionkey      string `mapstructure:"SESSION_SECRET"`
	MaxAge          int    `mapstructure:"MAX_AGE"`
	IsProd          bool   `mapstructure:"IS_PROD"`
	SessionPath     string `mapstructure:"SESSION_PATH"`
	SessionHttpOnly bool   `mapstructure:"SESSION_HTTP_ONLY"`
	PasswordPepper  string `mapstructure:"PASSWORD_PEPPER"`
	HashKey         string `mapstructure:"HASH_KEY"`
	DefaultPwd      string `mapstructure:"DEFAULT_PWD"`
}

// LoadConfig reads configuration from env file or env variables
func LoadConfig(path string) (config Config, err error) {
	//general configuration
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	//reading
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
