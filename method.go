package ddp

// rpc structures

type Method struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`

	RandomSeed string `json:"randomSeed"`
}

type Result struct {
	Error  Error  `json:"error"`
	Result string `json:"result"`
}
