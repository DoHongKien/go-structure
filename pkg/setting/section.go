package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Logger LoggerSetting `mapstructure:"logger"`
	JWT    JWTSetting    `mapstructure:"jwt"`
}

type ServerSetting struct {
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int `mapstructure:"db"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxSize     int    `mapstructure:"max_size"`
	Compress    bool   `mapstructure:"compress"`
}

type JWTSetting struct {
	SecretKey  string `mapstructure:"secret_key"`
	ExpireTime string    `mapstructure:"expire_time"`
	RefreshTime string `mapstructure:"refresh_time"`
}
