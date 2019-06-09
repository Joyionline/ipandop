package model

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// 数据库

// 表结构
type Student struct {
	Name   string `bson:"name"`   // 姓名
	Age    int    `bson:"age"`    // 年龄
	Sid    string `bson:"sid"`    // 学号
	Status int    `bson:"status"` //状态
}

type Per struct {
	Per []Student
}

var (
	client *mgo.Collection
	mongo  *mgo.Session
)

func Init() {
	err := connect()
	if err != nil {
		log.Println("建立数据库连接错误", err)
	}
	// 执行添加
	// InsertData()
	QueryByOne()
	// DeleteData()

}

// 建立连接
func connect() error {
	mongo, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return err
	}
	client = mongo.DB("mydb_tutorial").C("t_student")
	fmt.Println("数据库连接成功")
	return err
}

// 添加数据
func InsertData() error {
	data := Student{
		Name:   "Joyi2",
		Age:    21,
		Sid:    "s19848485",
		Status: 1,
	}
	err := client.Insert(&data)
	if err != nil {
		return nil
	} else {
		fmt.Println("添加数据成功")
	}
	return err
}

func QueryByOne() error {
	mongo, err := mgo.Dial("127.0.0.1")

	defer mongo.Close()
	if err != nil {
		return err
	}
	client := mongo.DB("mydb_tutorial").C("t_student")

	user := Student{}
	users := []Student{}
	// 根据name查询数据
	// cErr := client.Find(bson.M{"name": "Joyi"}).One(&user)
	cErr := client.Find(bson.M{"sid": "s19848485"}).One(&user)
	cErr = client.Find(bson.M{"status": 1}).All(&users)
	// cErr := client.Update()
	if cErr != nil {
		fmt.Println("Error", cErr)
		return err
	}
	fmt.Println("----分割线---")
	fmt.Println("单个数据：", user)
	fmt.Println("所有数据：", users)
	return nil
}

// 更新数据
func UpdateData() error {
	mongo, err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil {
		return err
	}

	client := mongo.DB("mydb_tutorial").C("t_student")

	// 更新一条
	cErr := client.Update(bson.M{"Joyi2": 1}, bson.M{"&set": bson.M{"age": 30}})
	if cErr != nil {
		return cErr
	}
	return nil
}

// 删除数据
func DeleteData() error {
	mongo, err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil {
		return err
	}
	client := mongo.DB("mydb_tutorial").C("t_student")
	//删除一条数据
	cErr := client.Remove(bson.M{"sid": "s19848485"})
	if cErr != nil {
		return cErr
	}
	fmt.Println("数据删除成功")
	return nil
}
