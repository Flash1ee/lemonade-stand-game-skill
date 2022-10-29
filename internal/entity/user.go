package entity

type VKUser struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type VKFriends struct {
	Count int64    `json:"count"`
	Ids   []string `json:"ids"`
}
