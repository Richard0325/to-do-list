package todolist

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
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
type MemoryDao struct {
	data  map[int]*Task
	maxID int
}

func (m MemoryDao) GetTasks() (Tasks, error) {
	ret := []*Task{}
	for _, task := range m.data {
		ret = append(ret, task)
	}
	return ret, nil
}

func (m MemoryDao) GetTask(idStr string) (*Task, error) {
	id , _ :=strconv.Atoi(idStr)
	if task, ok := m.data[id]; ok {
		return task, nil
	}
	return nil, ErrNotFound
}

const maxInt = int(^uint(0) >> 1)

func (m *MemoryDao) AddTask(task *Task) (string, error) {
	if len(m.data) == ListSize {
		return "-1", ErrNoSpace
	}
	id := m.maxID
	if m.maxID == maxInt {
		//ToDo
		m.maxID = ListSize
	} else {
		m.maxID++
	}
	m.data[id] = task
	idStr:= strconv.Itoa(id)
	return idStr, nil
}

func (m MemoryDao) DeleteTask(idStr string) error {
	id, _:= strconv.Atoi(idStr)
	delete(m.data, id)
	return nil
}
