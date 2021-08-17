package config

type Project struct {
	Port         string `json:"port" yaml:"port" mapstructure:"port"`
	OSSType      string `json:"oss-type" yaml:"oss-type" mapstructure:"oss-type"`
	RelativePath string `json:"relative-path" yaml:"relative-path" mapstructure:"relative-path"`
}
