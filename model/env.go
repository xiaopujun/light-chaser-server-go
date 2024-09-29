package model

type DataSource struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
}

type Server struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LCENV struct {
	Separator              string     `mapstructure:"separator"`
	AllowFileType          []string   `mapstructure:"allow_file_type"`
	AllowImageType         []string   `mapstructure:"allow_image_type"`
	FileMaxSize            int64      `mapstructure:"file_max_size"`
	ImageMaxSize           int64      `mapstructure:"image_max_size"`
	RootPath               string     `mapstructure:"root_path"`
	ImagePath              string     `mapstructure:"image_path"`
	CoverPath              string     `mapstructure:"cover_path"`
	AvatarPath             string     `mapstructure:"avatar_path"`
	RemoteComponentPath    string     `mapstructure:"remote_component_path"`
	LicensePath            string     `mapstructure:"license_path"`
	TestDatabaseMaxTimeout int64      `mapstructure:"test_database_max_timeout"`
	Database               DataSource `mapstructure:"database"`
	Server                 Server     `mapstructure:"server"`
}
