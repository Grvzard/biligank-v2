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
		from := req.From
		to := req.To

		ts_now := uint32(time.Now().Unix())
		if from == 0 {
			from = ts_now
		} else {
			from = min(from, ts_now)
		}
		// 2023-01-01T00:00:00+8:00
		to = max(to, 1672502400)

		ctx.JSON(200, crud.StreamLogByTstampSlice(uid, from, to))
	})
}
