package ioc

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	aliSms "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

func InitAliSmsClient() *aliSms.Client {
	accessKeyID := os.Getenv("ALICLOUD_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALICLOUD_ACCESS_KEY_SECRET")
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(accessKeyID),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	client, err := aliSms.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
