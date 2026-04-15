package config

import (
	"os"
	"strconv"
	"sync"
)

var (
	mu sync.RWMutex

	// Server configuration
	ServerPort = GetEnvOrDefault("SERVER_PORT", "3000")
	ServerAddress = GetEnvOrDefault("SERVER_ADDRESS", "0.0.0.0")

	// Database configuration
	SQLitePath = GetEnvOrDefault("SQLITE_PATH", "new-api.db")
	MySQLDSN = GetEnvOrDefault("SQL_DSN", "")

	// Redis configuration
	RedisConnString = GetEnvOrDefault("REDIS_CONN_STRING", "")
	RedisPassword = GetEnvOrDefault("REDIS_PASSWORD", "")

	// Session and security
	SessionSecret = GetEnvOrDefault("SESSION_SECRET", "new-api-secret")
	CryptoSecret = GetEnvOrDefault("CRYPTO_SECRET", "")

	// System settings
	DebugEnabled = GetEnvOrDefaultBool("DEBUG", false)
	LogDir = GetEnvOrDefault("LOG_DIR", "")

	// Rate limiting
	// Increased from 180 to 300 to be less restrictive for personal use
	GlobalApiRateLimitNum = GetEnvOrDefaultInt("GLOBAL_API_RATE_LIMIT", 300)
	GlobalApiRateLimitDuration int64 = 3 * 60 // seconds

	// Token and quota settings
	InitialRootToken = GetEnvOrDefault("INITIAL_ROOT_TOKEN", "")
	InitialRootAccessToken = GetEnvOrDefault("INITIAL_ROOT_ACCESS_TOKEN", "")

	// Frontend settings
	FrontendBaseURL = GetEnvOrDefault("FRONTEND_BASE_URL", "")

	// Worker concurrency
	WorkerCount = GetEnvOrDefaultInt("WORKER_COUNT", 4)

	// Version info (set at build time)
	Version = "v0.0.1"
	StartTime int64
)

// GetEnvOrDefault returns the value of the environment variable named by key,
// or defaultValue if the variable is not present or empty.
func GetEnvOrDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// GetEnvOrDefaultBool returns the boolean value of the environment variable,
// or defaultValue if not set or invalid.
func GetEnvOrDefaultBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}
	return b
}

// GetEnvOrDefaultInt returns the integer value of the environment variable,
// or defaultValue if not set or invalid.
func GetEnvOrDefaultInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return i
}

// GetEnvOrDefaultInt64 returns the int64 value of the environment variable,
// or defaultValue if not set or invalid.
func GetEnvOrDefaultInt64(key string, defaultValue int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultValue
	}
	return i
}

// IsMasterNode returns true if this instance is configured as the master node.
func IsMasterNode() bool {
	return GetEnvOrDefault("NODE_TYPE", "master") == "master"
}

// GetLock acquires the global config read-write mutex for writing.
func GetLock() {
	mu.Lock()
}

// ReleaseLock releases the global config read-write mutex.
func ReleaseLock() {
	mu.Unlock()
}
