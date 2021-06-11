package todo

import (
	t "ms-hex-arch-golang-k8s/core/domain/todo"
	i "ms-hex-arch-golang-k8s/core/infrastructure/ioc/repositories"
)

type ITodoService interface {
	save(todo t.Todo) t.Todo
	update(todo t.Todo) t.Todo
}

type TodoService t.Todo

func (o TodoService) Save(todo t.Todo) t.Todo {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		config.TodoRepository.Save(todo)
	})

	return todo
}

func (o TodoService) Update(todo t.Todo) t.Todo {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		config.TodoRepository.Update(todo)
	})

	return todo
}
