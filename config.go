package test

import (
	"log"
	"path/filepath"
	"runtime"

	config "github.com/mateothegreat/go-config"
	"github.com/mateothegreat/go-config/plugins/sources"
	"github.com/mateothegreat/go-multilog/multilog"
)

var TestConfig *Conf

func init() {
	TestConfig = Setup()
}

type Conf struct {
	NotionAPIKey   string `validate:"required" yaml:"notion_api_key"`
	RedisAddress   string `validate:"required" yaml:"redis_address"`
	RedisDatabase  int    `validate:"required" yaml:"redis_database"`
	RedisKeyPrefix string `validate:"required" yaml:"redis_key_prefix"`
	RedisUsername  string `validate:"required" yaml:"redis_username"`
	RedisPassword  string `validate:"required" yaml:"redis_password"`
}

func Setup() *Conf {
	cfg := &Conf{}

	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  multilog.DEBUG,
		Format: multilog.FormatText,
	}))

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "config.yaml")

	err := config.LoadWithPlugins(
		config.FromYAML(sources.YAMLOpts{Path: dir}),
		config.FromEnv(sources.EnvOpts{Prefix: "APP"}),
	).WithValidationStrategy(config.StrategyAuto).Build(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
