package crud

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StreamerInfoDb struct {
	Uid        int64
	AreaName   string `bson:"area_name"`
	ParentName string `bson:"parent_name"`
	Roomid     int64
	Uname      string
	Id         primitive.ObjectID `bson:"_id"`
}

type StreamerInfo struct {
	Uid         int64  `json:"uid"`
	AreaName    string `json:"area_name"`
	ParentName  string `json:"parent_name"`
	Roomid      int64  `json:"roomid"`
	ShortRoomid int64  `json:"short_roomid"`
	Uname       string `json:"uname"`
	FirstLogTs  uint32 `json:"first_log_ts"`
}

func streamerByFilter(filter interface{}) []StreamerInfo {
	var results_db []StreamerInfoDb
	var results []StreamerInfo

	cursor, err := Coll.Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			log.Fatal(err)
			return nil
		}
	}
	if err = cursor.All(context.TODO(), &results_db); err != nil {
		log.Fatal(err)
		return nil
	}
	for _, e := range results_db {
		results = append(results, StreamerInfo{
			Uid:         e.Uid,
			AreaName:    e.AreaName,
			ParentName:  e.ParentName,
			Roomid:      e.Roomid,
			ShortRoomid: 0,
			Uname:       e.Uname,
			FirstLogTs:  uint32(e.Id.Timestamp().Unix()),
		})
	}
	return results
}

func StreamerByUid(uid int64) []StreamerInfo {
	return streamerByFilter(bson.M{"uid": uid})
}

func StreamerByRoomid(roomid int64) []StreamerInfo {
	return streamerByFilter(bson.M{"roomid": roomid})
}

func StreamerBy(id int64) []StreamerInfo {
	var results []StreamerInfo
	if id <= 1000 {
		results = StreamerByUid(id)
		origin_roomid := ShortRoomid[id]
		// 0 means none, 1 ~ 1000 are reserved
		if origin_roomid > 1000 {
			if tmp := StreamerByRoomid(origin_roomid); tmp != nil {
				// should be only 1 element
				for _, e := range tmp {
					e.ShortRoomid = id
					results = append(results, e)
				}
			}
		}
	} else {
		results = streamerByFilter(bson.M{
			"$or": []bson.M{
				{"uid": id},
				{"roomid": id},
			},
		})
	}
	return results
}
