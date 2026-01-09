package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(c *Conf) (*redis.Client, error) {
	if c == nil {
		return nil, errors.New("redis: nil conf")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := parseDatabase(c.Database)
	if err != nil {
		return nil, err
	}

	clientOptions := redis.Options{
		// Addr 为 redis 地址，一般为 host:port。
		Addr: c.Address,
		// DB 为 redis 逻辑库号。
		DB: db,
	}

	if c.Username != "" {
		clientOptions.Username = c.Username
	}
	if c.Password != "" {
		clientOptions.Password = c.Password
	}

	// 从配置生成 TLSConfig；tlsEnabled 表示是否启用 TLS。
	tlsConfig, tlsEnabled, err := NewTLSConfig(c.Tls)
	// TLS 配置构造失败时直接返回错误。
	if err != nil {
		return nil, err
	}
	// 启用 TLS 时，将 TLSConfig 写入 clientOptions。
	if tlsEnabled {
		clientOptions.TLSConfig = tlsConfig
	}

	if c.MaxIdleConnects > 0 {
		clientOptions.MaxIdleConns = c.MaxIdleConnects
	}

	if c.MaxOpenConnects > 0 {
		clientOptions.PoolSize = c.MaxOpenConnects
	}

	if c.ConnMaxLifeTime > 0 {
		clientOptions.ConnMaxLifetime = time.Duration(c.ConnMaxLifeTime) * time.Second
	}

	// 创建 redis 客户端实例
	client := redis.NewClient(&clientOptions)

	// 通过 Ping 进行连通性校验，确保返回的 client 可用。
	if c := client.Ping(ctx); c.Err() != nil {
		client.Close()
		return nil, c.Err()
	}

	return client, nil
}
