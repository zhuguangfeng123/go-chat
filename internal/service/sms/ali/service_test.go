package ali

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sms "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSender(t *testing.T) {
	config := &openapi.Config{}

	client, err := sms.NewClient(config)
	assert.NoError(t, err)

	s := NewAliService(client)

	err = s.Send(context.Background(), "物联科技", "SMS_467460115", "123456", "18860313695")
	assert.NoError(t, err)

}
