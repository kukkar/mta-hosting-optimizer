package hostingoptimizer

import (
	"context"
	"fmt"
	"testing"

	"github.com/kukkar/common-golang/pkg/factory/sql"
	csql "github.com/kukkar/mta-hosting-optimizer/src/common/factory/sql"
)

func TestGetUnusedHosts(t *testing.T) {

	//c := context.TODO()
	var st testStorageAdapterImpl
	instace := hostingOptimizerImpl{
		stAdapter: &st,
	}

	res, err := instace.GetInefficientHosts(1)
	if err != nil {
		t.Errorf("Enable to signup expected success got %s", err.Error())
		return
	}

	t.Logf("output is %v", res)
}

func TestGetUnusedHostsWithEmptyValue(t *testing.T) {

	//c := context.TODO()
	var st testStorageAdapterImpl
	instace := hostingOptimizerImpl{
		stAdapter: &st,
	}

	res, err := instace.GetInefficientHosts(0)
	if err != nil {
		t.Errorf("Enable to signup expected success got %s", err.Error())
		return
	}
	t.Logf("output is %v", res)
}

func TestGetUnusedHostsWithError(t *testing.T) {

	//c := context.TODO()
	var st testStorageAdapterImplForError
	instace := hostingOptimizerImpl{
		stAdapter: &st,
	}

	res, err := instace.GetInefficientHosts(1)
	if err != nil {
		t.Logf("Enable to get data from sql got %s", err.Error())
		return
	}
	t.Logf("output is %v", res)
}

func TestGetInactiveIPHosts(t *testing.T) {

	c := context.TODO()
	pool, err := sql.Initiate((sql.MysqlConfig{
		User:            "root",
		Password:        "Golang@1994",
		DBName:          "hostingoptimizer",
		Host:            "127.0.0.1",
		Port:            "3306",
		DefaultTimeZone: "Europe/Paris",
	}))
	if err != nil {
		t.Errorf("Enable to signup expected success got %s", err.Error())
	}
	adapter := sqlAdapter{
		adatper: &csql.MysqlPool{*pool},
	}
	res, err := adapter.getInactiveIPHosts(c, 1)
	if err != nil {
		t.Errorf("Enable to signup expected success got %s", err.Error())
	}
	t.Logf("output is %v", res)
}

type testStorageAdapterImpl struct {
}

type testStorageAdapterImplForError struct {
}

func (this *testStorageAdapterImplForError) getInactiveIPHosts(ctx context.Context,
	activeIPsCount int) ([]string, error) {
	return nil, fmt.Errorf("unable to get data from sql")
}

func (this *testStorageAdapterImpl) getInactiveIPHosts(ctx context.Context,
	activeIPsCount int) ([]string, error) {
	return []string{"mta-prod-1", "mta-prod-2"}, nil
}
