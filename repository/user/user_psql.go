package userRepository

import (
	"books-app/models"
	"database/sql"
	"log"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	logFatal(err)

	user.Password = ""
	return user
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (b UserRepository) GetUsers(db *sql.DB, user models.User, users []models.User) []models.User {
	rows, err := db.Query("select * from users")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		logFatal(err)

		users = append(users, user)
	}

	return users
}

func (b UserRepository) GetUser(db *sql.DB, user models.User, id int) models.User {
	rows := db.QueryRow("select * from users where id=$1", id)

	err := rows.Scan(&user.ID, &user.Email, &user.Password)
	logFatal(err)

	return user
}

func (b UserRepository) RemoveUser(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from users where id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
