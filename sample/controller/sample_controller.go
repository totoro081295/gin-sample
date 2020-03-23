package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SampleController sample controller
type SampleController struct{}

// NewSampleController mount sample controller
func NewSampleController(r *gin.Engine) {
	handler := &SampleController{}
	test := r.Group("/test")
	test.GET("/sample/:name", handler.Get)
}

// Get get
func (c *SampleController) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, name)
}
