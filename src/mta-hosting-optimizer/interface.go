package sendinbluetask

type SendingBlueTask interface {
	GetInefficientHosts(ipCount int) ([]ResInactiveIpCount, error)
}
