package config

type Local struct {
	Path     string `json:"path" yaml:"path" mapstructure:"path"`
	DatePath bool   `json:"date-path" yaml:"date-path" mapstructure:"date-path"`
}

type QiNiu struct {
	AccessKey     string `json:"access-key" yaml:"access-key" mapstructure:"access-key"`
	SecretKey     string `json:"secret-key" yaml:"secret-key" mapstructure:"secret-key"`
	BucketName    string `json:"bucket-name" yaml:"bucket-name" mapstructure:"bucket-name"`
	StorageZone   string `json:"storage-zone" yaml:"storage-zone" mapstructure:"storage-zone"`
	UseHTTPS      bool   `json:"use-https" yaml:"use-https" mapstructure:"use-https"`
	UseCdnDomains bool   `json:"use-cdn-domains" yaml:"use-cdn-domains" mapstructure:"use-cdn-domains"`
	CdnPath       string `json:"cdn-path" yaml:"cdn-path" mapstructure:"cdn-path"`
	DatePath      bool   `json:"date-path" yaml:"date-path" mapstructure:"date-path"`
}

type AliyunOSS struct {
	URL             string `json:"url" yaml:"url" mapstructure:"url"`
	Endpoint        string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
	AccessKeyId     string `json:"access-key-id" yaml:"access-key-id" mapstructure:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret" yaml:"access-key-secret" mapstructure:"access-key-secret"`
	BucketName      string `json:"bucket-name" yaml:"bucket-name" mapstructure:"bucket-name"`
	DatePath        bool   `json:"date-path" yaml:"date-path" mapstructure:"date-path"`
}
