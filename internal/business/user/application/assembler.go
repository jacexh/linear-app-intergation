package application

import (
	"github.com/jacexh/linear-app-intergation/internal/business/user/domain"
	"github.com/jinzhu/copier"
)

func assembleDomainUser(entity *domain.User) (*User, error) {
	du := new(User)
	if err := copier.Copy(du, entity); err != nil {
		return nil, err
	}
	return du, nil
}
