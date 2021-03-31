package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	host = "smtp.ethereal.email"
	port = 587
	user = "cory74@ethereal.email"
	pass = "x4NY7ww2DpCfVueHba"
)

func TestSendMail(t *testing.T) {
	mailCleint := NewMailer(host, port, user, pass)
	toMail := []string{"akshit@gmail.com"}
	err := mailCleint.SendMail(toMail, "this is a test mail")

	require.NoError(t, err)
}
