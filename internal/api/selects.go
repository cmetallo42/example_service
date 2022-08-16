package api

import (
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func (api *API) GetDatas() (handler fasthttp.RequestHandler) {
	handler = func(ctx *fasthttp.RequestCtx) {
		datas, err := api.Database.SelectDatas(ctx)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(datas)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}
	}

	return
}
