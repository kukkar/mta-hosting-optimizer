//+build !test

package hostingoptimizer

import (
	"context"
	"fmt"

	mfactory "github.com/kukkar/mta-hosting-optimizer/src/common/factory/mongof"
	"github.com/kukkar/mta-hosting-optimizer/src/common/factory/sql"
)

func GetInterface(c context.Context,
	config Config) (SendingBlueTask, error) {

	stAdapter, err := getStorageAdapter(mfactory.DEFAULT_KEY, config.StorageAdapter)
	if err != nil {
		return nil, err
	}
	return &hostingOptimizerImpl{
		stAdapter: stAdapter,
		ctx:       c,
	}, nil
}

func getStorageAdapter(key string,
	stAdapter string) (storageAdapter, error) {
	switch stAdapter {
	case StInMemory:
	case StMongo:
	case StMysql:
		msql, err := getMysqlAdapter("")
		if err != nil {
			return nil, err
		}
		return &sqlAdapter{
			msql,
		}, nil
	}
	return nil, fmt.Errorf("wrong choice of Adapter")
}

func getMysqlAdapter(key string) (*sql.MysqlPool, error) {
	if key == "" {
		key = sql.DefaultKey
	}
	return sql.GetDbPool(key)
}
