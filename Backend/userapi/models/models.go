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

//Email service data
type Email struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Subject  string `json:"subject,omitempty"`
	HTMLBody string `json:"htmlbody,omitempty"`
	TextBody string `json:"textbody,omitempty"`
}

//TemplateData service data
type TemplateData struct {
	Firstname        string `json:"firstname,omitempty"`
	Verificationcode string `json:"verificationcodestring,omitempty"`
	Password         string `json:"password,omitempty"`
	URL              string `json:"url,omitempty"`
	Price            int64  `json:"price"`
}
type Picture struct {
	ImageId     string `json:"imageid"`
	UserId      string `json:"userid"`
	Title       string `json:"title"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	IsAvailable bool   `json:"isavailable"`
	TumbnailUrl string `json:"tumbnailUrl"`
	OrigUrl     string `json:"origurl"`
}
