package test

import (
	"os"
	"reflect"
	"testing"
)

const Const_empty_string string = ""


/**
*  测试从文件中读取Yaml配置
*  go test -v -run=TestYamlConfig -args ./test.yaml
*/
func TestYamlConfig(t *testing.T)  {
	var config_file_path string = Const_empty_string

	/** 正式环境下使用下面代码来获取
	 *  flag.StringVar(&config_file_path, "config_file", "", "Path to configuration file.")
	 *  flag.Parse()
	 */

	// 获取配置文件路径
	config_file_path = os.Args[len(os.Args)-1]
	got, err := ReadConfigFromFile(config_file_path)
	if err != nil {
		t.Errorf("Error:%v'", err)
	}
	want := &Config{
		Application:     "go_config",
		ReadTimeoutSec:  300,
		WriteTimeoutSec: 900,
		BaseDir:         "/usr/local/cache",
		Test:			&TestS{
			"yellow",
			"chinese",
		},
		ConfigFile:      "./test.yaml",
		Port:	         8888,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%+v' want '%+v'", *got, *want)
	}
}

/**
*  测试从Json文件中读取配置
*  go test -v -run=TestJsonConfig -args ./test.json
*/
func TestJsonConfig(t *testing.T){
	var config_file_path string = Const_empty_string

	/** 正式环境下使用下面代码来获取
	 *  flag.StringVar(&config_file_path, "config_file", "", "Path to configuration file.")
	 *  flag.Parse()
	 */

	//获取配置文件路径
	config_file_path = os.Args[len(os.Args)-1]
	got, err := ReadConfigFromJsonFile(config_file_path)
	if err != nil {
		t.Errorf("Error:%v'", err)
	}
	want := &Config{
		Application:     "go_config",
		ReadTimeoutSec:  300,
		WriteTimeoutSec: 900,
		BaseDir:         "/usr/local/cache",
		Test:			&TestS{
			"yellow",
			"chinese",
		},
		ConfigFile:      "./test.json",
		Port:	         8888,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%+v' want '%+v'", *got, *want)
	}
}