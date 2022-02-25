package entities

import "time"

type School struct {
	Id         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Address    string    `json:"address" db:"address"`
	City       string    `json:"city" db:"city"`
	Providence string    `json:"providence" db:"providence"`
	Zip        string    `json:"zip" db:"zip"`
	Phone      string    `json:"phone" db:"phone"`
	Fax        string    `json:"fax" db:"fax"`
	Email      string    `json:"email" db:"email"`
	Website    string    `json:"website" db:"website"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}
