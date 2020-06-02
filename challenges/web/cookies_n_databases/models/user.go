package models

//DBSession stores the session per user
var DBSession = map[string]User{}

//DBUsers stores the users
var DBUsers = map[string]string{
	//Hash = naruto; it's in rock you
	`'admin@cybercastors.com'`:    `'cf9ee5bcb36b4936dd7064ee9b2f139e'`,
	`'admin@powerpuffgirls.com'`:  `'fe87c92e83ff6523d677b7fd36c3252d'`,
	`'jeff@homeaddress.com'`:      `'d1833805515fc34b46c2b9de553f599d'`,
	`'moreusers@leakingdata.com'`: `'77004ea213d5fc71acf74a8c9c6795fb'`,
}

//User struct holds user data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
