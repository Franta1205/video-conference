package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CallController struct{}

func NewCallController() *CallController {
	return &CallController{}
}

func (cc *CallController) StartCall(c *gin.Context) {
	c.String(http.StatusOK, "call started")
}

func (cc *CallController) JoinCall(c *gin.Context) {
	c.String(http.StatusOK, "call Joined")
}

func (cc *CallController) EndCall(c *gin.Context) {
	c.String(http.StatusOK, "call ended")
}

func (cc *CallController) CallPage(c *gin.Context) {
	c.File("views/call/show.html")
}
