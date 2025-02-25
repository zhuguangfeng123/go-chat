package code

import (
	"context"
	"fmt"
	"github.com/zhuguangfeng123/go-chat/internal/repository"
	"github.com/zhuguangfeng123/go-chat/internal/service/sms"
	"math/rand"
)

type CodeService interface {
	SendCode(ctx context.Context, biz string, phone string) error
	VerifyCode(ctx context.Context, biz string, inputCode string, phone string) (bool, error)
}

type codeService struct {
	smsSvc   sms.Service
	codeRepo repository.CodeRepository
}

func NewCodeService(smsSvc sms.Service, codeRepo repository.CodeRepository) CodeService {
	return &codeService{
		smsSvc:   smsSvc,
		codeRepo: codeRepo,
	}
}

// SendCode 发送验证码
func (svc *codeService) SendCode(ctx context.Context, biz string, phone string) error {
	//生成验证码
	code := svc.generateCode()
	//存储到redis
	err := svc.codeRepo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	//发送
	signName, tpl := svc.selectSmsByBiz(biz)
	err = svc.smsSvc.Send(ctx, signName, tpl, svc.selectArgByBiz(biz, code), phone)
	if err != nil {
		err1 := svc.codeRepo.DeleteCode(ctx, biz, phone)
		if err1 != nil {
			fmt.Println(err)
		}
	}
	return err
}

// VerifyCode 校验验证码
func (svc *codeService) VerifyCode(ctx context.Context, biz string, inputCode string, phone string) (bool, error) {
	return svc.codeRepo.Verify(ctx, biz, phone, inputCode)
}

func (svc *codeService) generateCode() string {
	num := rand.Intn(1000000)
	return fmt.Sprintf("%06d", num)
}

func (svc *codeService) selectSmsByBiz(biz string) (string, string) {
	switch biz {
	case "login":
		return "物联科技", "SMS_467460115"
	default:
		return "", ""
	}
}

func (svc *codeService) selectArgByBiz(biz string, code string) string {
	switch biz {
	case "login":
		return fmt.Sprintf("{\"code\":\"%s\"}", code)
	}
	return ""
}
