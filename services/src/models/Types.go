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

type Config struct {
	Env  string `yaml:"env"`
	Http struct {
		WriteTimeout int `yaml:"writeTimeout"`
		ReadTimeout  int `yaml:"readTimeout"`
	}
	Log LogInfo  `yaml:"log"`
	Db  DBConfig `yaml:"db"`
}
