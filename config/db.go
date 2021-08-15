package config

type Mysql struct {
	Hostname       string         `json:"hostname" yaml:"hostname" mapstructure:"hostname"`
	Port           string         `json:"port" yaml:"port" mapstructure:"port"`
	Username       string         `json:"username" yaml:"username" mapstructure:"username"`
	Password       string         `json:"password" yaml:"password" mapstructure:"password"`
	Dbname         string         `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	PersonalConfig PersonalConfig `mapstructure:"personal-config" json:"personal-config" yaml:"personal-config"`
}

type PersonalConfig struct {
	LogMode                string `json:"log-mode" yaml:"log-mode" mapstructure:"log-mode"`
	LogZap                 bool `json:"log-zap" yaml:"log-zap" mapstructure:"log-zap"`
	SkipDefaultTransaction bool   `json:"skip-default-transaction" yaml:"skip-default-transaction" mapstructure:"skip-default-transaction"`
	TablePrefix            string `json:"table-prefix" yaml:"table-prefix" mapstructure:"table-prefix"`
	SingularTable          bool   `json:"singular-table" yaml:"singular-table" mapstructure:"singular-table"`
	MaxIdleConn            int    `json:"max-idle-conn" yaml:"max-idle-conn" mapstructure:"max-idle-conn"`
	MaxOpenConn            int    `json:"max-open-conn" yaml:"max-open-conn" mapstructure:"max-open-conn"`
}
