package setting

type Config struct {
	MYSQL  MysqlSetting  `mapstructure:"mysql"`
	LOGGER LoggerSetting `mapstructure:"log"`
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

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	LogFileName string `mapstructure:"logFileName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackup   int    `mapstructure:"maxBackup"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}
