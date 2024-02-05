package api

import (
	"strconv"
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
		uid, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
		if err != nil || uid < 1 {
			ctx.String(400, "bad request")
			return
		}
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
		} else if to <= 1672488000 {
			// 2023-01-01T00:00:00+12:00
			ctx.String(400, "bad request")
			return
		}
		for to <= from {
			results = append(results, crud.StreamLogByTstamp(from, uid)...)
			from -= 86400
		}
		ctx.JSON(200, results)
	})
}
