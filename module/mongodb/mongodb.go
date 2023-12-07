package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Person struct {
	Name  string
	Age   int
	City  string
	Phone string
}

func WirteData() {

	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI("mongodb://192.168.60.152")

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 获取数据库和集合的引用
	database := client.Database("mydatabase")
	collection := database.Collection("people")

	// 插入文档
	person := Person{
		Name:  "Timog",
		Age:   30,
		City:  "Ortigas City",
		Phone: "1123-7890",
	}

	insertResult, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted document with ID:", insertResult.InsertedID)

	// 查询文档
	var result Person
	err = collection.FindOne(context.TODO(), nil).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found document: %+v\n", result)

	// 关闭连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection closed.")
}

func FindOne() {

	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI("mongodb://192.168.60.152")

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 获取数据库和集合的引用
	database := client.Database("mydatabase")
	collection := database.Collection("people")

	// 构建正则表达式对象
	filter := bson.D{{"age", 30}}

	// 执行查询
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// 查询文档
	var results []Person
	for cursor.Next(context.TODO()) {
		var result Person
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	// 检查是否有匹配的文档
	if len(results) == 0 {
		fmt.Println("No matching documents found.")
	} else {
		fmt.Printf("Found documents: %+v\n", results)
	}

	// 关闭连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection closed.")
}
