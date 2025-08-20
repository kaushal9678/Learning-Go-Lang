package models

import (
	"errors"

	"example.com/rest-api-go/db"
	"example.com/rest-api-go/utils"
)

type User struct {
	ID int64 `json:"id"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u User) Save() error{
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query); if err != nil{
		return err
	}
	defer stmt.Close();
	hashedpassword, err := utils.HashPassWord(u.Password); if err != nil{
		return err
	}

	result, err := stmt.Exec(u.Email, hashedpassword); if err != nil{
		return err
	}
	userId, err := result.LastInsertId(); if err != nil{
		return err;
	}
	u.ID = userId
	return nil
}
func GetUsers() ([]User, error) {
	query := `SELECT * FROM users`;
	rows, err := db.DB.Query(query); if err != nil{
		return nil, err;
	}
	defer rows.Close();
  var users []User;

  for rows.Next(){
	var user User;
	if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil{
		return nil, err;
	}
	users = append(users, user);
  }
  if err := rows.Err(); err != nil{
	return nil,err;
  }
return users, nil
}

func (u User) ValidateCredentials() error{
	query := `SELECT password FROM users WHERE email = ?`;
	row := db.DB.QueryRow(query, u.Email);
	
	var retrievedPassword string;
	if err := row.Scan(&retrievedPassword); err != nil {
		return err;
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials are not valid")
	}

	return nil
}
func (u User) GetUserByEmailId() (*User, error) {
	query := `SELECT * FROM users WHERE email = ?`
	stmt, err := db.DB.Prepare(query); if err != nil{
		return nil, err;
	}
	defer stmt.Close();
	row := stmt.QueryRow(u.Email);
	var user User;
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil{
		return nil, err;
	}
	return &user, nil
}
func (u User)Signin() (*User, error){
	query := `SELECT * FROM users WHERE email = ? AND password = ?`;
	stmt, err := db.DB.Prepare(query); if err != nil{
		return nil, err;
	}
	defer stmt.Close();
	row := stmt.QueryRow(u.Email, u.Password);
	var user User;
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, err;
	}
	return &user, nil
}