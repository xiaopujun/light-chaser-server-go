package global

import (
	"github.com/light-chaser/server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	LC_DB  *gorm.DB
	LC_ENV *model.LCENV
	LC_LOG *zap.Logger
)
