package ddp

// connect structures

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
