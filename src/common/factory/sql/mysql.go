package sql

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/factory/sql"
	concurrenthashmap "github.com/kukkar/common-golang/pkg/utils/concurrenthashmap"
	appConf "github.com/kukkar/mta-hosting-optimizer/conf"
)

// DefaultKey default pool key for mysql conn
const DefaultKey = "default"

var mysqlMap = concurrenthashmap.New()

type MysqlPool struct {
	sql.MySqlPool
}

func GetDbPool(key string) (*MysqlPool, error) {
	if val, ok := mysqlMap.Get(key); !ok {
		//we dont have a pool by this key, initiate new pool.
		pool, err := InitNewDBPool(key)
		if err != nil {
			return nil, fmt.Errorf("Could not initiate pool for key:%s, Error:%s",
				key, err.Error())
		}
		mysqlMap.Put(key, pool)
		return pool, nil
	} else {
		return val.(*MysqlPool), nil
	}
}

// InitPool Generally not to be called explicitely, but if you are in desperate need
// do not hesitate to call.
func InitNewDBPool(key string) (*MysqlPool, error) {

	conf, err := appConf.GetAppConfig()
	if err != nil {
		return nil, err
	}
	if conf.MySql == nil {
		return nil, fmt.Errorf("mysql config can not be empty")
	}
	pool, err := sql.Initiate((*conf.MySql))
	if err != nil {
		return nil, err
	}
	return &MysqlPool{*pool}, nil
}
