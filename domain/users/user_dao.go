package users

import (
	"fmt"
	// "strings"
	// "github.com/go-sql-driver/mysql"
	"github.com/selvamshan/bookstore_user-api/utils/errors"
	//"github.com/selvamshan/bookstore_user-api/utils/date_utils"
	//"github.com/selvamshan/bookstore_user-api/utils/mysql_utils"
	"github.com/selvamshan/bookstore_user-api/datasources/mysql/users_db"
	"github.com/selvamshan/bookstore_user-api/logger"
)

const(
	indexUniqueEamil = "email_UNIQUE"	
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created, status FROM users Where id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get(method string) *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to prepare %s user statment", method), err)
		return errors.NewInternalServerError("data base error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {	
		// return mysql_utils.ParseError(getErr)
		logger.Error(fmt.Sprintf("error when trying to scan row in user struct %s user by id", method), getErr)
		return errors.NewInternalServerError("data base error")		
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statment", err)
		return errors.NewInternalServerError("data base error")
	}
	defer stmt.Close()
    if err := user.Validate(); err!= nil {
		return err
	}
	
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("data base error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating new user", err)
		return errors.NewInternalServerError("data base error")
	}
	
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statment", err)
		return errors.NewInternalServerError("data base error")
	}
	defer stmt.Close()

    _, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to excute update user stmt", err)
		return errors.NewInternalServerError("data base error")
	}
	return nil
}


func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user stmt", err)
		return errors.NewInternalServerError("data base error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to excute delete user stmt ", err)
		return errors.NewInternalServerError("data base error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find by status user stmt", err)
		return nil, errors.NewInternalServerError("data base error")		
	}
	defer stmt.Close()

	// users = []User
	// rows, err := stmt.QueryRows(&users)
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to query find by status user stmt", err)
		return nil, errors.NewInternalServerError("data base error")	
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan user row into users struct in find by status", err)
			return nil, errors.NewInternalServerError("data base error")	
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}