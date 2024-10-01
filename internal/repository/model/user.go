package model

import (
	"database/sql"
	"time"
)

// UserRepoModel - модель User
type UserRepoModel struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      int64        `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserData struct {
	Username string `json:"username"`
	Role     int64  `json:"role"`
}
