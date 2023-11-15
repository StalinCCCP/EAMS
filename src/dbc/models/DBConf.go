package models

type DBConf struct {
	Server   string `json:"Server"`
	DBName   string `json:"DBName"`
	User     string `json:"User"`
	Password string `json:"Password"`
}
