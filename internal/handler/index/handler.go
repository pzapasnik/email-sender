package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	renderer "github.com/pzapasnik/email-sender/internal/renderer/html"
	"github.com/pzapasnik/email-sender/web/templates"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (h *Handler) HandleTempl(c *gin.Context) {
	r := renderer.NewRenderer(c.Request.Context(), http.StatusOK, templates.Index())
	c.Render(http.StatusOK, r)
}
