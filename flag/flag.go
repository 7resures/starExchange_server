package flag

import (
	"flag"
)

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := flag.Bool("db", false, "run in debug mode")
	flag.Parse()
	return Option{*db}
}

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
}
