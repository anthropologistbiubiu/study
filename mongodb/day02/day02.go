package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongodb/model"
	"time"
)

func main() {
	var (
		collection *mongo.Collection
		err        error
		cursor     *mongo.Cursor
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second))
	if err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db := client.Database("my_db")

	//3.选择表 my_collection
	collection = db.Collection("my_collection")

	// 这一段分组求和其实可以使用某个方法做个封装
	groupStage := []model.Group{}
	groupStage = append(groupStage, model.Group{
		Group: bson.D{
			{"_id", "$jobName"},
			{"countJob", model.Sum{Sum: 1}},
		},
	})

	if cursor, err = collection.Aggregate(context.TODO(), groupStage); err != nil {
		log.Fatal(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	//遍历游标
	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
