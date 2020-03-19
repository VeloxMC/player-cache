package api

import "github.com/valyala/fasthttp"

// endNameHandler handles the '/name/:uuid' endpoint
func endNameHandler(ctx *fasthttp.RequestCtx) {
	// Validate the authorization header
	if !validateAuthorization(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBody(errorResponse(errUnauthorized))
		return
	}

	// Get the corresponding name
	name, err := cacheInstance.GetName(ctx.UserValue("uuid").(string)); if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(errorResponse(err))
		return
	}

	// Print the fetched name
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(successResponse(name))
}

