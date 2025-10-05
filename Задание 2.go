package main

import (
    "errors"
    "fmt"
)

type Product struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

type Inventory struct {
    products map[int]Product
}

func (i *Inventory) AddProduct(product Product) {
    if i.products == nil {
        i.products = make(map[int]Product)
    }
    i.products[product.ID] = product
}

func (i *Inventory) WriteOff(productID int, quantity int) error {
    product, ok := i.products[productID]
    if !ok {
        return errors.New("товар не найден")
    }
    if product.Quantity < quantity {
        return errors.New("недостаточно товара на складе")
    }
    product.Quantity -= quantity
    i.products[productID] = product 
    return nil
}

func (i *Inventory) RemoveProduct(productID int) error {
    _, ok := i.products[productID]
    if !ok {
        return errors.New("товар не найден")
    }
    delete(i.products, productID)
    return nil
}

func (i *Inventory) GetTotalValue() float64 {
    total := 0.0
    for _, product := range i.products {
        total += product.Price * float64(product.Quantity)
    }
    return total
}

func main() {
    inventory := Inventory{products: make(map[int]Product)}
    inventory.AddProduct(Product{ID: 1, Name: "Яблоко", Price: 1.0, Quantity: 10})
    inventory.AddProduct(Product{ID: 2, Name: "Банан", Price: 0.5, Quantity: 20})
    fmt.Println("Общая стоимость склада:", inventory.GetTotalValue()) 
    err := inventory.WriteOff(1, 5)
    if err == nil {
        fmt.Println("Общая стоимость склада после списания:", inventory.GetTotalValue()) 
    }
}
