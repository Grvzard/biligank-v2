package api

import (
	"time"

	"github.com/Grvzard/biligank-v2/backend/crud"
	"github.com/gin-gonic/gin"
)

type ReqStreamLog struct {
	// timestamp
	From uint32 `form:"from"`
	To   uint32 `form:"to"`
}

func RegStreamlog(r *gin.RouterGroup) {
	r.GET("/streamlog/:uid", func(ctx *gin.Context) {
		var uid uint64
		ctx.BindUri(&uid)
		var req ReqStreamLog
		ctx.BindQuery(&req)
		var results = []crud.StreamlogRow{}
		from := req.From
		to := req.To
		if from == 0 {
			from = uint32(time.Now().Unix())
		}
		if to == 0 {
			to = from
		}
		for to <= from {
			results = append(results, crud.StreamLogByTstamp(from, uid)...)
			from -= 86000
		}
		ctx.JSON(200, results)
	})
}
