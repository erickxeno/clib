package config

import (
	"os"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

type Configure struct {
	Member   string   `yaml:"Member"`
	Elements []string `yaml:"Elements"`
	St       struct {
		Member   string   `yaml:"Member"`
		Elements []string `yaml:"Elements"`
	} `yaml:"Struct"`
	First struct {
		Second struct {
			Member   string   `yaml:"Member"`
			Elements []string `yaml:"Elements"`
		} `yaml:"Second"`
	} `yaml:"First"`
	NoExit string `yaml:"NoExit"`
}

func TestFindConfDir(t *testing.T) {
	os.Setenv(CiEnvName, CiEnvValue)
	PatchConvey("test FindConfDir", t, func() {
		PatchConvey("指定具体目录", func() {
			path := FindConfDir("conf")
			So(path, ShouldResemble, "../conf")
		})
		PatchConvey("指定level", func() {
			path := FindConfDir("conf", 10)
			So(path, ShouldResemble, "")
		})
		PatchConvey("不存在", func() {
			path := FindConfDir("aaa")
			So(path, ShouldResemble, "")
		})
	})
}

func TestGetConfigYmlFile(t *testing.T) {
	os.Setenv(CiEnvName, CiEnvValue)
	PatchConvey("GetConfigYmlFile测试", t, func() {
		PatchConvey("指定具体目录", func() {
			cf, err := GetConfigYmlFile("../conf")
			So(err, ShouldBeNil)
			So(cf, ShouldResemble, "../conf/ci.yml")
		})
		PatchConvey("当前目录", func() {
			cf, err := GetConfigYmlFile(".")
			So(err, ShouldNotBeNil)
			So(cf, ShouldResemble, "")
		})
		PatchConvey("不存在的目录", func() {
			cf, err := GetConfigYmlFile("aaa")
			So(err, ShouldNotBeNil)
			So(cf, ShouldResemble, "")
		})
	})
}

func TestLoadFile(t *testing.T) {
	os.Setenv(CiEnvName, CiEnvValue)
	PatchConvey("Test LoadFile", t, func() {
		PatchConvey("no configure struct", func() {
			err := LoadFile("../conf/ci.yml")
			So(err, ShouldBeNil)
		})
		PatchConvey("load with configure struct", func() {
			cfg := Configure{}
			err := LoadFile("../conf/ci.yml", &cfg)
			So(err, ShouldBeNil)
			So(cfg.Member, ShouldResemble, "A1")
			So(len(cfg.Elements), ShouldEqual, 2)
			So(cfg.Elements[0], ShouldResemble, "E1")
			So(cfg.Elements[1], ShouldResemble, "E2")

			So(cfg.St.Member, ShouldEqual, "A2")
			So(len(cfg.St.Elements), ShouldEqual, 2)
			So(cfg.St.Elements[0], ShouldResemble, "E11")
			So(cfg.St.Elements[1], ShouldResemble, "E12")

			So(cfg.First.Second.Member, ShouldEqual, "A3")
			So(len(cfg.First.Second.Elements), ShouldEqual, 2)
			So(cfg.First.Second.Elements[0], ShouldResemble, "E111")
			So(cfg.First.Second.Elements[1], ShouldResemble, "E112")

			So(cfg.NoExit, ShouldResemble, "")
		})
	})
}

func TestInit(t *testing.T) {
	os.Setenv(CiEnvName, CiEnvValue)
	PatchConvey("Test Init", t, func() {
		PatchConvey("load with configure struct", func() {
			cfg := Configure{}
			err := Init(&cfg)
			So(err, ShouldBeNil)
			So(cfg.Member, ShouldResemble, "A1")
			So(len(cfg.Elements), ShouldResemble, 2)
			So(cfg.Elements[0], ShouldResemble, "E1")
			So(cfg.Elements[1], ShouldResemble, "E2")

			So(cfg.St.Member, ShouldResemble, "A2")
			So(len(cfg.St.Elements), ShouldResemble, 2)
			So(cfg.St.Elements[0], ShouldResemble, "E11")
			So(cfg.St.Elements[1], ShouldResemble, "E12")

			So(cfg.First.Second.Member, ShouldResemble, "A3")
			So(len(cfg.First.Second.Elements), ShouldResemble, 2)
			So(cfg.First.Second.Elements[0], ShouldResemble, "E111")
			So(cfg.First.Second.Elements[1], ShouldResemble, "E112")
			So(cfg.NoExit, ShouldResemble, "")
		})
	})
}

func TestBaseFunc(t *testing.T) {
	os.Setenv(CiEnvName, CiEnvValue)
	PatchConvey("Test Base Func", t, func() {
		PatchConvey("load with no configure struct", func() {
			err := LoadFile("../conf/ci.yml")
			So(err, ShouldBeNil)

			So(Has("Member"), ShouldBeTrue)
			So(Has("Elements"), ShouldBeTrue)
			So(Has("Struct"), ShouldBeTrue)
			So(Has("Struct.Member"), ShouldBeTrue)
			So(Has("Struct.Elements"), ShouldBeTrue)
			So(Has("First"), ShouldBeTrue)
			So(Has("First.Second"), ShouldBeTrue)
			So(Has("First.Second.Member"), ShouldBeTrue)
			So(Has("First.Second.Elements"), ShouldBeTrue)

			So(String("Member"), ShouldResemble, "A1")
			So(len(StringSlice("Elements")), ShouldEqual, 2)
			So(StringSlice("Elements")[0], ShouldResemble, "E1")
			So(StringSlice("Elements")[1], ShouldResemble, "E2")

			So(String("Struct.Member"), ShouldResemble, "A2")
			So(len(StringSlice("Struct.Elements")), ShouldEqual, 2)
			So(StringSlice("Struct.Elements")[0], ShouldResemble, "E11")
			So(StringSlice("Struct.Elements")[1], ShouldResemble, "E12")

			So(String("First.Second.Member"), ShouldResemble, "A3")
			So(len(StringSlice("First.Second.Elements")), ShouldEqual, 2)
			So(StringSlice("First.Second.Elements")[0], ShouldResemble, "E111")
			So(StringSlice("First.Second.Elements")[1], ShouldResemble, "E112")
		})
	})
}
