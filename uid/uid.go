package uid

import (
	"context"
	"fmt"
	"sync"

	"github.com/erickxeno/clib/config"
	"github.com/erickxeno/clib/utils"
)

const (
	DefaultNamespace = "default_namespace"
)

type Generator interface {
	NewID(ctx context.Context) (int64, error)
	NewIDs(ctx context.Context, count int) []int64
}

type IDGenerator struct {
	namespace string
	//client
}

var (
	generatorMap     sync.Map
	lock             sync.Mutex
	defaultGenerator Generator
)

func init() {
	defaultGenerator = GetGenerator(DefaultNamespace)
	if utils.IsNil(defaultGenerator) {
		panic(fmt.Sprintf("create generator failed, namesapce:%v", DefaultNamespace))
	}
}

func GetDefaultGenerator() Generator {
	return defaultGenerator
}

func GetGenerator(namespace string) Generator {
	key := namespace
	generator, ok := generatorMap.Load(key)
	if ok {
		ig, _ := generator.(Generator)
		return ig
	}
	lock.Lock()
	defer lock.Unlock()
	generator, ok = generatorMap.Load(key)
	if ok {
		ig, _ := generator.(Generator)
		return ig
	}
	var ig Generator
	if !config.IsCIEnv() {
		//client, err := idgenerator.NewNtIdGeneratorBuilder().
		//	WithNamespace(namespace).
		//	Build()
		//if err != nil {
		//	lg.Error(context.Background(), "NewNtIdGeneratorBuilder fail", lg.Err(err), lg.Data("namespace", namespace))
		//	return nil
		//}
		ig = &IDGenerator{
			namespace: namespace,
			// client: client,
		}
	} else {
		ig = &MockIDGenerator{}
	}
	generatorMap.Store(key, ig)
	return ig
}

func (i *IDGenerator) NewID(ctx context.Context) (int64, error) {
	//ID, err := i.client.Get(ctx)
	return int64(0), nil
}
func (i *IDGenerator) NewIDs(ctx context.Context, count int) []int64 {
	//IDs, err := i.client.MGet(ctx, count)
	return []int64{}
}
