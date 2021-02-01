package lib

// GetUpdatesResponse model
type GetUpdatesResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update model
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message,omitempty"`
}

// Message model
type Message struct {
	MessageID int    `json:"message_id"`
	From      User   `json:"from"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

// User model
type User struct {
	Username string `json:"username"`
}

// Chat model
type Chat struct {
	ID int `json:"id"`
}

// SendMessage model
type SendMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

// GetMeResult model
type GetMeResult struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name,omitempty"`
	Username  string `json:"username"`
}

// GetMeResponse model
type GetMeResponse struct {
	OK     bool        `json:"ok"`
	Result GetMeResult `json:"result"`
}
