package models

type LogType string

const (
	INFO  LogType = "INFO"
	ERROR LogType = "ERROR"
	WARN  LogType = "WARN"
	DEBUG LogType = "DEBUG"
)

type LogInfo struct {
	Dir   string  `yaml:"dir"`
	File  string  `yaml:"file"`
	Level LogType `yaml:"level"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
	Driver   string `yaml:"driver"`
	EndPoint string `yaml:"endPoint"`
}

type RedisConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Password     string `yaml:"password"`
	Database     int    `yaml:"database"`
	DialTimeout  int    `yaml:"dialTimeout"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
	PoolSize     int    `yaml:"poolSize"`
	EndPoint     string `yaml:"endpoint"`
}

type Config struct {
	Env  string `yaml:"env"`
	Http struct {
		WriteTimeout int `yaml:"writeTimeout"`
		ReadTimeout  int `yaml:"readTimeout"`
	}
	Log   LogInfo     `yaml:"log"`
	Db    DBConfig    `yaml:"db"`
	Redis RedisConfig `yaml:"redis"`
}
