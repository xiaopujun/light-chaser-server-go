package config

import (
	"github.com/light-chaser/server/global"
	"github.com/light-chaser/server/model"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

func LoadEvn() *model.LCENV {
	//初始化应用配置
	config := newConfig()
	//读取配置文件
	viper.SetConfigName("light-chaser")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//解析特殊配置项
	parseFileCapacityConfiguration(viper.GetViper(), "file_max_size")
	parseFileCapacityConfiguration(viper.GetViper(), "image_max_size")
	//绑定配置项到结构体
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}

// 解析文件、磁盘容量相关配置项。一般用于限制上传文件的尺寸大小。配置规范为：
// 1. 支持纯数字配置，单位为字节。如：1024，
// 2. 支持带单位配置，单位为KB、MB。如：1024KB，1024MB
// 3. 配置项为空时，默认为5MB
// 4. 若上述的其它情况，将抛出异常
func parseFileCapacityConfiguration(viper *viper.Viper, key string) {
	if viper == nil || key == "" {
		panic("Unable to resolve, parser exception, or key is empty")
	}
	capacityValue := viper.Get(key)
	if value, ok := capacityValue.(string); ok {
		value = strings.ToUpper(value)
		mb := strings.HasSuffix(value, "MB")
		kb := strings.HasSuffix(value, "KB")
		if mb || kb {
			num, err := strconv.Atoi(strings.TrimSpace(value[:len(value)-2]))
			if err != nil {
				panic(err)
			} else {
				if mb {
					num = num * 1024 * 1024 * 1024
				}
				if kb {
					num = num * 1024 * 1024
				}
				viper.Set(key, int64(num))
			}
		} else {
			panic(key + " is not valid : " + value)
		}
	} else if value, ok := capacityValue.(int64); ok {
		viper.Set(key, value)
	} else {
		global.LC_LOG.Info(key + " was unplaced")
	}
}

func newConfig() *model.LCENV {
	env := model.LCENV{
		Separator:              "/",
		AllowFileType:          []string{".js"},
		AllowImageType:         []string{"jpg", "jpeg", "png", "gif"},
		FileMaxSize:            1024 * 1024 * 1024 * 5,
		ImageMaxSize:           1024 * 1024 * 1024 * 5,
		RootPath:               ".",
		ImagePath:              "./images/",
		CoverPath:              "./covers/",
		AvatarPath:             "./avatars/",
		RemoteComponentPath:    "./remote_components/",
		LicensePath:            "./license/",
		TestDatabaseMaxTimeout: 10,
		Server: model.Server{
			Port: 8080,
			Mode: "debug",
		},
	}
	return &env
}
