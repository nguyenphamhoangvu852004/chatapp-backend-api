package setting

type Config struct {
	Redis      Redis      `mapstructure:"redis"`
	Mysql      Mysql      `mapstructure:"mysql"`
	Log        Log        `mapstructure:"log"`
	Server     Server     `mapstructure:"server"`
	Jwt        Jwt        `mapstructure:"jwt"`
	Cors       Cors       `mapstructure:"cors"`
	Cloudinary Cloudinary `mapstructure:"cloudinary"`
	Admin      Admin      `mapstructure:"admin"`
}
type Admin struct {
	Email       string `mapstructure:"email"`
	Password    string `mapstructure:"password"`
	PhoneNumber string `mapstructure:"phoneNumber"`
	Username    string `mapstructure:"username"`
}
type Cloudinary struct {
	CloudName string `mapstructure:"cloud_name"`
	ApiKey    string `mapstructure:"api_key"`
	ApiSecret string `mapstructure:"api_secret"`
}

type Cors struct {
	Url string `mapstructure:"url"`
}
type Jwt struct {
	AccessTokenSecret        string `mapstructure:"accessSecret"`
	AccessTokenExpiriedTime  int    `mapstructure:"accessSecretExpiriedTime"`
	RefreshTokenSecret       string `mapstructure:"refreshSecret"`
	RefreshTokenExpiriedTime int    `mapstructure:"refreshSecretExpiriedTime"`
}
type Server struct {
	Mode       string `mapstructure:"mode"`
	Port       int    `mapstructure:"port"`
	SocketPort int    `mapstructure:"socketPort"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type Mysql struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type Log struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}
