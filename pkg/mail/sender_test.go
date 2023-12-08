package mail

import (
	"github.com/c1tad3l/backend-wedo/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGmailSender_SendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	info, err := config.LoadConfig()
	require.NoError(t, err)
	t.Errorf(
		"For", info,
	)

	sender := NewGmailSender(info.EmailSenderName, info.EmailSenderAddress, info.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	`
	to := []string{"fellowgram@gmail.com"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)

	require.NoError(t, err)

}
