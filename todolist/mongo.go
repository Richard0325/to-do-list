package todolist

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDao struct {
	DB *mongo.Database
}

func (dao MariaDao) GetTasks() (Tasks, error) {

}
func (dao MariaDao) GetTask(id int) (*Task, error) {

}
func (dao MariaDao) AddTask(task *Task) (int, error) {

}
func (dao MariaDao) DeleteTask(id int) error {

}
