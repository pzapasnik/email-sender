package index

import (
	"log/slog"
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
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func (h *Handler) HandleTempl(c *gin.Context) {
	slog.Info("HandleTempl")
	c.HTML(http.StatusOK, "index.templ", templates.Index())
	// r := renderer.NewRenderer(c.Request.Context(), http.StatusOK, templates.Index())
	// _ = r.Render(c.Writer)
}
