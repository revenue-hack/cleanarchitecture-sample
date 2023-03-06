package useroutput

type UserList struct {
	Users []*UserItem `json:"users"`
}
type UserItem struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
