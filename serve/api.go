package serve

import "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

type API struct {
	smsapi *dysmsapi.Client
}

func NewAPI(accessKeyID, accessKeySecret string) (*API, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}

	return &API{
		smsapi: client,
	}, nil
}
