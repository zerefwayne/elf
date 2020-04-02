package models

type Response struct {
	Success 		bool			`json:"success"`
	AlreadyExists	bool			`json:"alreadyExists"`
	Payload			interface{}		`json:"payload"`
}
