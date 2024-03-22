package providers

import (
	"simplylab/model"
)

type UserProvider struct {

}

func (p UserProvider) GetUserByName(userName string) model.User {
	return model.User{
			ID:        "",
			Name:      "",
			CreatedAt: model.CreatedAt{},
			UpdatedAt: model.UpdatedAt{},
		}
}