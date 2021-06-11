package auth

import (
	t "ms-hex-arch-golang-k8s/core/domain/users"
	n "ms-hex-arch-golang-k8s/core/infrastructure/ioc/commons"
	i "ms-hex-arch-golang-k8s/core/infrastructure/ioc/repositories"
)

type IUserService interface {
	Save(user t.User) (string, error)
}

type UserService int

func (o UserService) Save(user t.User) (t.User, error) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {

		x := n.Resolver()

		x.Invoke(func(con *n.Config) {
			user.Password = con.EncrityHash.HashPassword(user.Password)
		})

		config.UserRepository.Save(user)

	})

	return user, nil
}
