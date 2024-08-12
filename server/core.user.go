package server

import (
	"database/sql"
)

type User struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUser() ([]*User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Đảm bảo đóng rows sau khi sử dụng

	var users []*User

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//scan to get 1 object

func scanRowIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)

	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Age,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) PostUser(user *User) error {
	_, err := s.db.Exec("insert into users (name,age) values (? ,?)", user.Name, user.Age)

	if err != nil {
		return err
	}

	return nil
}
