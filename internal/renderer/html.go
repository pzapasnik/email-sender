package html

import "github.com/gin-gonic/gin/render"

type HTMLRenderer struct {
}

func NewHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{}
}

func (r *HTMLRenderer) Instance(string, any) render.Render {

}

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}
