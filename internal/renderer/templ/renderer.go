package renderer

import (
	"context"
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

type TemplRenderer struct {
	fallbackHtmlRenderer render.HTMLRender
}

func NewTemplRenderer(fallbackHtmlRenderer render.HTMLRender) *TemplRenderer {
	return &TemplRenderer{fallbackHtmlRenderer: fallbackHtmlRenderer}
}

func (r *TemplRenderer) Instance(name string, data any) render.Render {
	templData, ok := data.(templ.Component)
	if !ok {
		// if no component is found, fallback to the default HTML renderer
		return r.fallbackHtmlRenderer.Instance(name, data)
	}

	return &Renderer{
		ctx:       context.Background(),
		component: templData,
		name:      name,
		data:      data,
	}
}

type Renderer struct {
	ctx       context.Context
	component templ.Component
	name      string
	data      any
}

func NewRenderer(ctx context.Context, component templ.Component) *Renderer {
	return &Renderer{ctx: ctx, component: component}
}

func (r Renderer) Render(writer http.ResponseWriter) error {
	r.WriteContentType(writer)
	if r.component == nil {
		return errors.New("no component to render")
	}

	return r.component.Render(r.ctx, writer)
}

func (r Renderer) WriteContentType(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
}
