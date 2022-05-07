package types

import "time"

// 配置

type Config struct {
	Id         int64     `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Endpoint   string    `json:"endpoint,omitempty"`
	Method     string    `json:"method,omitempty"`
	Body       string    `json:"body,omitempty"`
	CreateTime time.Time `json:"crateTime,omitempty"`
	ModifyTime time.Time `json:"modifyTime,omitempty"`
}
