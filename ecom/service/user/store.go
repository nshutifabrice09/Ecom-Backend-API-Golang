package user

import (
	"database/sql"
	"ecom/types"
	"fmt"
)

type Store struct {
	db *sql.DB
}

// CreateUser implements types.UserStore.
func (s *Store) CreateUser(types.User) error {
	panic("unimplemented")
}

// GetUserById implements types.UserStore.
func (s *Store) GetUserById(id int) (*types.User, error) {
	panic("unimplemented")
}

// UpdateUser implements types.UserStore.
func (s *Store) UpdateUser(types.User) error {
	panic("unimplemented")
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		if user.ID == 0 {
			return nil, fmt.Errorf("user not found!")
		}
	}
}

func scanRowIntoUser(rows *sql.Row) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Address,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
