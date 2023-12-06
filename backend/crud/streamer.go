package crud

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StreamerInfo struct {
	Uid        uint64 `json:"uid" `
	AreaName   string `json:"area_name" bson:"area_name"`
	ParentName string `json:"parent_name" bson:"parent_name"`
	Roomid     uint64 `json:"roomid"`
	Uname      string `json:"uname"`
}

func streamerByFilter(filter interface{}) []StreamerInfo {
	var results []StreamerInfo
	cursor, err := Coll.Find(context.TODO(), filter)
	if err == nil {
		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
			return nil
		}
		return results
	} else if err == mongo.ErrNoDocuments {
		return nil
	} else {
		log.Fatal(err)
		return nil
	}
}

func StreamerByUid(uid uint64) []StreamerInfo {
	return streamerByFilter(bson.M{"uid": uid})
}

func StreamerByRoomid(roomid uint64) []StreamerInfo {
	return streamerByFilter(bson.M{"roomid": roomid})
}
