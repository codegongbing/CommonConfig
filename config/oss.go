package config

type QiNiu struct {
	AccessKey     string `json:"access_key" yaml:"access_key" mapstructure:"access_key"`
	SecretKey     string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
	BucketName    string `json:"bucket_name" yaml:"bucket_name" mapstructure:"bucket_name"`
	StorageZone   string `json:"storage_zone" yaml:"storage_zone" mapstructure:"storage_zone"`
	UseHTTPS      bool   `json:"use_https" yaml:"use_https" mapstructure:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains" yaml:"use_cdn_domains" mapstructure:"use_cdn_domains"`
	CdnPath       string `json:"cdn_path" yaml:"cdn_path" mapstructure:"cdn_path"`
}

type AliyunOSS struct {
	URL             string `json:"url" yaml:"url" mapstructure:"url"`
	Endpoint        string `json:"endpoint" yaml:"endpoint" mapstructure:"endpoint"`
	AccessKeyId     string `json:"access_key_id" yaml:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret" mapstructure:"access_key_secret"`
	BucketName      string `json:"bucket_name" yaml:"bucket_name" mapstructure:"bucket_name"`
}
