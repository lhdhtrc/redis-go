package redis

import "github.com/fireflycore/go-utils/tlsx"

// Conf 定义 redis 连接初始化所需的配置项。
type Conf struct {
	Address  string `json:"address"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`

	// Tls 为 TLS 配置，非空且字段齐全时启用双向 TLS。
	Tls *tlsx.TLS `json:"tls"`

	// MaxOpenConnects 为连接池大小（近似最大连接数）。
	MaxOpenConnects int `json:"max_open_connects"`
	// MaxIdleConnects 为连接池最大空闲连接数。
	MaxIdleConnects int `json:"max_idle_connects"`
	// ConnMaxLifeTime 为连接最大生命周期，单位秒；0 表示不设置。
	ConnMaxLifeTime int `json:"conn_max_life_time"`
}
