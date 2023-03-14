package telegram

type Update struct {
	UpdateID int     `json:"update_id,omitempty"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int    `json:"message_id,omitempty"`
	From      User   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date,omitempty"`
	Text      string `json:"text,omitempty"`
}

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

type Chat struct {
	ID   int    `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
