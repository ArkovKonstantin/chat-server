package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

type Chat struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Users     []int  `json:"users,omitempty"`
	CreatedAt string `json:"created_at"`
}

type Message struct {
	ID        int    `json:"id"`
	Chat      string `json:"chat"`
	Author    string `json:"author"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type AddMsgForm struct {
	Chat   int    `json:"chat"`
	Author int    `json:"author"`
	Text   string `json:"text"`
}
