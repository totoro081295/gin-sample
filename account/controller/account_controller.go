package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AccountController account controller
type AccountController struct{}

// NewAccountController mount account controller
func NewAccountController(r *gin.Engine) {
	handler := &AccountController{}
	test := r.Group("/test")
	test.GET("/account/:name", handler.Get)
}

// Get get
func (c *AccountController) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, name)
}
