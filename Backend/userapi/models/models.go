package models

//User this is user type
type User struct {
	Userid      string `json:"userid,omitempty"`
	Password    string `json:"password,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty"`
}

//Registration this is user type
type Registration struct {
	Userid           string `json:"userid,omitempty"`
	Password         string `json:"password,omitempty"`
	Firstname        string `json:"firstname,omitempty"`
	Lastname         string `json:"lastname,omitempty"`
	Mailvalid        bool   `json:"mailvalid,omitempty"`
	Phonenumber      string `json:"phonenumber,omitempty"`
	Verificationcode string `json:"verificationcode,omitempty"`
	Timestamp        string `json:"timestamp,omitempty"`
}
