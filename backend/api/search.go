package api

import (
	"strconv"

	"github.com/Grvzard/biligank-v2/backend/crud"
	"github.com/gin-gonic/gin"
)

func RegSearch(r *gin.RouterGroup) {
	r.GET("/search", func(ctx *gin.Context) {
		pattern := ctx.Query("q")
		var results []crud.StreamerInfo
		if 1 <= len(pattern) && len(pattern) <= 64 {
			id, err := strconv.ParseInt(pattern, 10, 64)
			if err == nil && id > 0 {
				results = crud.StreamerBy(id)
			}
			// StreamerByUname(pattern)
		}
		if results != nil {
			ctx.JSON(200, results)
		} else {
			ctx.JSON(200, []crud.StreamerInfo{})
		}
	})
}
