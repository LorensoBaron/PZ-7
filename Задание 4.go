package main

import "fmt"

type Order struct {
    ID         int
    CustomerID int
    Items      []OrderItem
    Status     string 
}

type OrderItem struct {
    ProductID int
    Quantity  int
    Price     float64 
}

type Customer struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
}

func (o *Order) CalculateTotal() float64 {
    total := 0.0
    for _, item := range o.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

func (o *Order) AddItem(item OrderItem) {
    o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(productID int) {
    for i, item := range o.Items {
        if item.ProductID == productID {
            o.Items = append(o.Items[:i], o.Items[i+1:]...) 
            return                                         
        }
    }
}

func (o *Order) SetStatus(status string) {
    o.Status = status
}

func main() {
    order := Order{
        ID:         1,
        CustomerID: 101,
        Items: []OrderItem{
            {ProductID: 1, Quantity: 2, Price: 10.0},
            {ProductID: 2, Quantity: 1, Price: 25.0},
        },
        Status: "Создан",
    }

    fmt.Println("Общая сумма заказа:", order.CalculateTotal()) 
    order.AddItem(OrderItem{ProductID: 3, Quantity: 3, Price: 5.0})
    fmt.Println("Общая сумма заказа после добавления:", order.CalculateTotal())
    order.RemoveItem(1)
    fmt.Println("Общая сумма заказа после удаления:", order.CalculateTotal()) 
    order.SetStatus("Подтвержден")
    fmt.Println("Статус заказа", order.Status) 
}
