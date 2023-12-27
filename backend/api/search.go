package api

import (
	"strconv"

	"github.com/Grvzard/biligank-v2/backend/crud"
	"github.com/gin-gonic/gin"
)

func RegSearch(r *gin.RouterGroup) {
	r.GET("/search", func(ctx *gin.Context) {
		pattern := ctx.Query("q")
		var results = []crud.StreamerInfo{}
		if 1 <= len(pattern) && len(pattern) <= 64 {
			id, err := strconv.ParseInt(pattern, 10, 64)
			if err == nil && id > 0 {
				if tmp := crud.StreamerByUid(id); tmp != nil {
					results = append(results, tmp...)
				}
				if tmp := crud.StreamerByRoomid(id); tmp != nil {
					results = append(results, tmp...)
				}
			}
			// StreamerByUname(pattern)
		}
		ctx.JSON(200, results)
	})
}
