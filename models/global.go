package models

import "hold-door/config"

var Timeout int = config.GetConfig().Get("service_secret.AuthCenter").(int)
