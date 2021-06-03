package user

import (
	t "clean-code-golang/core/domain/users"
	i "clean-code-golang/core/infrastructure/ioc/commons"

	_ "github.com/go-sql-driver/mysql"
)

type IUserRepository interface {
	FindAll() []t.User
	FindById(id string) t.User
	Save(user t.User)
	Update(user t.User)
	DeleteById(id string)
	FindByEmail(email string) (t.User, error)
}

type UserRepository t.User

func (u UserRepository) FindAll() []t.User {

	c := i.Resolver()

	var users []t.User

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Find(&users)
	})

	return users
}

func (u UserRepository) FindByEmail(email string) (t.User, error) {

	c := i.Resolver()

	var user t.User

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.First(&user, "email = ?", email)
	})

	return user, nil
}

func (u UserRepository) FindById(id string) t.User {

	c := i.Resolver()

	var user t.User

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.First(&user, "id = ?", id)
	})

	return user
}

func (u UserRepository) Save(user t.User) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Create(&user)
	})

}

func (u UserRepository) Update(user t.User) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		db := config.Connection.Context()
		db.Model(&user).Updates(t.User{
			Email: user.Email,
		})
	})

}

func (u UserRepository) DeleteById(id string) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		var user t.User
		db := config.Connection.Context()
		db.Delete(&user, id)
	})

}
