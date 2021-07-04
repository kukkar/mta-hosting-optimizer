package hostingoptimizer

type SendingBlueTask interface {
	GetInefficientHosts(ipCount int) ([]ResInactiveIpCount, error)
}
