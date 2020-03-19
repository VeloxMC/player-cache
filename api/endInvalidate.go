package api

import "github.com/valyala/fasthttp"

// endInvalidateHandler handles the '/invalidate' endpoint
func endInvalidateHandler(ctx *fasthttp.RequestCtx) {
	// Validate the authorization header
	if !validateAuthorization(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBody(errorResponse(errUnauthorized))
		return
	}

	// Invalidate the cache
	cacheInstance.Invalidate()

	// Print a success message
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(successResponse("invalidated cache"))
}
