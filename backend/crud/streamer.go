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
	Uid        int64  `json:"uid"`
	AreaName   string `json:"area_name"`
	ParentName string `json:"parent_name"`
	Roomid     int64  `json:"roomid"`
	Uname      string `json:"uname"`
	FirstLogTs uint32 `json:"first_log_ts"`
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
			Uid:        e.Uid,
			AreaName:   e.AreaName,
			ParentName: e.ParentName,
			Roomid:     e.Roomid,
			Uname:      e.Uname,
			FirstLogTs: uint32(e.Id.Timestamp().Unix()),
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
