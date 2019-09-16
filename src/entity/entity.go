package entity

type Terminals struct {
	Id               int64  `json:"id" gm:"id"`
	Dbid             int64  `json:"dbid"`
	Pos              int64  `json:"pos"`
	Name             string `json:"name"`
	Shop_name        string `json:"shop_name"`
	Ip_addr          string `json:"ip_addr"`
	Run_env          string `json:"run_env"`
	Core_ver         string `json:"core_ver"`
	Db_ver           string `json:"db_ver"`
	Vertion          string `json:"vertion"`
	Release_id       int64  `json:"release_id"`
	Await_release_id int64  `json:"await_release_id"`
	Status           int    `json:"status"`
}

type TerminalsMapper struct {
	SelectAll  func() ([]Terminals, error)
	SelectById func(id string) (Terminals, error) `mapperParams:"id"`
}
