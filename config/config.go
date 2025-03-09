package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Server        ServerConfig       `envconfig:"SERVER"`
	Database      DatabaseConfig     `envconfig:"DATABASE"`
	Authorization Authorization      `envconfig:"AUTHORIZATION"`
	Logger        Logger             `envconfig:"LOGGER"`
	Redis         RedisConfig        `envconfig:"REDIS"`
	RedisCluster  RedisClusterConfig `envconfig:"REDIS_CLUSTER"`
	Minio         MinioConfig        `envconfig:"MINIO"`
}

type ServerConfig struct {
	Name                string        `envconfig:"NAME" default:"MyServer"`
	AppVersion          string        `envconfig:"APP_VERSION" default:"1.0.0"`
	Port                string        `envconfig:"PORT" default:"8080"`
	BaseURI             string        `envconfig:"BASE_URI" default:""`
	Mode                string        `envconfig:"MODE" default:"debug"`
	ReadTimeout         time.Duration `envconfig:"READ_TIMEOUT" default:"30s"`
	WriteTimeout        time.Duration `envconfig:"WRITE_TIMEOUT" default:"30s"`
	SSL                 bool          `envconfig:"SSL" default:"false"`
	CSRF                bool          `envconfig:"CSRF" default:"false"`
	Debug               bool          `envconfig:"DEBUG" default:"true"`
	GrRunningThreshold  int           `envconfig:"GR_RUNNING_THRESHOLD" default:"100"`
	GcPauseThreshold    int           `envconfig:"GC_PAUSE_THRESHOLD" default:"200"`
	CacheDeploymentType int           `envconfig:"CACHE_DEPLOYMENT_TYPE" default:"1"`
}

type DatabaseConfig struct {
	Host            string `envconfig:"HOST" default:"localhost"`
	Port            string `envconfig:"PORT" default:"5432"`
	User            string `envconfig:"USER" default:"postgres"`
	Password        string `envconfig:"PASSWORD" default:""`
	Database        string `envconfig:"NAME" default:"hotel_reservation"`
	MaxPoolOpen     int    `envconfig:"MAX_POOL_OPEN" default:"50"`
	MaxPoolIdle     int    `envconfig:"MAX_POOL_IDLE" default:"10"`
	MaxPollLifeTime int    `envconfig:"MAX_POLL_LIFETIME" default:"300000"`
	LogLevel        string `envconfig:"LOG_LEVEL" default:"SILENT"` // Options: "SILENT", "INFO", "WARN", "ERROR"
}
type Authorization struct {
	JWTSecret           string `envconfig:"JWT_SECRET" default:"ais-jwt"`
	JwtExpired          int    `envconfig:"JWT_EXPIRATION" default:"3600"`
	RefreshTokenExpired int    `envconfig:"JWT_REFRESH_EXPIRATION" default:"360000"`
}

type Logger struct {
	Development       bool   `envconfig:"DEVELOPMENT" default:"true"`
	DisableCaller     bool   `envconfig:"DISABLE_CALLER" default:"false"`
	DisableStacktrace bool   `envconfig:"DISABLE_STACKTRACE" default:"false"`
	Encoding          string `envconfig:"ENCODING" default:"json"`
	Level             string `envconfig:"LEVEL"`
}

type RedisConfig struct {
	Address     string `envconfig:"REDIS_ADDRESS"`
	Password    string `envconfig:"REDIS_PASSWORD"`
	DefaultDb   string `envconfig:"REDIS_DEFAULT_DB"`
	MinIdleCons int    `envconfig:"REDIS_MIN_IDLE_CONNS"`
	PoolSize    int    `envconfig:"REDIS_POOL_SIZE"`
	PoolTimeout int    `envconfig:"REDIS_POOL_TIMEOUT"`
	DB          int    `envconfig:"REDIS_DB"`
}

type RedisClusterConfig struct {
	Delimiter   string `envconfig:"REDIS_CLUSTER_DELIMITER"`
	ReadOnly    bool   `envconfig:"REDIS_CLUSTER_READ_ONLY"`
	Address     string `envconfig:"REDIS_CLUSTER_ADDRESS"`
	DefaultDb   string `envconfig:"REDIS_CLUSTER_DEFAULT_DB"`
	MinIdleCons int    `envconfig:"REDIS_CLUSTER_MIN_IDLE_CONNS"`
	PoolSize    int    `envconfig:"REDIS_CLUSTER_POOL_SIZE"`
	PoolTimeout int    `envconfig:"REDIS_CLUSTER_POOL_TIMEOUT"`
	Password    string `envconfig:"REDIS_CLUSTER_PASSWORD"`
	DB          int    `envconfig:"REDIS_CLUSTER_DB"`
}

type MinioConfig struct {
	Endpoint  string `envconfig:"ENDPOINT" default:"http://localhost"`
	Port      string `envconfig:"PORT" default:"9000"`
	AccessKey string `envconfig:"ACCESS_KEY" default:"admin"`
	SecretKey string `envconfig:"SECRET_KEY" default:"admin123"`
	Bucket    string `envconfig:"BUCKET" default:"hotel-storage"`
	UseSSL    bool   `envconfig:"USE_SSL" default:"false"`
}

func NewConfig() (*Configuration, error) {
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	var cfg Configuration
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration from environment variables: %v", err)
	}

	return &cfg, nil
}
