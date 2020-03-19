package api

import "github.com/valyala/fasthttp"

// endUUIDHandler handles the '/uuid/:name' endpoint
func endUUIDHandler(ctx *fasthttp.RequestCtx) {
	// Validate the authorization header
	if !validateAuthorization(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBody(errorResponse(errUnauthorized))
		return
	}

	// Get the corresponding UUID
	uuid, err := cacheInstance.GetUUID(ctx.UserValue("name").(string)); if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(errorResponse(err))
		return
	}

	// Print the fetched UUID
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(successResponse(uuid))
}
