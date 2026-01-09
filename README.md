## go-redis

基于 go-redis/v9 的工程化封装，提供：
- Redis Client 初始化（可选 TLS、连接池、连接生命周期）
- 初始化连通性检查（New 内置 Ping，默认超时 10s）

## 安装

```bash
go get github.com/fireflycore/go-redis
```

## 快速开始

```go
package main

import "github.com/fireflycore/go-redis"

func main() {
	conf := &redis.Conf{
		Address:         "127.0.0.1:6379",
		Database:        "0",
		MaxOpenConnects: 100,
		MaxIdleConnects: 10,
		ConnMaxLifeTime: 600,
	}

	client, err := redis.New(conf)
	if err != nil {
		panic(err)
	}

	_ = client
}
```

## 配置说明

初始化配置结构为 redis.Conf。

常用字段：
- Address：Redis 地址，一般为 host:port
- Database：逻辑库号（字符串形式），空表示默认 0
- Username/Password：鉴权信息（可选）
- MaxOpenConnects/MaxIdleConnects：连接池配置（分别映射到 PoolSize/MaxIdleConns）
- ConnMaxLifeTime：连接最大生命周期（单位秒，<=0 表示不设置）

### TLS

当 Conf.Tls 同时配置了 CaCert / ClientCert / ClientCertKey 三个文件路径时启用 TLS，否则视为不启用：

```go
conf := &redis.Conf{
	Address:  "127.0.0.1:6379",
	Database: "0",
	Tls: &redis.TLS{
		CaCert:        "/path/to/ca.pem",
		ClientCert:    "/path/to/client.pem",
		ClientCertKey: "/path/to/client.key",
	},
}

client, err := redis.New(conf)
_ = client
_ = err
```
