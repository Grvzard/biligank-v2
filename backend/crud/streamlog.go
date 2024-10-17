package crud

import (
	"time"

	"github.com/Grvzard/biligank-v2/backend/config"
)

type StreamlogRow struct {
	Title      string `json:"title,omitempty"`
	CTs        uint32 `json:"c_ts,omitempty"`
	LastUpdate uint32 `json:"last_update,omitempty"`
	WatchedNum uint64 `json:"watched_num,omitempty"`
	AreaName   string `json:"area_name,omitempty"`
	Cover      string `json:"cover,omitempty"`
}

func FetchStreamLog(table string, uid int64, to_ *time.Time, from_ *time.Time) []StreamlogRow {
	results := []StreamlogRow{}
	sql := Sqldb.Table(table).Where("uid = ?", uid)
	if to_ != nil {
		sql = sql.Where("last_update >= ?", uint32(to_.Unix()))
	}
	if from_ != nil {
		sql = sql.Where("last_update <= ?", uint32(from_.Unix()))
	}
	sql.Order("_id DESC").Find(&results)
	return results
}

var beijing = time.FixedZone("UTC+8", 8*60*60)

// `from` as near, `to` as far
func StreamLogByTstampSlice(uid int64, from uint32, to uint32) []StreamlogRow {
	results := []StreamlogRow{}
	from_ := time.Unix(int64(from), 0).In(beijing)
	to_ := time.Unix(int64(to), 0).In(beijing)
	step := 1
	// recent data are divided into tables by day (in tables like 2024_02_01)
	// data prior to Stage1 are archived monthly (in tables like 2023_10)
	is_orig_from := true
	// critical values: to_ == from_ -> STOP
	for to_.Before(from_) {
		// critical values: from_ == Stage1 -> step 1
		if from_.Before(config.GlobalConfig.Stage1) {
			step = 2
		}
		if step == 1 {
			tablename := from_.Format("2006_01_02")
			// here we retrieve all entries regardless of the value of from/to
			results = append(results, FetchStreamLog(tablename, uid, nil, nil)...)

			from_ = from_.AddDate(0, 0, -1)
			is_orig_from = false
		} else if step == 2 {
			tablename := from_.Format("2006_01")
			// the search_to here might be smaller than the smallest value in the table,
			// which introduces an overhead, but that's okay
			search_to := &to_
			search_from := &from_
			if !is_orig_from {
				search_from = nil
			}
			results = append(results, FetchStreamLog(tablename, uid, search_to, search_from)...)

			from_ = from_.AddDate(0, -1, 0)
			is_orig_from = false
		}
	}
	return results
}
