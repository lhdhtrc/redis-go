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

func New(conf *Conf) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db, _ := strconv.Atoi(conf.Database)
	clientOptions := redis.Options{
		Addr:         conf.Address,
		DB:           db,
		MaxIdleConns: conf.MaxIdleConnects,
		PoolSize:     conf.MaxOpenConnects,
	}

	if conf.Username != "" && conf.Password != "" {
		clientOptions.Username = conf.Username
		clientOptions.Password = conf.Password
	}

	if conf.Tls != nil && conf.Tls.CaCert != "" && conf.Tls.ClientCert != "" && conf.Tls.ClientCertKey != "" {
		certPool := x509.NewCertPool()
		CAFile, CAErr := os.ReadFile(conf.Tls.CaCert)
		if CAErr != nil {
			return nil, CAErr
		}
		certPool.AppendCertsFromPEM(CAFile)

		clientCert, clientCertErr := tls.LoadX509KeyPair(conf.Tls.ClientCert, conf.Tls.ClientCertKey)
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
