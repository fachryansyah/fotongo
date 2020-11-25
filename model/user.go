package model

import "log"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Level    int32  `json:"level"`
	Password string `json:""`
}

func InsertUser(user User) (*string, error) {
	var id string
	err := db.QueryRow("INSERT INTO users (name, email, password, level) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name,
		user.Email,
		user.Password,
		user.Level,
	).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &id, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User

	err := db.QueryRow("SELECT id, name, email, level, password FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Level, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
