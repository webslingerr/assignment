package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"os"

	"io/ioutil"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewBranchRepo(fileName string) *branchRepo {
	return &branchRepo{
		fileName: fileName,
	}
}

func (b *branchRepo) Create(req *models.CreateBranch) (string, error) {
	branches, err := b.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	branches = append(branches, models.Branch{
		Id:   uuid,
		Name: req.Name,
	})

	err = b.WriteFileToJson(branches)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (b *branchRepo) GetById(req *models.BranchPrimaryKey) (models.Branch, error) {
	branches, err := b.Read()
	if err != nil {
		return models.Branch{}, err
	}

	for _, v := range branches {
		if req.Id == v.Id {
			return v, nil
		}
	}

	return models.Branch{}, errors.New("There is no branch with this id")
}

func (b *branchRepo) GetAll(req *models.GetAllBranchRequest) (models.GetAllBranchResponse, error) {
	branches, err := b.Read()
	if err != nil {
		return models.GetAllBranchResponse{}, err
	}

	if req.Limit + req.Offset > len(branches) {
		if req.Offset > len(branches) {
			return models.GetAllBranchResponse{}, errors.New("Invalid offset or limit")
		}
		br := branches[req.Offset:]
		return models.GetAllBranchResponse{
			Count: len(br),
			Branches: br,
		}, nil
	}

	br := branches[req.Offset : req.Offset + req.Limit]
	return models.GetAllBranchResponse{
		Count: len(br),
		Branches: br,
	}, nil
}

func (b *branchRepo) Update(req *models.UpdateBranch) error {
	branches, err := b.Read()
	if err != nil {
		return err
	}

	flag := false
	for i, v := range branches {
		if v.Id == req.Id {
			branches[i].Name = req.Name
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("There is no branch with this id")
	}

	err = b.WriteFileToJson(branches)
	if err != nil {
		return err
	}
	return nil
}

func (b *branchRepo) Delete(req *models.BranchPrimaryKey) error {
	branches, err := b.Read()
	if err != nil {
		return err
	}

	flag := false
	for i, v := range branches {
		if v.Id == req.Id {
			branches = append(branches[:i], branches[i+1:]...)
			flag = true
		}
	}

	if !flag {
		return errors.New("There is no branch with this id")
	}

	err = b.WriteFileToJson(branches)
	if err != nil {
		return err
	}

	return nil
}

func (b *branchRepo) Read() ([]models.Branch, error) {
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return []models.Branch{}, err
	}

	var branches []models.Branch
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return []models.Branch{}, err
	}
	return branches, nil
}

func (b *branchRepo) WriteFileToJson(items interface{}) error {
	body, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
