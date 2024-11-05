package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CallController struct{}

func NewCallController() *CallController {
	return &CallController{}
}

func (vcc *CallController) StartCall(c *gin.Context) {
	c.String(http.StatusOK, "call started")
}
