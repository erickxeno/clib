package config

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/erickxeno/clib/errors"
	"github.com/jinzhu/configor"
	"github.com/spf13/viper"
)

const (
	EnvKeyBMTInitOff = "BMT_INIT_OFF"
)

var (
	vipe       = newViper()
	done       uint32
	m          sync.Mutex
	BMTInitOff bool
)

func Init(configuredStructs ...interface{}) error {
	if atomic.LoadUint32(&done) != 0 {
		return nil
	}
	m.Lock()
	defer m.Unlock()
	if atomic.LoadUint32(&done) != 0 {
		return nil
	}
	defer atomic.StoreUint32(&done, 1)

	confFile, err := GetConfigYmlFile(FindConfDir(FolderCnf))
	if err != nil {
		return errors.WithStack(err)
	}
	return LoadFile(confFile, configuredStructs...)
}

func IsInitDone() bool {
	return atomic.LoadUint32(&done) != 0
}

func newViper() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}

// LoadFile config file (full path with name and extension)
func LoadFile(configFile string, configuredStructs ...interface{}) error {
	vipe.SetConfigFile(configFile)
	configor.New(&configor.Config{Verbose: true}).Load(configuredStructs, configFile)
	err := vipe.ReadInConfig()
	if err != nil {
		fmt.Printf("clib load config to vipe fail, file=%v, err=%v", configFile, err)
		return err
	}
	if len(configuredStructs) > 0 {
		err = configor.New(&configor.Config{Verbose: true}).Load(configuredStructs[0], configFile)
		if err != nil {
			fmt.Printf("clib load config to structure file=%v, err=%v", configFile, err)
			return err
		}
	}
	fmt.Printf("clib load config file=%v finish", configFile)
	return nil
}

func Has(key string) bool {
	return vipe.IsSet(key)
}

func Empty(key string) bool {
	return String(key) == ""
}

func Set(key string, val interface{}) {
	vipe.Set(key, val)
}

func SetDefault(key string, val interface{}) {
	vipe.SetDefault(key, val)
}

func String(key string) string {
	return vipe.GetString(key)
}
func DefaultString(key string, val string) string {
	if !vipe.IsSet(key) {
		return val
	}
	return vipe.GetString(key)
}

func Bool(key string) bool {
	return vipe.GetBool(key)
}
func DefaultBool(key string, val bool) bool {
	if Has(key) {
		return Bool(key)
	}
	return val
}

func Int(key string) int {
	return vipe.GetInt(key)
}
func DefaultInt(key string, val int) int {
	if Has(key) {
		return Int(key)
	}
	return val
}

func Int32(key string) int32 {
	return vipe.GetInt32(key)
}
func DefaultInt32(key string, val int32) int32 {
	if Has(key) {
		return Int32(key)
	}
	return val
}

func Int64(key string) int64 {
	return vipe.GetInt64(key)
}
func DefaultInt64(key string, val int64) int64 {
	if Has(key) {
		return Int64(key)
	}
	return val
}

func Float64(key string) float64 {
	return vipe.GetFloat64(key)
}
func DefaultFloat64(key string, val float64) float64 {
	if Has(key) {
		return Float64(key)
	}
	return val
}

func IntSlice(key string) []int {
	return vipe.GetIntSlice(key)
}
func DefaultIntSlice(key string, val []int) []int {
	if Has(key) {
		return IntSlice(key)
	}
	return val
}

func Time(key string) time.Time {
	return vipe.GetTime(key)
}
func DefaultTime(key string, val time.Time) time.Time {
	if Has(key) {
		return Time(key)
	}
	return val
}

func Duration(key string) time.Duration {
	return vipe.GetDuration(key)
}
func DefaultDuration(key string, val time.Duration) time.Duration {
	if Has(key) {
		return Duration(key)
	}
	return val
}

func StringSlice(key string) []string {
	return vipe.GetStringSlice(key)
}
func DefaultStringSlice(key string, val []string) []string {
	if Has(key) {
		return StringSlice(key)
	}
	return val
}

func StringMap(key string) map[string]interface{} {
	return vipe.GetStringMap(key)
}
func DefaultStringMap(key string, val map[string]interface{}) map[string]interface{} {
	if Has(key) {
		return StringMap(key)
	}
	return val
}

func StringMapString(key string) map[string]string {
	return vipe.GetStringMapString(key)
}
func DefaultStringMapString(key string, val map[string]string) map[string]string {
	if Has(key) {
		return StringMapString(key)
	}
	return val
}

func StringMapStringSlice(key string) map[string][]string {
	return vipe.GetStringMapStringSlice(key)
}
func DefaultStringMapStringSlice(key string, val map[string][]string) map[string][]string {
	if Has(key) {
		return StringMapStringSlice(key)
	}
	return val
}

func SizeInBytes(key string) uint {
	return vipe.GetSizeInBytes(key)
}
func DefaultSizeInBytes(key string, val uint) uint {
	if Has(key) {
		return SizeInBytes(key)
	}
	return val
}

func UnmarshalKey(key string, rawVal interface{}) error {
	return vipe.UnmarshalKey(key, rawVal)
}

func All() map[string]interface{} {
	return vipe.AllSettings()
}
