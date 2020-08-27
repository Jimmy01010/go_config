package test

import (
	"go_config/common"
)

/*
* 	这个文件定义配置文件有关的结构，并提供一些查询方法和一个从文件中读取配置的函数。
*/

type TestS struct {
	Colour string
	Language string
}

type Config struct {
	Application string `yaml:"Application"`
	WriteTimeoutSec int `yaml:"write_timeout_sec"`
	ReadTimeoutSec  int `yaml:"read_timeout_sec"`

	Test *TestS `yaml:"test"`
	ConfigFile string
	Port       int `yaml:"port"`
	BaseDir    string `yaml:"base_dir"`
}

/**
 * 对于一些不想被修改的字段，只提供一个查询方法
 */
func (cnfg *Config) GetBaseDir() string {
	return cnfg.BaseDir
}

func (cnfg *Config) GetTest() *TestS {
	return cnfg.Test
}

/**
 * 返回一个带有默认值的新Config结构体。
 */
func NewConfig(configFile string) *Config {
	return &Config{
		ConfigFile:      configFile,
		ReadTimeoutSec:  300,
		WriteTimeoutSec: 900,
		BaseDir:         "/usr/local/cache",
	}
}


/**
 * 从文件中读取配置
 */
func ReadConfigFromFile(config_file_path string) (*Config, error) {
	config := NewConfig(config_file_path)
	if err := common.ParseYAMLFile(config_file_path, config); err != nil {
		return nil, err
	}

	// 确保目录（BaseDir）存在，不存在则创建
	if config.GetBaseDir() != "" {
		if err := common.CheckAndMakeDir(config.GetBaseDir()); err != nil {
			return nil, err
		}
	}
	return config, nil
}

func ReadConfigFromJsonFile(config_file_path string) (*Config, error) {
	config := NewConfig(config_file_path)
	if err := common.ParseJsonFile(config_file_path, config); err != nil {
		return nil, err
	}

	// 确保目录存在，不存在则创建
	if config.GetBaseDir() != "" {
		if err := common.CheckAndMakeDir(config.GetBaseDir()); err != nil {
			return nil, err
		}
	}
	return config, nil
}