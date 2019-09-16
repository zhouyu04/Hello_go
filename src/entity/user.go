package entity

type User struct {
	Id        int64  `json:"id" gm:"id"`
	Username  string `json:"username"`
	Pwd       string `json:"pwd"`
	Nickname  string `json:"nickname"`
	Create_at string `json:"create_at"`
}

type UserMapper struct {
	SelectByUserAndPwd func(username string, pwd string) (User, error)
}
