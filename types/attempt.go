package types

// Attempt - An attempt to login
type Attempt struct {
	Username  *Username `json:"username"`
	Token     *Token    `json:"token"`
	Origin    string    `json:"origin"`
	Timestamp string    `json:"timestamp"`
	Hash      []byte    `json:"hash"`
}

// NewAttempt - Create a new attempt struct
func NewAttempt() {

}
