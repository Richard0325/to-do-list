package todolist

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Dao interface {
	GetTasks() (Tasks, error)
	GetTask(int) (*Task, error)
	AddTask(*Task) (int, error)
	DeleteTask(int) error
}

var ListSize int

type MemoryDao struct {
	data  map[int]*Task
	maxID int
}

type DaoType int

var ErrNotFound error = errors.New("not found")
var ErrNoSpace error = errors.New("no space")
var ErrInvalidInput error = errors.New("invalid input")

var MemoryDaoType DaoType = 1
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
			fmt.Printf("SQL init fail %s", err.Error())
			return nil
		}
		return &MariaDao{
			DB: conn,
		}
	}
	case MongoDaoType:
		// credential := options.Credential{
		// 	Username: "user",
		// 	Password: "password",
		// }
		// clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").
		// 	SetAuth(credential)
		// client, err := mongo.Connect(context.TODO(), clientOpts)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// _ = client
	return nil
}

func (m MemoryDao) GetTasks() (Tasks, error) {
	ret := []*Task{}
	for _, task := range m.data {
		ret = append(ret, task)
	}
	return ret, nil
}

func (m MemoryDao) GetTask(id int) (*Task, error) {
	if task, ok := m.data[id]; ok {
		return task, nil
	}
	return nil, ErrNotFound
}

const maxInt = int(^uint(0) >> 1)

func (m *MemoryDao) AddTask(task *Task) (int, error) {
	if len(m.data) == ListSize {
		return -1, ErrNoSpace
	}
	id := m.maxID
	if m.maxID == maxInt {
		//ToDo
		m.maxID = ListSize
	} else {
		m.maxID++
	}
	m.data[id] = task
	return id, nil
}

func (m MemoryDao) DeleteTask(id int) error {
	delete(m.data, id)
	return nil
}
