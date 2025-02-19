package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
