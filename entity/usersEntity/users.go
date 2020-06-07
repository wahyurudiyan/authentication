package usersEntity

import "time"

type Users struct {
	ID              uint   `json:"id"`
	AccountUniqueID string `json:"account_unique_id"`
	Firstname       string `json:"firstname"`
	Surename        string `json:"surename"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Timestamp
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
