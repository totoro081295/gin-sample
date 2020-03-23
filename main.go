package main

import (
	accountC "gin-sample/account/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	accountC.NewAccountController(r)
	r.Run()
}
