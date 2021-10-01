package todolist

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	MongoID     primitive.ObjectID `json:"-" bson:"_id"`
	ID          string `json:"id" bson:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

type Tasks []*Task
