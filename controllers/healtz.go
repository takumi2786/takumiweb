package controllers

import (
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type HealthzController struct{}

func (pc HealthzController) Healthz(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "ok"})
	// NOTE: H is a shortcut for map[string]interface{}
	// type H map[string]any
}
