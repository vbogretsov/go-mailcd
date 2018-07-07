package mock_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vbogretsov/go-mailcd"
	"github.com/vbogretsov/go-mailcd/mock"
)

func TestSendAddEmailToInbox(t *testing.T) {
	sender := mock.New()

	req := mailcd.Request{
		TemplateLang: "en",
		TemplateName: "test",
		TemplateArgs: map[string]interface{}{
			"test": "test",
		},
		To: []mailcd.Address{
			{
				Email: "to1@mail.com",
			},
		},
	}

	err := sender.Send(req)

	require.Nil(t, err, "send error: %v", err)
	require.Equal(t, 1, len(sender.Inbox))
	require.Equal(t, req, sender.Inbox[0])
}
