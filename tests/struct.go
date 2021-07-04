package tests

type ServiceErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DebugMsg string `json:"debug_msg"`
}

type ResUnusedHosts struct {
	Data   []UnusedHost `json:"data"`
	Status bool         `json:"status"`
}

type UnusedHost struct {
	Host string `json:"host"`
}
