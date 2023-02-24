package jsonDb

import (
	"app/config"
	"app/storage"
	"io/ioutil"
	"os"
)

type Store struct {
	branch   *branchRepo
	shopCart *shopCartRepo
	user     *userRepo
	product  *productRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {
	if !fileExists(cfg.BranchFileName) {
		os.Create(cfg.BranchFileName)
		err := ioutil.WriteFile(cfg.BranchFileName, []byte("[]"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	if !fileExists(cfg.ShopCartFileName) {
		os.Create(cfg.ShopCartFileName)
		err := ioutil.WriteFile(cfg.ShopCartFileName, []byte("[]"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	if !fileExists(cfg.UserFileName) {
		os.Create(cfg.UserFileName)
		err := ioutil.WriteFile(cfg.UserFileName, []byte("[]"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	if !fileExists(cfg.ProductFileName) {
		os.Create(cfg.ProductFileName)
		err := ioutil.WriteFile(cfg.ProductFileName, []byte("[]"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return &Store{
		branch:   NewBranchRepo(cfg.BranchFileName),
		shopCart: NewShopCart(cfg.ShopCartFileName),
		user:     NewUserRepo(cfg.UserFileName),
		product:  NewProductRepo(cfg.ProductFileName),
	}, nil
}

func (s *Store) CloseDb() {}

func (s *Store) Branch() storage.BranchRepoI {
	return s.branch
}

func (s *Store) ShopCart() storage.ShopCartRepoI {
	return s.shopCart
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Product() storage.ProductRepoI {
	return s.product
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}
