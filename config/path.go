package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/erickxeno/clib/errors"
	"github.com/erickxeno/clib/utils"
)

const (
	FolderCnf  = "conf"
	SuffixYml  = ".yml"
	CiEnvName  = "SUB_ENV"
	CiEnvValue = "ci"
)

func IsCIEnv() bool {
	return /*env.IsBoe() &&*/ strings.ToLower(os.Getenv("SUB_ENV")) == CiEnvValue
}

func GetConfigYmlFile(confDir string) (string, error) {
	//paths := caesar.GetOrderedConfKeys(caesar.NewOption())
	paths := []string{"base.yml", "boe.yml", "boe.unknown_env.yml", "boe.unknown_env.default.yml"}
	newPaths := []string{}
	var tenantSuffix string
	localTenant := "local" // TODO: get run env tenant
	switch localTenant {
	case "local":
		tenantSuffix = "_local"
	case "other":
		tenantSuffix = "_other"
	default:
		panic("tenant not support currently")
	}
	for _, path := range paths {
		newPaths = append(newPaths, strings.TrimSuffix(path, SuffixYml)+tenantSuffix+SuffixYml)
	}
	if IsCIEnv() {
		newPaths = append(newPaths, "ci"+SuffixYml)
	}
	configFile := ""
	for _, path := range newPaths {
		path = filepath.Join(confDir, path)
		if !utils.FileExist(path) {
			continue
		}
		configFile = path
	}
	if configFile == "" {
		return "", errors.Errorf("no config file found, dir:%s, IsCIEnv:%v", confDir, IsCIEnv())
	}
	fmt.Printf("config_file loaded: %s, IDC: %s, Region: %s, HostEnv: %s, TceEnv: %s", configFile, "env.IDC()", "env.Region()", os.Getenv("TCE_HOST_ENV"), os.Getenv("TCE_ENV"))
	return configFile, nil
}

func FindConfDir(confDir string, level ...int) string {
	depth := 0
	if len(level) > 0 {
		depth = level[0]
	}
	maxSearchDepth := 10
	if depth > maxSearchDepth {
		return ""
	}
	if utils.IsDir(confDir) {
		return confDir
	}
	return FindConfDir("../"+confDir, depth+1)
}
