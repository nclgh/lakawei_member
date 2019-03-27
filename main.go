package main

import (
	"github.com/nclgh/lakawei_rpc/server"
	"github.com/nclgh/lakawei_member/model"
)

func initCommon() {
	model.Init()
}

func main() {
	server.Init()
	initCommon()
	server.Run(new(ServiceMember))
}
