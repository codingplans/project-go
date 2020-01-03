package config

import (
	"iceberg/frame/config"
	"time"
)

var DefaultConfig *Config

type Config struct {
	Env                      config.Environment `json:"env"`
	Base                     config.BaseCfg     `json:"baseCfg"`
	WSAddr                   string             `json:"-"`
	PublicWSAddr             string             `json:"-"`
	HTTPTimeoutSeconds       time.Duration      `json:"http_timeout_seconds"`
	ClientTimeoutMinutes     time.Duration      `json:"client_timeout_minutes"`
	BroadcastQueueSize       int                `json:"broadcast_queue_size"`
	HeartbeatIntervalSeconds time.Duration      `json:"heartbeat_interval_seconds"`
	RedisProdOrder           config.RedisCfg    `json:"redis_prod_order"`
	RedisAuth                config.RedisCfg    `json:"redis_auth"`
	RedisLoop                config.RedisCfg    `json:"redis_loop"`
	NsqCfg                   config.NsqCfg      `json:"nsq_cfg"`
	MysqlAppMain             config.MysqlCfg    `json:"mysql_app_main_cfg"`
	MysqlGameserverCfg       config.MysqlCfg    `json:"mysql_gameserver_cfg"`
	MysqlChatroomCfg         config.MysqlCfg    `json:"mysql_chatroom_cfg"`
	MysqlFeedCfg             config.MysqlCfg    `json:"mysql_feed_cfg"`
	MaxClientCount           int                `json:"max_client_count"`
	ShenceCfg                struct {
		Switch  bool   `json:"switch"`
		Host    string `json:"host"`
		Timeout int    `json:"timeout"`
		Project string `json:"project"`
	} `json:"shenceCfg"` // 神策埋点统计
	DotKafka config.KafkaCfg `json:"dotKafka"`
}

func Init(filepath string) {
	config.Parseconfig(filepath, &DefaultConfig)
	if DefaultConfig.BroadcastQueueSize == 0 {
		DefaultConfig.BroadcastQueueSize = 10000
	}
	if DefaultConfig.HTTPTimeoutSeconds == 0 {
		DefaultConfig.HTTPTimeoutSeconds = 10 * time.Second
	} else {
		DefaultConfig.HTTPTimeoutSeconds = DefaultConfig.HTTPTimeoutSeconds * time.Second
	}
	if DefaultConfig.ClientTimeoutMinutes == 0 {
		DefaultConfig.ClientTimeoutMinutes = 3 * time.Minute
	} else {
		DefaultConfig.ClientTimeoutMinutes = DefaultConfig.ClientTimeoutMinutes * time.Minute
	}
	if DefaultConfig.HeartbeatIntervalSeconds == 0 {
		DefaultConfig.HeartbeatIntervalSeconds = 10 * time.Second
	} else {
		DefaultConfig.HeartbeatIntervalSeconds = DefaultConfig.HeartbeatIntervalSeconds * time.Second
	}
	if DefaultConfig.MaxClientCount <= 0 {
		DefaultConfig.MaxClientCount = 10000
	}
}
