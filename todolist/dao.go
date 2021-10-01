package todolist

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Dao interface {
	GetTasks() (Tasks, error)
	GetTask(string) (*Task, error)
	AddTask(*Task) (string, error)
	DeleteTask(string) error
}

var ListSize int

type DaoType int

var ErrNotFound error = errors.New("not found")
var ErrNoSpace error = errors.New("no space")
var ErrInvalidInput error = errors.New("invalid input")

var MemoryDaoType DaoType = 1
var MongoDaoType DaoType = 2
var MariaDaoType DaoType = 3

func InitDao(daoT DaoType, size int) Dao {
	ListSize = size
	switch daoT {
	case MemoryDaoType:
		return &MemoryDao{
			data:  map[int]*Task{},
			maxID: 0,
		}
	case MariaDaoType:
		conn, err := sql.Open("mysql", "root:bearathome@tcp(minecraft.litttlebear.tw:11100)/Richard")
		if err != nil {
			panic("MariaDB init error")
			return nil
		}
		return &MariaDao{
			DB: conn,
		}
	case MongoDaoType:
		credential := options.Credential{
			Username: "bearathome",
			Password: "bearathome",
		}
		clientOpts := options.Client().ApplyURI("mongodb://minecraft.litttlebear.tw:11101").
			SetAuth(credential)
		client, err := mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			panic("MongoDB init error")
		}
		collection := client.Database("richard").Collection("todolist")
		return &MongoDao{
			Coll: collection,
			maxID : 0,
		}
	}
	return nil
}
