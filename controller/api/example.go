package api

import (
	"github.com/gin-gonic/gin"
	"github.com/worryry/8-pigeons/pkg/server/router"
	"net/http"
)

func init() {
	router.Register(&Example{})
}

type Example struct {
}

func (e *Example) NameGet(c *gin.Context) {
	user := []int{1, 23, 4}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "ok", "data": user})
}

func (e *Example) GetNameList(c *gin.Context) {
	user := []int{1, 23, 4}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "ok", "data": user})
}
