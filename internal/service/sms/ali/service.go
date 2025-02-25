package ali

import (
	"context"
	"errors"
	aliSms "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/zhuguangfeng123/go-chat/internal/service/sms"
)

type Service struct {
	client *aliSms.Client
}

func NewAliService(client *aliSms.Client) sms.Service {
	return &Service{
		client: client,
	}
}

// Send 发送验证码阿里云实现
func (s *Service) Send(ctx context.Context, signName string, tpl string, arg string, phone string) error {
	sendSmsRequest := &aliSms.SendSmsRequest{
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(tpl),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(arg),
	}

	res, err := s.client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}

	if *res.Body.Code != "OK" {
		return errors.New("发送短信失败" + *res.Body.Message)
	}
	return nil
}
