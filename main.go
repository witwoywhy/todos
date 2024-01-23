package main

import (
	"todos/httpserv"
	"todos/infra"
	"todos/utils/validate"
)

func init() {
	infra.InitConfig()
}

func main() {
	validate.InitValidate()
	infra.InitFileStorage()
	httpserv.Run()
}
