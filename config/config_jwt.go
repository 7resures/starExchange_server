package config

type Jwt struct {
	Secret_key                    string `yaml:"secret_key"`
	Expiration_time               int    `yaml:"expiration_time"`
	Issuer                        string `yaml:"issuer"`
	Refresh_token_expiration_time int    `yaml:"refresh_token_expiration_time"`
}
