package todolist

import "strconv"

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
	id ,err :=strconv.Atoi(idStr)
	if err != nil{
		return nil, ErrInvalidInput
	}
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
	id, err:= strconv.Atoi(idStr)
	if err != nil{
		return ErrInvalidInput
	}
	delete(m.data, id)
	return nil
}