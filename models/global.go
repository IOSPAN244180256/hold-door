package models

import (
	"hold-door/config"
	"time"
)

var Timeout time.Duration = time.Duration(config.GetConfig().Get("timeout").(int)) * time.Second

var DatetimeTemplate string = "2006/1/2 15:04:05" //时间格式化模板
