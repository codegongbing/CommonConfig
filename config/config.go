package config

type Server struct {
	Project Project `mapstructure:"project" json:"project" yaml:"project"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu     QiNiu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	//TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
}
