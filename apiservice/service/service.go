package service

import (
	"fire-press/api/types"
	"fire-press/apiservice/store"
)

type ConfigService struct{}

func NewConfigService() ConfigService {
	return ConfigService{}
}

func (cs ConfigService) Save(config types.Config) {
	store.SaveConfig(config)
}

func (cs ConfigService) Delete(id int64) int64 {
	return store.DeleteConfigById(id)
}

func (cs ConfigService) Get(id int64) types.Config {
	return types.Config{}
}

func (cs ConfigService) ListByPage(name string, pageNo int, pageSize int) []types.Config {
	return []types.Config{}
}
