package config

type Mysql struct {
	Hostname       string         `json:"hostname" yaml:"hostname" mapstructure:"hostname"`
	Port           string         `json:"port" yaml:"port" mapstructure:"port"`
	Username       string         `json:"username" yaml:"username" mapstructure:"username"`
	Password       string         `json:"password" yaml:"password" mapstructure:"password"`
	Dbname         string         `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	PersonalConfig PersonalConfig `mapstructure:"personal_config"`
}

type PersonalConfig struct {

}
