package config

type Sql struct {
	MysqlConfig    MysqlConfig    `mapstructure:"mysql-config" json:"mysql-config" yaml:"mysql-config"`
	PostgresConfig PostgresConfig `mapstructure:"postgres-config" json:"postgres-config" yaml:"postgres-config"`
	GormConfig     GormConfig     `mapstructure:"gorm-config" json:"gorm-config" yaml:"gorm-config"`
}

type MysqlConfig struct {
	Hostname string `json:"hostname" yaml:"hostname" mapstructure:"hostname"`
	Port     string `json:"port" yaml:"port" mapstructure:"port"`
	Username string `json:"username" yaml:"username" mapstructure:"username"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	Dbname   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
}
type PostgresConfig struct {
	Hostname string `json:"hostname" yaml:"hostname" mapstructure:"hostname"`
	Port     string `json:"port" yaml:"port" mapstructure:"port"`
	Username string `json:"username" yaml:"username" mapstructure:"username"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	Dbname   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
}

type GormConfig struct {
	SqlType                string `json:"sql-type" yaml:"sql-type" mapstructure:"sql-type"`
	LogMode                string `json:"log-mode" yaml:"log-mode" mapstructure:"log-mode"`
	LogZap                 bool   `json:"log-zap" yaml:"log-zap" mapstructure:"log-zap"`
	SkipDefaultTransaction bool   `json:"skip-default-transaction" yaml:"skip-default-transaction" mapstructure:"skip-default-transaction"`
	TablePrefix            string `json:"table-prefix" yaml:"table-prefix" mapstructure:"table-prefix"`
	SingularTable          bool   `json:"singular-table" yaml:"singular-table" mapstructure:"singular-table"`
	MaxIdleConn            int    `json:"max-idle-conn" yaml:"max-idle-conn" mapstructure:"max-idle-conn"`
	MaxOpenConn            int    `json:"max-open-conn" yaml:"max-open-conn" mapstructure:"max-open-conn"`
}
