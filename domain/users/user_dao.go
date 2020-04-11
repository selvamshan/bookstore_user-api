package users

import (
	"fmt"
	"time"
	"github.com/selvamshan/bookstore_user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr{
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("email %s alreay registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d alread exists", user.Id))
	}
	currTime := time.Now()
	user.DateCreated = currTime.Format("02-01-2006")
	userDB[user.Id] = user
	return nil
}