package todolist

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		t := Task{}
		err := cur.Decode(&t)
		if err != nil {
			fmt.Println("Decode error")
			return nil, err
		}
		ret = append(ret, &t)
	}
	return ret, nil
}
func (dao MongoDao) GetTask(id int) (*Task, error) {
	t := Task{}
	err := dao.Coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&t)
	if err != nil {
		fmt.Println("GetTask error")
		return nil, ErrNotFound
	}
	return &t, nil
}
func (dao *MongoDao) AddTask(task *Task) (int, error) {
	_, err := dao.Coll.InsertOne(context.TODO(), bson.D{{"id", dao.maxID}, {"title", task.Title}, {"description", task.Description}, {"deadline", task.Deadline}})
	if err != nil {
		return -1, err
	}
	id := dao.maxID
	dao.maxID++
	return id, nil
}
func (dao MongoDao) DeleteTask(id int) error {
	err := dao.Coll.FindOneAndDelete(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		fmt.Println("DeleteTask error")
	}
	return nil
}
