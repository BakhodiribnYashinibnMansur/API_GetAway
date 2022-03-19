package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config
type Config struct {
	Environment string // develop, staging, production

	PositionServiceHost string
	PositionServicePort int

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string
	MinioBucketName  string
	MinioLocation    string
	MinioHost        string

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	config.PositionServiceHost = cast.ToString(getOrReturnDefault("POSITION_SERVICE_HOST", "localhost"))
	config.PositionServicePort = cast.ToInt(getOrReturnDefault("POSITION_SERVICE_PORT", 5004))

	config.MinioEndpoint = cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", ""))
	config.MinioAccessKeyID = cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", ""))
	config.MinioSecretKey = cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY_ID", ""))
	config.MinioBucketName = cast.ToString(getOrReturnDefault("MINIO_BACKET_NAME", ""))
	config.MinioLocation = cast.ToString(getOrReturnDefault("MINIO_LOCATION", ""))
	config.MinioHost = cast.ToString(getOrReturnDefault("MINIO_HOST", ""))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
