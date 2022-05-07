package interfaces

import "fire-press/api/types"

type ConfigService interface {
	Save(config types.Config)
	Delete(id int64) int64
	Get(id int64) types.Config
	ListByPage(name string, pageNo int, pageSize int) []types.Config
}
