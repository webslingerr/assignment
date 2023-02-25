package models

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type CategoryStatistics struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Category struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}
