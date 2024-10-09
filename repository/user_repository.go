// repository/user_repository.go
package repository

import (
	"database/sql"
	"github.com/mateusfaustino/go-rest-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := ur.connection.QueryRow(query, username)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
