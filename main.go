package main

import (
	"red_envelop_server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.LoadSnatch(r)
	routers.LoadOpen(r)
	routers.LoadWalletList(r)
	r.Run()
}
