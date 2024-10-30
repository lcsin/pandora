package domain

// Music 数据库音乐表实体映射
type Music struct {
	ID     int64   `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Author string  `json:"author,omitempty"`
	URL    string  `json:"url,omitempty"`
}
