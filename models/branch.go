package models

type Branch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateBranch struct {
	Name string `json:"name"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type GetAllBranchRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetAllBranchResponse struct {
	Count    int      `json:"count"`
	Branches []Branch `json:"branches"`
}

type UpdateBranch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
