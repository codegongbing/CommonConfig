package config

type Project struct {
	Port    string `json:"port" yaml:"port" mapstructure:"port"`
	OSSType string `json:"oss-type" yaml:"oss-type" mapstructure:"oss-type"`
}
