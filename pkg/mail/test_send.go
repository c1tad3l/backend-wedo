package mail

import (
	"github.com/c1tad3l/backend-wedo/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func testsend(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	info, err := config.LoadConfig()
	require.NoError(t, err)

	sender := NewGmailSender(info.EmailSenderName, info.EmailSenderAddress, info.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	`
	to := []string{"zloymolodoy88@gmail.com"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
