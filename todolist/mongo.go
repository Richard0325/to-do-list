package todolist

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoDao struct {
	Coll  *mongo.Collection
	maxID int
}

func (dao MongoDao) GetTasks() (Tasks, error) {
	cur, err := dao.Coll.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("GetTasks error")
		return nil, err
	}
	ret := []*Task{}
	for cur.Next(context.Background()) {
		t:= Task{}
		err := cur.Decode(&t)
		if err != nil {
			fmt.Println("Decode error")
			return nil, err
		}
		t.ID = t.MongoID.Hex()

		ret = append(ret, &t)
	}
	return ret, nil
}
func (dao MongoDao) GetTask(idStr string) (*Task, error) {
	objID, _ := primitive.ObjectIDFromHex(idStr)
	t := Task{}
	err := dao.Coll.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&t)
	if err != nil {
		fmt.Println("GetTask error")
		return nil, ErrNotFound
	}
	t.ID = idStr
	return &t, nil
}
func (dao *MongoDao) AddTask(task *Task) (string, error) {
	opts := options.Count().SetMaxTime(2 * time.Second)
	count, err := dao.Coll.CountDocuments(context.TODO(),bson.D{{}},opts)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("get count error")
		return "-1", err
	}
	if count >= int64(ListSize){
		return "-1", ErrNoSpace
	}
	res, err := dao.Coll.InsertOne(context.TODO(), bson.D{{"title", task.Title}, {"description", task.Description}, {"deadline", task.Deadline}})
	if err != nil {
		return "-1", err
	}
	idStr := res.InsertedID.(primitive.ObjectID).Hex()
	return idStr, nil
}
func (dao MongoDao) DeleteTask(idStr string) error {
	objID, _ := primitive.ObjectIDFromHex(idStr)
	err := dao.Coll.FindOneAndDelete(context.TODO(), bson.D{{"_id", objID}})
	if err != nil {
		fmt.Println("DeleteTask error")
	}
	return nil
}
