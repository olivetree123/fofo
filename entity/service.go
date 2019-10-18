package entity

type Service struct {
	ID      string                 `json:"id"`
	Name    string                 `json:"name"`
	Group   string                 `json:"group"`
	Port    int                    `json:"port"`
	Address string                 `json:"address"`
	Extra   map[string]interface{} `json:"extra"`
}
