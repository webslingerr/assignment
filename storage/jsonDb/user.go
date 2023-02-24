package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"os"

	"io/ioutil"
)

type userRepo struct {
	fileName string
}

func NewUserRepo(fileName string) *userRepo {
	return &userRepo{
		fileName: fileName,
	}
}

func (u *userRepo) GetById(req *models.UserPrimaryKey) (models.User, error) {
	users, err := u.Read()
	if err != nil {
		return models.User{}, err
	}

	for _, v := range users {
		if req.Id == v.Id {
			return v, nil
		}
	}

	return models.User{}, errors.New("There is no user with this id")
}

func (u *userRepo) Update(req *models.UpdateUserBalance) error {
	users, err := u.Read()
	if err != nil {
		return err
	}

	flag := false
	for i, v := range users {
		if req.Id == v.Id {
			users[i].Balance = req.Balance
			flag = true
		}
	}

	if !flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	
	return nil
}

func (b *userRepo) Read() ([]models.User, error) {
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return []models.User{}, err
	}

	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}
