package cconfig

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/erickxeno/clib/convert"
	"github.com/erickxeno/clib/errors"
	"github.com/jinzhu/configor"

	"github.com/BurntSushi/toml"
	_jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

var (
	jsoniter = _jsoniter.ConfigCompatibleWithStandardLibrary
)

// FileConfig Configuration base file or string
type FileConfig struct {
	data map[string]interface{}
}

// GetFileConfigFromYamlStream get Configuration from io stream(yaml format)
func GetFileConfigFromYamlStream(r io.Reader) (Configuration, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return GetFileConfigFromYamlString(string(d))
}

// GetFileConfigFromYamlString get Configuration from string(yaml format)
func GetFileConfigFromYamlString(str string) (Configuration, error) {
	data := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(str), &data)
	if err != nil {
		return nil, err
	}
	return &FileConfig{data: data}, nil
}

// GetFileConfigFromJsonStream  get Configuration from io Stream(json format)
func GetFileConfigFromJsonStream(r io.Reader) (Configuration, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return GetFileConfigFromJsonString(string(d))
}

// GetFileConfigFromJsonString get Configuration from string(json format)
func GetFileConfigFromJsonString(str string) (Configuration, error) {
	data := make(map[string]interface{})
	decoder := jsoniter.NewDecoder(bytes.NewBufferString(str))
	decoder.UseNumber()
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return &FileConfig{data: data}, nil
}

// GetFileConfigFromTomlStream  get Configuration from io Stream(Toml format)
func GetFileConfigFromTomlStream(r io.Reader) (Configuration, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return GetFileConfigFromTomlString(string(d))
}

// GetFileConfigFromTomlString get Configuration from string(Toml format)
func GetFileConfigFromTomlString(str string) (Configuration, error) {
	data := make(map[string]interface{})
	err := toml.Unmarshal([]byte(str), &data)
	if err != nil {
		return nil, err
	}
	return &FileConfig{data: data}, nil
}

// GetFileConfig get Configuration from file(normal k-v format)
func GetFileConfig(fileNames ...string) (Configuration, error) {
	if len(fileNames) == 1 && strings.HasSuffix(fileNames[0], ".json") {
		f, err := os.Open(fileNames[0])
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return GetFileConfigFromJsonStream(f)
	}
	data := make(map[string]interface{})
	err := configor.New(&configor.Config{Verbose: true}).Load(&data, fileNames...)
	if err != nil && !strings.Contains(err.Error(), "invalid config, should be struct") {
		return nil, err
	}
	return &FileConfig{data: data}, nil
}

// GetConfigItem get sub child Configuration from a FileConfig
func GetConfigItem(c Configuration, item string) Configuration {
	yc, ok := c.(*FileConfig)
	if !ok {
		return nil
	}
	if v, ok := yc.data[item]; ok {
		if d, ok := v.(map[string]interface{}); ok {
			return &FileConfig{data: d}
		}
	}
	return nil
}

// GetInt Get the value by key from FileConfig, value format int
func (fc *FileConfig) GetInt(ctx context.Context, key string) (int, error) {
	if v, ok := fc.data[key].(int); ok {
		return v, nil
	}
	if v, ok := fc.data[key].(int32); ok {
		return int(v), nil
	}
	if v, ok := fc.data[key].(int64); ok {
		return int(v), nil
	}
	if v, ok := fc.data[key].(_jsoniter.Number); ok {
		tmp, err := v.Int64()
		return int(tmp), err
	}
	if v, ok := fc.data[key].(json.Number); ok {
		tmp, err := v.Int64()
		return int(tmp), err
	}
	if v, ok := fc.data[key].(float64); ok {
		return int(v), nil
	}
	if v, ok := fc.data[key].(string); ok {
		return convert.ToIntE(v)
	}
	return 0, errors.Errorf("notint value")
}

// GetInt32 Get the value by key from FileConfig, value format int32
func (fc *FileConfig) GetInt32(ctx context.Context, key string) (int32, error) {
	if v, ok := fc.data[key].(int32); ok {
		return v, nil
	}
	if v, ok := fc.data[key].(int); ok {
		return int32(v), nil
	}
	if v, ok := fc.data[key].(int64); ok {
		return int32(v), nil
	}
	if v, ok := fc.data[key].(_jsoniter.Number); ok {
		tmp, err := v.Int64()
		return int32(tmp), err
	}
	if v, ok := fc.data[key].(json.Number); ok {
		tmp, err := v.Int64()
		return int32(tmp), err
	}
	if v, ok := fc.data[key].(float64); ok {
		return int32(v), nil
	}
	if v, ok := fc.data[key].(string); ok {
		return convert.ToInt32E(v)
	}
	return 0, errors.Errorf("not int32 value")
}

