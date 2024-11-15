package config

import "fmt"

type Redis struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Password  string `yaml:"password"`
	Pool_size int    `yaml:"pool_size"`
}

func (r *Redis) GetAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
