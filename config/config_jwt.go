package config

type Jwt struct {
	Secret_key                    string `yaml:"secret_key"`
	Expiration_time               uint   `yaml:"expiration_time"`
	Issuer                        string `yaml:"issuer"`
	Refresh_token_expiration_time uint   `yaml:"refresh_token_expiration_time"`
}
