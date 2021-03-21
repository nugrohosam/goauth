package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	helpers "github.com/nugrohosam/gosampleapi/helpers"
	database "github.com/nugrohosam/gosampleapi/services/databases"
	grpcConn "github.com/nugrohosam/gosampleapi/services/grpc"
	httpConn "github.com/nugrohosam/gosampleapi/services/http"
	infrastructure "github.com/nugrohosam/gosampleapi/services/infrastructure"
	"github.com/spf13/viper"
)

func main() {

	envRootPath := flag.String("env-root-path", "none", "--")
	serviceUse := flag.String("service", "none", "--")

	flag.Parse()

	if *envRootPath == "none" {
		fmt.Println("flag [--env-root-path=?] must be spellied")
		return
	}

	if !helpers.InArray(*serviceUse, []string{"http", "grpc"}) || *serviceUse == "none" {
		fmt.Println("flag [--service=?] must be spellied in (http or grpc)")
		return
	}

	loadConfigFile(*envRootPath)

	infrastructure.PrepareSentry()

	if err := database.ConnOrm(); err != nil {
		panic(err)
	}

	if *serviceUse == "grpc" {
		if err := grpcConn.Serve(); err != nil {
			panic(err)
		}
	} else if *serviceUse == "http" {
		if err := httpConn.Serve(); err != nil {
			panic(err)
		}
	}
}

func initiateRedisCache() {
	cacheRedisPrefixKey := viper.GetString("cache.redis.prefix-key")
	driverCache := viper.GetString("redis.driver")

	switch driverCache {
	case "redis":
		configCacheRedis := viper.GetStringMap("redis")
		redisHostsCache := make(map[string]string)
	
		for key, value := range configCacheRedis {
			keyRedis := cacheRedisPrefixKey + key
			valueReal := value.(map[string]string)
			redisHostsCache[keyRedis] = valueReal["host"] + ":" + valueReal["port"]
		}
	
		infrastructure.InitiateRedisCache(redisHostsCache)
	}
}

func loadConfigFile(envRootPath string) {

	viper.SetConfigType("yaml")
	viper.SetConfigName(".env")
	viper.AddConfigPath(envRootPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var files []string
	configFolderName := "config"
	rootPathConfig := helpers.SetPath(envRootPath, configFolderName)

	if err := filepath.Walk(rootPathConfig, func(path string, info os.FileInfo, err error) error {
		if info.Name() != configFolderName {
			files = append(files, info.Name())
		}
		return nil
	}); err != nil {
		panic(err)
	}

	var nameConfig string

	for _, file := range files {
		nameConfig = strings.ReplaceAll(file, ".yaml", "")

		viper.SetConfigName(nameConfig)
		viper.AddConfigPath(rootPathConfig)

		if err := viper.MergeInConfig(); err != nil {
			panic(err)
		}
	}
}
