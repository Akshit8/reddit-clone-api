package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	host = "smtp.ethereal.email"
	port = 587
	user = "kira61@ethereal.email"
	pass = "SuxjFxNeSbQ6FzEhbR"
)

func TestSendMail(t *testing.T) {
	mailCleint := NewMailer(host, port, user, pass)
	toMail := []string{"akshitsadana@gmail.com"}
	err := mailCleint.SendMail(toMail, "this is a test mail")

	require.NoError(t, err)
}
