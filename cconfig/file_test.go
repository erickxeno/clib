package cconfig

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/facebookgo/ensure"
)

var (
	ctx = context.Background()
)

func TestGetInt(t *testing.T) {
	str := `key: 12345`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetInt(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, int(12345))
}

func TestGetInt01(t *testing.T) {
	str := `{"key": 12345}`
	c, err := GetFileConfigFromJsonString(str)
	ensure.Nil(t, err)
	val, err := c.GetInt64(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, int64(12345))
}

func TestGetInt02(t *testing.T) {
	str := `{"key": "12345"}`
	c, err := GetFileConfigFromJsonString(str)
	ensure.Nil(t, err)
	val, err := c.GetString(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, "12345")
}

func TestGetInt03(t *testing.T) {
	str := `
key = 12345
`
	c, err := GetFileConfigFromTomlString(str)
	ensure.Nil(t, err)
	val, err := c.GetInt(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, int(12345))
}

func TestGetInt04(t *testing.T) {
	str := `
key = "12345"
`
	c, err := GetFileConfigFromTomlString(str)
	ensure.Nil(t, err)
	val, err := c.GetString(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, "12345")
}

func TestGetInt32(t *testing.T) {
	str := `key: 12345`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetInt32(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, int32(12345))
}

func TestGetInt64(t *testing.T) {
	str := `key: 12345`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetInt64(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, int64(12345))
}

func TestGetBool(t *testing.T) {
	str := `
key1: true
key2: 1
key3: false
key4: 0`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetBool(ctx, "key1")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, true)
	val, err = c.GetBool(ctx, "key2")
	ensure.NotNil(t, err)
	val, err = c.GetBool(ctx, "key3")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, false)
	val, err = c.GetBool(ctx, "key4")
	ensure.NotNil(t, err)
}

func TestGetFloat64(t *testing.T) {
	str := `key: 12345.0`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetFloat64(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, float64(12345.0))
}

func TestGetString(t *testing.T) {
	str := `key: "12345"`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetString(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val, "12345")
}

func TestGetStrings(t *testing.T) {
	str := `key: ["1", "2"]`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val, err := c.GetStrings(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, len(val), 2)
	ensure.DeepEqual(t, val[0], "1")
	ensure.DeepEqual(t, val[1], "2")
}

func TestGetStrings01(t *testing.T) {
	str := `{"key": ["1", "2"]}`
	c, err := GetFileConfigFromJsonString(str)
	ensure.Nil(t, err)
	val, err := c.GetStrings(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, len(val), 2)
	ensure.DeepEqual(t, val[0], "1")
	ensure.DeepEqual(t, val[1], "2")
}

func TestGetStrings02(t *testing.T) {
	str := `key= ["1", "2"]`
	c, err := GetFileConfigFromTomlString(str)
	ensure.Nil(t, err)
	val, err := c.GetStrings(ctx, "key")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, len(val), 2)
	ensure.DeepEqual(t, val[0], "1")
	ensure.DeepEqual(t, val[1], "2")
}

func TestGetJson(t *testing.T) {
	str := `
key:
  sub_key1: 1
  sub_key2: "34"`
	c, err := GetFileConfigFromYamlString(str)
	ensure.Nil(t, err)
	val := struct {
		SubKey1 int    `json:"sub_key1"`
		SubKey2 string `json:"sub_key2"`
	}{}
	err = c.GetJson(ctx, "key", &val)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val.SubKey1, 1)
	ensure.DeepEqual(t, val.SubKey2, "34")
}

func TestGetJson01(t *testing.T) {
	str := `
{
"key": {
  "sub_key1": 1,
  "sub_key2": "34"
 }
}
`
	c, err := GetFileConfigFromJsonString(str)
	ensure.Nil(t, err)
	val := struct {
		SubKey1 int    `json:"sub_key1"`
		SubKey2 string `json:"sub_key2"`
	}{}
	err = c.GetJson(ctx, "key", &val)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val.SubKey1, 1)
	ensure.DeepEqual(t, val.SubKey2, "34")
}

func TestGetJson02(t *testing.T) {
	str := `

[key] 
  sub_key1 = 1
  sub_key2 = "34"
`
	c, err := GetFileConfigFromTomlString(str)
	ensure.Nil(t, err)
	val := struct {
		SubKey1 int    `json:"sub_key1"`
		SubKey2 string `json:"sub_key2"`
	}{}
	err = c.GetJson(ctx, "key", &val)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val.SubKey1, 1)
	ensure.DeepEqual(t, val.SubKey2, "34")
}

func TestFileGet(t *testing.T) {
	str := `
{
	"key1": 1,
	"key2": "34",
	"key3": 11.11,
	"key4": 1111111111111111111,
	"key5": "1111111111111111111",
	"key": {
		"sub_key1": 1,
		"sub_key2": "34",
		"sub_key3": 11.11,
		"sub_key4": 1111111111111111111,
		"sub_key5": "1111111111111111111"
	}
}
`

	val := struct {
		SubKey1 int     `json:"sub_key1"`
		SubKey2 string  `json:"sub_key2"`
		SubKey3 float64 `json:"sub_key3"`
		SubKey4 int64   `json:"sub_key4"`
		SubKey5 string  `json:"sub_key5"`
	}{}

	file, err := ioutil.TempFile("./", "cconfig*.json")
	if err != nil {
		t.Fatal("Could not create temp file")
	}
	file.WriteString(str)
	file.Close()
	defer os.Remove(file.Name())
	c, err := GetFileConfig(file.Name())
	ensure.Nil(t, err)
	err = c.GetJson(ctx, "key", &val)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val.SubKey1, 1)
	ensure.DeepEqual(t, val.SubKey2, "34")
	ensure.DeepEqual(t, val.SubKey3, 11.11)
	ensure.DeepEqual(t, val.SubKey4, int64(1111111111111111111))
	ensure.DeepEqual(t, val.SubKey5, "1111111111111111111")

	val1, err := c.GetInt(ctx, "key1")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val1, 1)

	val11, err := c.GetBool(ctx, "key1")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val11, true)

	val2, err := c.GetString(ctx, "key2")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val2, "34")
	val3, err := c.GetFloat64(ctx, "key3")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val3, 11.11)
	val4, err := c.GetInt64(ctx, "key4")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val4, int64(1111111111111111111))
	val5, err := c.GetInt64(ctx, "key5")
	ensure.Nil(t, err)
	ensure.DeepEqual(t, val5, int64(1111111111111111111))
	defer os.Remove(file.Name())
}
