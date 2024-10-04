package entity

// import "github.com/google/uuid"


type Token struct {
	// Id          uuid.UUID
	Id          string
	Login       string
	LoginType 	string
	PassHash    string
	Token       string
	Permissions *[]Permission
}