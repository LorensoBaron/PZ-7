package main

import (
    "errors"
    "fmt"
)


type BankAccount struct {
    accountNumber string
    holderName    string
    balance       float64
}

func (acc *BankAccount) Deposit(amount float64) {
    acc.balance += amount
}

func (acc *BankAccount) Withdraw(amount float64) error {
    if amount > acc.balance {
        return errors.New("недостаточно средств")
    }
    acc.balance -= amount
    return nil
}

func (acc *BankAccount) GetBalance() float64 {
    return acc.balance
}

func main() {
    account := BankAccount{
        accountNumber: "1234567890",
        holderName:    "Иван Иванов",
        balance:       1000.0,
    }

    fmt.Printf("Баланс до операций: %.2f\n", account.GetBalance())

    account.Deposit(500.0)
    fmt.Printf("Баланс после пополнения: %.2f\n", account.GetBalance())

    err := account.Withdraw(200.0)
    if err != nil {
        fmt.Println("Ошибка при снятии:", err)
    } else {
        fmt.Printf("Баланс после снятия: %.2f\n", account.GetBalance())
    }

    err = account.Withdraw(2000.0)
    if err != nil {
        fmt.Println("Ошибка при снятии:", err)
    }
}
