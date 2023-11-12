package config

import (
	"time"

	"github.com/spf13/viper"
)

type config struct {
	Web struct {
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		IdleTimeout     time.Duration
		ShutdownTimeout time.Duration
		APIHost         string
		DebugHost       string
	}
	DB struct {
		User         string
		Password     string
		Host         string
		Name         string
		Schema       string
		MaxIdleConns int
		MaxOpenConns int
		DisableTLS   bool
	}
	Redis struct {
		Host struct {
			Master  string
			Replica string
		}
		Password string
		DB       int
	}
	Tempo struct {
		ReporterURI string
		ServiceName string
		Probability float64
	}
	Auth struct {
		Audience string
	}
}

func New() (config, error) {
	vConfig := viper.New()

	// Set Web defaults.
	vConfig.SetDefault("Web.ReadTimeout", time.Second*15)
	vConfig.SetDefault("Web.WriteTimeout", time.Second*15)
	vConfig.SetDefault("Web.IdleTimeout", time.Second*60)
	vConfig.SetDefault("Web.ShutdownTimeout", time.Second*5)
	vConfig.SetDefault("Web.APIHost", "0.0.0.0:3000")
	vConfig.SetDefault("Web.DebugHost", "0.0.0.0:4000")

	// Set DB defaults.
	vConfig.SetDefault("DB.User", "postgres")
	vConfig.SetDefault("DB.Password", "postgres")
	vConfig.SetDefault("DB.Host", "localhost:5432")
	vConfig.SetDefault("DB.Name", "postgres")
	vConfig.SetDefault("DB.Schema", "public")
	vConfig.SetDefault("DB.MaxIdleConns", 10)
	vConfig.SetDefault("DB.MaxOpenConns", 10)
	vConfig.SetDefault("DB.DisableTLS", true)

	// Set Redis defaults.
	vConfig.SetDefault("Redis.Host.Master", "redis-master.tuber-system.svc.cluster.local:6379")
	vConfig.SetDefault("Redis.Host.Replica", "redis-replicas.tuber-system.svc.cluster.local:6379")
	vConfig.SetDefault("Redis.Password", "redis")
	vConfig.SetDefault("Redis.DB", 0)

	// Set Tempo defaults.
	vConfig.SetDefault("Tempo.ReporterURI", "tempo.tuber-system.svc.cluster.local:4317")
	vConfig.SetDefault("Tempo.ServiceName", "tuber-api")
	vConfig.SetDefault("Tempo.Probability", 1)

	// Set Auth defaults.
	vConfig.SetDefault("Auth.Audience", "808161745758-73upckm7gpah36uejvar457fhn6qtsd4.apps.googleusercontent.com")

	// Enable environment variable overriding for all.
	vConfig.AutomaticEnv()
	conf := &config{}
	// Unmarshal the config into the conf struct.
	if err := vConfig.Unmarshal(conf); err != nil {
		return *conf, err
	}
	return *conf, nil
}
