package rediskey

import "template-project/config"

var (
	// 验证码 redis key
	CaptchaCodeKey = config.Data.Lili.Name + ":captcha:code:"

	// 登录账户密码错误次数 redis key
	LoginPasswordErrorKey = config.Data.Lili.Name + ":login:password:error:"

	// 登录用户 redis key
	UserTokenKey = config.Data.Lili.Name + ":user:token:"

	// 防重提交 redis key
	RepeatSubmitKey = config.Data.Lili.Name + ":repeat:submit:"

	// 配置表数据 redis key
	SysConfigKey = config.Data.Lili.Name + "system:config"

	// 字典表数据 redis key
	SysDictKey = config.Data.Lili.Name + "system:dict:data"
)
