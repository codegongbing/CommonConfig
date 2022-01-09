package config

type Server struct {
	Project   Project   `mapstructure:"project" json:"project" yaml:"project"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Sql       Sql       `mapstructure:"sql" json:"sql" yaml:"sql"`
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu     QiNiu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	//TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
}
