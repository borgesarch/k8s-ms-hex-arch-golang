package todo

import (
	t "ms-hex-arch-golang-k8s/core/domain/todo"
	i "ms-hex-arch-golang-k8s/core/infrastructure/ioc/commons"

	_ "github.com/go-sql-driver/mysql"
)

type ITodoRepository interface {
	FindAll() []t.Todo
	FindById(id string) t.Todo
	Save(todo t.Todo)
	Update(todo t.Todo)
	DeleteById(id string)
}

type TodoRepository t.Todo

func (u TodoRepository) FindAll() []t.Todo {

	c := i.Resolver()

	var todos []t.Todo

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Find(&todos)
	})

	return todos
}

func (u TodoRepository) FindById(id string) t.Todo {

	c := i.Resolver()
	var todo t.Todo

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.First(&todo)
	})

	return todo
}

func (u TodoRepository) Save(todo t.Todo) {
	c := i.Resolver()
	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Create(&todo)
	})
}

func (u TodoRepository) Update(todo t.Todo) {
	c := i.Resolver()
	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Model(&todo).Updates(t.Todo{
			Description: todo.Description,
			Completed:   todo.Completed,
		})
	})
}

func (u TodoRepository) DeleteById(id string) {
	c := i.Resolver()
	var todo t.Todo
	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Delete(&todo, id)
	})
}
