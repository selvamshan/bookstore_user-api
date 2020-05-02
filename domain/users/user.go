package users

import (
	"strings"
    "regexp"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

var (
	re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

const (
	StatusActive="active"
)

type User struct {
	Id 			int64  `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string  `json:"password"`
}

type Users []User


func (user *User) Validate() *rest_errors.RestErr {
    user.TrimSpaceInNames()
    if err := user.ValidateEmail(); err != nil {
		return err
	}
	if err := user.ValidatePassword(); err != nil {
		return err
	}
	
	return nil
}

func (user *User) TrimSpaceInNames() {
	user.FirstName = strings.TrimSpace(user.FirstName)
    user.LastName = strings.TrimSpace(user.LastName)
}


func (user *User) ValidateEmail() *rest_errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == ""{
		return rest_errors.NewBadRequestError("invalid email address")
	}
	if !re.MatchString(user.Email) {
		// fmt.Println(user.Email)
		return rest_errors.NewBadRequestError("invalid email address")	
	}
	return nil
}

func (user *User) ValidatePassword() *rest_errors.RestErr {
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")	
	}
	return nil
}