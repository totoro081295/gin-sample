package main

import (
	sampleC "gin-sample/sample/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	sampleC.NewSampleController(r)
	r.Run()
}