// GetInt64 Get the value by key from FileConfig, value format int64
func (fc *FileConfig) GetInt64(ctx context.Context, key string) (int64, error) {
	if v, ok := fc.data[key].(int64); ok {
		return v, nil
	}
	if v, ok := fc.data[key].(int32); ok {
		return int64(v), nil
	}
	if v, ok := fc.data[key].(int); ok {
		return int64(v), nil
	}
	if v, ok := fc.data[key].(_jsoniter.Number); ok {
		return v.Int64()
	}
	if v, ok := fc.data[key].(json.Number); ok {
		return v.Int64()
	}
	if v, ok := fc.data[key].(float64); ok {
		return int64(v), nil
	}
	if v, ok := fc.data[key].(string); ok {
		return convert.ToInt64E(v)
	}
	return 0, errors.Errorf("not int64 value")
}

// GetBool Get the value by key from FileConfig, value format bool
func (fc *FileConfig) GetBool(ctx context.Context, key string) (bool, error) {
	if v, ok := fc.data[key].(bool); ok {
		return v, nil
	}
	if v, ok := fc.data[key].(string); ok {
		return convert.ToBoolE(v)
	}
	if v, ok := fc.data[key].(_jsoniter.Number); ok {
		return convert.ToBoolE(v.String())
	}
	if v, ok := fc.data[key].(json.Number); ok {
		return convert.ToBoolE(v.String())
	}
	return false, errors.Errorf("not bool value")
}

// GetFloat64 Get the value by key from FileConfig, value format float64
func (fc *FileConfig) GetFloat64(ctx context.Context, key string) (float64, error) {
	if v, ok := fc.data[key].(float64); ok {
		return v, nil
	}
	if v, ok := fc.data[key].(_jsoniter.Number); ok {
		return v.Float64()
	}
	if v, ok := fc.data[key].(json.Number); ok {
		return v.Float64()
	}
	if v, ok := fc.data[key].(string); ok {
		return convert.ToFloat64E(v)
	}
	return 0.0, errors.Errorf("not float64 value")
}

// GetString Get the value by key from FileConfig, value format string
func (fc *FileConfig) GetString(ctx context.Context, key string) (string, error) {
	if v, ok := fc.data[key].(string); ok {
		return v, nil
	}
	return "", errors.Errorf("not string value")
}

// GetStrings Get the value by key from FileConfig, value format []string
func (fc *FileConfig) GetStrings(ctx context.Context, key string) ([]string, error) {
	val := []string{}
	return val, fc.GetJson(ctx, key, &val)
}

// GetJson Get the value by key from FileConfig, value format json
func (fc *FileConfig) GetJson(ctx context.Context, key string, val interface{}) error {
	defaultValue := reflect.Indirect(reflect.ValueOf(val))
	if !defaultValue.CanAddr() {
		return errors.Errorf("val %v should be addressable", val)
	}
	if v, ok := fc.data[key]; ok {
		bytes, err := jsoniter.Marshal(v)
		if err != nil {
			return err
		}
		return jsoniter.Unmarshal(bytes, val)
	}
	return errors.Errorf("key not exist")
}

// GetIntWithDefault Get the value by key from FileConfig, value format int, return default val when err happen
func (fc *FileConfig) GetIntWithDefault(ctx context.Context, key string, defaultVal int) int {
	val, err := fc.GetInt(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetInt32WithDefault Get the value by key from FileConfig, value format int32, return default val when err happen
func (fc *FileConfig) GetInt32WithDefault(ctx context.Context, key string, defaultVal int32) int32 {
	val, err := fc.GetInt32(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetInt64WithDefault Get the value by key from FileConfig, value format int64, return default val when err happen
func (fc *FileConfig) GetInt64WithDefault(ctx context.Context, key string, defaultVal int64) int64 {
	val, err := fc.GetInt64(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetBoolWithDefault Get the value by key from FileConfig, value format bool, return default val when err happen
func (fc *FileConfig) GetBoolWithDefault(ctx context.Context, key string, defaultVal bool) bool {
	val, err := fc.GetBool(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetFloat64WithDefault Get the value by key from FileConfig, value format float64, return default val when err happen
func (fc *FileConfig) GetFloat64WithDefault(ctx context.Context, key string, defaultVal float64) float64 {
	val, err := fc.GetFloat64(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetStringWithDefault Get the value by key from FileConfig, value format string, return default val when err happen
func (fc *FileConfig) GetStringWithDefault(ctx context.Context, key string, defaultVal string) string {
	val, err := fc.GetString(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}

// GetStringsWithDefault Get the value by key from FileConfig, value format []string, return default val when err happen
func (fc *FileConfig) GetStringsWithDefault(ctx context.Context, key string, defaultVal []string) []string {
	val, err := fc.GetStrings(ctx, key)
	if err != nil {
		return defaultVal
	}
	return val
}
