package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ModelOption  map[string]interface{}
	ModelPath    string
	UserName     string
	SystemPrompt string
}

var C Config
