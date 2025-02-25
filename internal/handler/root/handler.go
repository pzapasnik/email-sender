package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pzapasnik/email-sender/web/templates"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "root.templ", templates.Root())
}
