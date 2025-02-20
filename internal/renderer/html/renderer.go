package renderer

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

type HTMLRenderer struct {
	fallbackHtmlRenderer render.HTMLRender
}

func NewHTMLRenderer(fallbackHtmlRenderer render.HTMLRender) *HTMLRenderer {
	return &HTMLRenderer{fallbackHtmlRenderer: fallbackHtmlRenderer}
}

func (r *HTMLRenderer) Instance(name string, data any) render.Render {
	templData, ok := data.(templ.Component)
	if !ok {
		return r.fallbackHtmlRenderer.Instance(name, data)
	}

	return &Renderer{
		ctx:       context.Background(),
		status:    -1,
		component: templData,
	}
}

type Renderer struct {
	ctx       context.Context
	status    int
	component templ.Component
}

func NewRenderer(ctx context.Context, status int, component templ.Component) *Renderer {
	return &Renderer{}
}

func (r Renderer) Render(writer http.ResponseWriter) error {
	r.WriteContentType(writer)
	if r.status != -1 {
		writer.WriteHeader(r.status)
	}

	writer.WriteHeader(r.status)

	if r.component == nil {
		return nil
	}

	return r.component.Render(r.ctx, writer)
}

func (r Renderer) WriteContentType(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
}
