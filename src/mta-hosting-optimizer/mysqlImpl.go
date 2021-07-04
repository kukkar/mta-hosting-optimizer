package hostingoptimizer

import (
	"context"

	"github.com/kukkar/mta-hosting-optimizer/src/common/factory/sql"
)

type sqlAdapter struct {
	adatper *sql.MysqlPool
}

func (this *sqlAdapter) getInactiveIPHosts(ctx context.Context,
	activeIPsCount int) ([]string, error) {

	dbData := make([]string, 0)
	conn := this.adatper.MySqlPool.GetConnection()

	query := conn.Conn.Table(IpHostNameTable)
	rows, err := query.Raw("select hostname from ipconfig where active='1' group by hostname having count(active) <= ?", activeIPsCount).Rows()
	defer rows.Close()
	if err != nil {
		return nil, query.Error
	}
	for rows.Next() {
		var d string
		rows.Scan(&d)
		dbData = append(dbData, d)
	}
	return dbData, nil
}
