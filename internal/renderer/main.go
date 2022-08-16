package renderer

import (
	"test1/templates"

	"github.com/valyala/fasthttp"
)

func (rr *Renderer) Main() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		ds, err := rr.Database.SelectDatas(ctx)
		if err != nil{
			panic(err)
		}

		ctx.Response.Header.Set("Content-Type", "text/html")

		templates.WriteCode(ctx, ds)
	}
	return
}
