package users

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id: 1, Username: "user1"},
	{Id: 2, Username: "user2"},
	{Id: 3, Username: "user3"},
}
