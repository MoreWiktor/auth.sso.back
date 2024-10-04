package models

type Permission struct {
	ServiceId   string
	Permissions []string
}

type Token struct {
	Id          string
	Login       string
	LoginType 	string
	PassHash    string
	Token       string
	Permissions *[]Permission
}
