package serve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
)

type SendSmsForm struct {
	PhoneNumbers    []string               `json:"phone_numbers" binding:"required"`
	SignName        string                 `json:"sign_name" binding:"required"`
	TemplateCode    string                 `json:"template_code" binding:"required"`
	TemplateParam   map[string]interface{} `json:"template_param"`
	SmsUpExtendCode string                 `json:"sms_up_extend_code"`
	OutId           string                 `json:"out_id"`
}

func (api *API) QuerySmsTemplate(c *gin.Context) {
	code := c.Param("template_code")

	request := dysmsapi.CreateQuerySmsTemplateRequest()
	request.Scheme = "https"
	request.TemplateCode = code

	resp, err := api.smsapi.QuerySmsTemplate(request)
	if err != nil {
		logrus.Errorf("获取模板 %s 失败，错误信息：\n%s", code, err.Error())
		api.writeError(c, gin.H{
			"error":       "系统错误",
			"message":     fmt.Sprintf("获取模板 %s 失败", code),
			"responseErr": err.Error(),
		})
		return
	}
	// fmt.Printf("response is %#v\n", resp)
	if resp.Code != "OK" {
		api.writeError(c, gin.H{
			"error":    "获取模板失败",
			"message":  fmt.Sprintf("获取模板 %s 失败：%s", code, resp.Message),
			"response": resp,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   resp,
	})
}

func (api *API) SendSms(c *gin.Context) {
	var form SendSmsForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = strings.Join(form.PhoneNumbers, ",")
	request.SignName = form.SignName
	request.TemplateCode = form.TemplateCode
	if form.TemplateParam != nil {
		temp, err := json.Marshal(form.TemplateParam)
		if err != nil {
			api.writeError(c, gin.H{
				"error":   "系统错误",
				"message": fmt.Sprintf("JSON转换 TemplateParam %#v 错误：%s", form.TemplateParam, err.Error()),
			})
			return
		}
		request.TemplateParam = string(temp)
	}
	if form.SmsUpExtendCode != "" {
		request.SmsUpExtendCode = form.SmsUpExtendCode
	}
	if form.OutId != "" {
		request.OutId = form.OutId
	}

	resp, err := api.smsapi.SendSms(request)
	if err != nil {
		logrus.Errorf("发送短信 %s 给 %s 失败，错误信息：\n%s", form.TemplateCode, form.PhoneNumbers, err.Error())
		api.writeError(c, gin.H{
			"error":       "系统错误",
			"message":     fmt.Sprintf("发送短信 %s 给 %s 失败", form.TemplateCode, form.PhoneNumbers),
			"responseErr": err.Error(),
		})
		return
	}
	// fmt.Printf("response is %#v\n", resp)
	if resp.Code != "OK" {
		api.writeError(c, gin.H{
			"error":    "发送短信失败",
			"message":  fmt.Sprintf("发送短信 %s 给 %s 失败", form.TemplateCode, form.PhoneNumbers),
			"response": resp,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   resp,
	})
}
