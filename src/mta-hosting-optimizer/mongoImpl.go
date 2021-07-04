package sendinbluetask

import (
	"context"
	"fmt"

	mfactory "github.com/kukkar/mta-hosting-optimizer/src/common/factory/mongof"
)

type mongoAdapter struct {
	adatper *mfactory.MDB
}

func (this *mongoAdapter) getInactiveIPHosts(ctx context.Context,
	activeIPsCount int) ([]string, error) {

	// rawConn := this.adatper.GetRawConn()

	// coll := rawConn.Collection(IpHostnameCollection)

	// groupStage := bson.D{{"$group", bson.D{{"_id", "$hostname"}, {"total", bson.D{{"$sum", "$hostname"}}}}}}
	// matchStage := bson.D{{"$match", bson.D{{"total", bson.D{{"$lte", activeIPsCount}}}}}}

	// showInfoCursor, err := coll.Aggregate(ctx, []bson.D{matchStage, groupStage})
	// if err != nil {
	// 	return nil, fmt.Errorf(err.Error())
	// }
	// listData := make([]DBResInactiveIpCount, 0)
	// dbData := make([]MongoIpStatusCollection, 0)
	// if err = showInfoCursor.All(ctx, &dbData); err != nil {
	// 	return nil, err
	// }
	// for _, eachData := range dbData {
	// 	listData = append(listData, DBResInactiveIpCount{
	// 		IP:       eachData.IP,
	// 		HostName: eachData.HostName,
	// 	})
	// }
	return nil, fmt.Errorf("todo")
}
