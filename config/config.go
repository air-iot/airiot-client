package config

import "time"

type Config struct {
	Gateway    string             `json:"gateway"`
	EtcdConfig string             `json:"etcdConfig"`
	Metadata   map[string]string  `json:"metadata"`
	Services   map[string]Service `json:"services"`
	Type       KeyType            `json:"type"`
	ProjectId  string             `json:"projectId"`
	AK         string             `json:"ak"`
	SK         string             `json:"sk"`
	Timeout    uint               `json:"timeout"`
	Limit      int                `json:"limit"`
	Debug      bool               `json:"debug"`
	Service    struct {
		//Enable bool          `json:"enable"`
		Expire time.Duration `json:"expire"`
	} `json:"service"`
}

type KeyType string

const (
	Tenant  KeyType = "tenant"
	Project KeyType = "project"
)

type Service struct {
	Metadata map[string]string `json:"metadata"`
}

const (
	XRequestProject             = "x-request-project"
	XRequestProjectDefault      = "default"
	XRequestHeaderAuthorization = "Authorization"
	XRequestQueryAuthorization  = "token"
)
