package sendinbluetask

import (
	"context"
)

type hostingOptimizerImpl struct {
	stAdapter storageAdapter
	ctx       context.Context
}

//
// ListTigers List all tigers registered previously on query basis
//
func (this *hostingOptimizerImpl) GetInefficientHosts(ipCount int) ([]ResInactiveIpCount, error) {

	res := make([]ResInactiveIpCount, 0)
	//default value for threshold is 1
	if ipCount == 0 {
		ipCount = 1
	}
	dbRes, err := this.stAdapter.getInactiveIPHosts(this.ctx, ipCount)
	if err != nil {
		return nil, err
	}
	for _, eachData := range dbRes {
		res = append(res, ResInactiveIpCount{
			HostName: eachData,
		})
	}
	return res, nil
}
