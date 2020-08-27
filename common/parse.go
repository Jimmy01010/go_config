package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

/**
* 	这个文件提供YAML和JSON的解析函数。
*/

// ParseYAMLBytes函数解析YAML编码的数据，并将结果存储在v指向的值中。
func ParseYAMLBytes(dat []byte, v interface{}) error {
	if err := yaml.Unmarshal(dat, v); err == io.EOF { // 文件为空返回EOF
		return fmt.Errorf("stream appears to be emtpy")
	} else if err != nil {
		return err
	}
	return nil
}

// ParseYAMLFile函数从yaml文件中读取配置。
func ParseYAMLFile(filename string, v interface{}) error {
	/**
	 * 因为yaml包没有实现Decoder，所以我们必须将整个文件读入内存，然后解析字节。
	 */
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("****************** Failed to read file '%v': %v", filename, err)
	}
	if err = ParseYAMLBytes(dat, v); err != nil {
		return fmt.Errorf("****************** failed to parse '%v': %v\n", filename, err)
	}
	return nil
}

// ParseJSONBytes函数解析json编码的数据，并将结果存储在v指向的值中。
func ParseJSONBytes(dat []byte, v interface{}) error {
	if err := json.Unmarshal(dat, v); err == io.EOF {
		return errors.New("stream appears to be emtpy")
	} else if err != nil {
		return err
	}
	return nil
}

// ParseJsonFile函数从json文件中读取配置。
func ParseJsonFile(filename string, v interface{}) error{
	file, err := os.Open(filename)
	if err != nil{
		return fmt.Errorf("****************** Failed to open file '%v': %v", filename, err)
	}
	defer file.Close()
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(v); err != nil{
		return fmt.Errorf("****************** failed to parse '%v': %v", filename, err)
	}
	return nil
}
