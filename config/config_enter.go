package config

type Config struct {
	Logger Logger `yaml:"logger"`
	Mysql  Mysql  `yaml:"mysql"`
	System System `yaml:"system"`
	Upload Upload `yaml:"upload"`
	Jwt    Jwt    `yaml:"jwt"`
	Wx     Wx     `yaml:"wx"`
	Redis  Redis  `yaml:"redis"`
}
