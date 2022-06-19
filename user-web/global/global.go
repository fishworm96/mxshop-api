package global

import (
	"mxshop-api/user-web/config"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
