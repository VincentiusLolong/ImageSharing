package models

import "time"

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	DB          int
}

var RedisSetting = &Redis{}

type Token struct {
	Refreshtoken string
}
