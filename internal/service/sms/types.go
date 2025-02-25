package sms

import "context"

type Service interface {
	Send(ctx context.Context, signName string, tpl string, arg string, phone string) error
}
