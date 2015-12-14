package ddp

type Connect struct {
	Session string   `json:"session"`
	Version string   `json:"version"`
	Support []string `json:"support"`
}

type Connected struct {
	Session string `json:"session"`
}

type ConnectFailed struct {
	Version string `json:"version"`
}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"msg"`
}
