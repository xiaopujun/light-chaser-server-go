package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	LC_DB     *gorm.DB
	LC_CONFIG map[string]any
	LC_LOG    *zap.Logger
)
