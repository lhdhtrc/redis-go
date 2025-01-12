package redis

type TLS struct {
	CaCert        string `json:"ca_cert" bson:"ca_cert" yaml:"ca_cert" mapstructure:"ca_cert"`
	ClientCert    string `json:"client_cert" bson:"client_cert" yaml:"client_cert" mapstructure:"client_cert"`
	ClientCertKey string `json:"client_cert_key" bson:"client_cert_key" yaml:"client_cert_key" mapstructure:"client_cert_key"`
}

type Config struct {
	Tls TLS `json:"tls" bson:"tls" yaml:"tls" mapstructure:"tls"`

	Account  string `json:"account" bson:"account" yaml:"account" mapstructure:"account"`
	Password string `json:"password" bson:"password" yaml:"password" mapstructure:"password"`

	Address  string `json:"address" yaml:"address" mapstructure:"address"`
	Database string `json:"database" yaml:"database" mapstructure:"database"`

	Mode bool `json:"mode" yaml:"mode" mapstructure:"mode"` // Mode is true cluster

	MaxOpenConnects int `json:"max_open_connects" bson:"max_open_connects" yaml:"max_open_connects" mapstructure:"max_open_connects"`
	MaxIdleConnects int `json:"max_idle_connects" bson:"max_idle_connects" yaml:"max_idle_connects" mapstructure:"max_idle_connects"`
	ConnMaxLifeTime int `json:"conn_max_life_time" bson:"conn_max_life_time" yaml:"conn_max_life_time" mapstructure:"conn_max_life_time"`
}
