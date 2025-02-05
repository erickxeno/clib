package common

const (
	HookOrderLowWeight     = 0
	HookOrderStartMqServer = 200
	HookOrderMiddleWeight  = 500
	HookOrderDefault       = HookOrderMiddleWeight
	HookOrderStartES       = 600
	HookOrderStartKV       = 600
	HookOrderStartGRAPH    = 600
	HookOrderStartDOC      = 600
	HookOrderStartMysql    = 600
	HookOrderStartRedis    = 700
	HookOrderHighWeight    = 1000
)
