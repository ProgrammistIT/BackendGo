package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID          int
	Name        string
	Description string
	IsDone      bool
}

type Storage interface {
	CreateTask(name string, description string)
	GetTask(id int) (Task, error)
	UpdateTask(id int, newDescription string, isDone bool) error
	DeleteTask(id int) error
}

type MemoryStorage struct {
	count int
	tasks map[int]Task
}

func (m *MemoryStorage) CreateTask(name string, description string) {
	m.count++
	id := m.count
	m.tasks[id] = Task{
		ID:          id,
		Name:        name,
		Description: description,
		IsDone:      false,
	}
}

func (m *MemoryStorage) GetTask(id int) (Task, error) {
	task, ok := m.tasks[id]
	if ok {
		return task, nil
	}

	return Task{}, errors.New("Задача не найдена")

}

func (m *MemoryStorage) UpdateTask(id int, newDescription string, status bool) error {
	task, ok := m.tasks[id]
	if ok {
		task.IsDone = status
		task.Description = newDescription
		m.tasks[id] = task
		return nil
	}
	return errors.New("Задача не найдена")
}

func (m *MemoryStorage) DeleteTask(id int) error {
	_, ok := m.tasks[id]

	if ok {
		delete(m.tasks, id)
		return nil
	}
	return errors.New("Задача не найдена")
}

func main() {
	var storage Storage = &MemoryStorage{
		tasks: make(map[int]Task),
	}

	storage.CreateTask("ДЗ по Backend", "Реализовать структуру Task по ТЗ и проверить код на работоспособность")
	task, err := storage.GetTask(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(task.Name)
		fmt.Println(task.Description)
	}
	fmt.Println("")

	err = storage.UpdateTask(1, "Проверить код на работоспособность", true)
	task, err = storage.GetTask(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(task.Name)
		fmt.Println(task.Description)
	}
	fmt.Println("")

	err = storage.DeleteTask(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Удаление произошло успешно!")
		_, err := storage.GetTask(1)
		fmt.Println(err)
	}
}
