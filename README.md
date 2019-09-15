# ga.sms.aliyun

阿里云短信服务

面向服务架构，我们系统提供 API 而不是 SDK 。
本服务封装阿里云 Go SDK ，提供短息服务接口。

## 参考

- [短信服务](https://api.aliyun.com/?spm=a2c4g.11186623.2.13.1d8051c14mUHUU#/?product=Dysmsapi&api=QuerySmsTemplate&tab=DEMO&lang=GO)

## 运行服务

配置环境变量：

```shell
export GA_SMS_ACCESS_KEY_ID=<你的阿里云短信服务 AccessKeyID>
export GA_SMS_ACCESS_KEY_SECRET=<你的阿里云短信服务 AccessKeySecret>
```

启动命令：

```shell
./ga.sms.aliyun serve -v --port=3000
```

## TODO

- [ ] 支持发送频次限制（来源 IP，发送到的手机号等）
