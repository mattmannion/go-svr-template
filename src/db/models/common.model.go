package models

type JSON struct {
	Status string `json:"status"`
	User   Users  `json:"user"`
}

type JSON_MSG struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
