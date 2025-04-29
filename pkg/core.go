package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
	"time"
)

func New(config *Config) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db, _ := strconv.Atoi(config.Database)
	clientOptions := redis.Options{
		Addr:         config.Address,
		DB:           db,
		MaxIdleConns: config.MaxIdleConnects,
		PoolSize:     config.MaxOpenConnects,
	}

	if config.Username != "" && config.Password != "" {
		clientOptions.Username = config.Username
		clientOptions.Password = config.Password
	}

	if config.Tls.CaCert != "" && config.Tls.ClientCert != "" && config.Tls.ClientCertKey != "" {
		certPool := x509.NewCertPool()
		CAFile, CAErr := os.ReadFile(config.Tls.CaCert)
		if CAErr != nil {
			return nil, CAErr
		}
		certPool.AppendCertsFromPEM(CAFile)

		clientCert, clientCertErr := tls.LoadX509KeyPair(config.Tls.ClientCert, config.Tls.ClientCertKey)
		if clientCertErr != nil {
			return nil, clientCertErr
		}

		tlsConfig := tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}
		clientOptions.TLSConfig = &tlsConfig
	}

	cli := redis.NewClient(&clientOptions)

	if c := cli.Ping(ctx); c.Err() != nil {
		return nil, c.Err()
	}

	return cli, nil
}
