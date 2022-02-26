package dtos

type SchoolCreateRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Providence string `json:"providence"`
	Zip        string `json:"zip"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	Email      string `json:"email"`
	Website    string `json:"website"`
}

type SchoolUpdateRequest struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Providence string `json:"providence"`
	Zip        string `json:"zip"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	Email      string `json:"email"`
	Website    string `json:"website"`
}
