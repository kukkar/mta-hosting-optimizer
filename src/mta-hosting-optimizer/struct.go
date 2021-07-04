package sendinbluetask

import "github.com/kukkar/common-golang/pkg/utils"

type ResInactiveIpCount struct {
	HostName string
	IP       string
}

type DBResInactiveIpCount struct {
	HostName string
	IP       string
}
type Config struct {
	StorageAdapter string
	RC             utils.RequestContext
}
