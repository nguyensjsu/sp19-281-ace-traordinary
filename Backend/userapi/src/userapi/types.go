type user struct {
	userid      string
	password    string
	firstname   string
	lastname    string
	phonenumber string
}

type registration struct {
	userid           string
	password         string
	firstname        string
	lastname         string
	mailvalid        bool
	phonenumber      string
	verificationcode string
	timestamp        string
}
