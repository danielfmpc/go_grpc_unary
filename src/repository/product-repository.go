package repository

import (
	"fmt"
	"os"

	"github.com/danielfmpc/go_rpc_unary/src/pb/products"
	"google.golang.org/protobuf/proto"
)

type ProductRepository struct {
}

const filename string = "products.txt"

func (pr *ProductRepository) LoadData() (products.ProductList, error) {
	productList := products.ProductList{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return productList, fmt.Errorf("error on read file: %+v", err)
	}

	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("error on unmarshal: %+v", err)
	}

	return productList, nil
}

func (pr *ProductRepository) SaveData(productList products.ProductList) error {
	data, err := proto.Marshal(&productList)
	if err != nil {
		return fmt.Errorf("error on marshal: %+v", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error on write file: %+v", err)
	}
	return nil
}

func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	productsList, err := pr.LoadData()
	if err != nil {
		return product, fmt.Errorf("error on load: %+v", err)
	}
	product.Id = int32(len(productsList.Products) + 1)
	productsList.Products = append(productsList.Products, &product)
	err = pr.SaveData(productsList)
	if err != nil {
		return product, fmt.Errorf("error on savedata: %+v", err)
	}

	return product, nil
}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	return pr.LoadData()
}
