package models

type GenerateResponse struct {
	Success 		bool
	AlreadyExists	bool
	Error			error
	GeneratedUrl	ElfUrl
}
