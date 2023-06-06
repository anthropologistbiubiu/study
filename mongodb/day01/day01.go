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
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second))
	if err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db = client.Database("my_db")
	fmt.Println(db.Name())
	/*
		//3.选择表 my_collection
		collection = db.Collection("my_collection")
		collection = collection
		fmt.Println(collection)
		//4.插入一条数据
		lr := model.LogRecord{
			JobName: "job10",
			Command: "echo 2",
			Err:     "",
			Content: "2",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		}
		iResult, err := collection.InsertOne(context.TODO(), lr)
		if err != nil {
			fmt.Print(err)
		}
		//_id:默认生成一个全局唯一ID
		id := iResult.InsertedID.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
		// 查找一条数据
		cond := model.FindByJobName{JobName: "job10"}
		cursor, err := collection.Find(context.TODO(), cond, options.Find().SetSkip(0), options.Find().SetLimit(2))
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	filter := bson.M{"jobName": "job10"}
	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(2))
	if err != nil {
		log.Fatal(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	/*
		for cursor.Next(context.TODO()) {
			var lr model.LogRecord
			//反序列化Bson到对象
			if cursor.Decode(&lr) != nil {
				fmt.Print(err)
				return
			}
			//打印结果数据
			fmt.Println(lr)
		}
	*/

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
