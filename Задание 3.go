package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

type User struct {
    Username string
    Email    string
    Password string 
}

func (u *User) SetPassword(password string) {
    hash := sha256.Sum256([]byte(password))
    u.Password = hex.EncodeToString(hash[:])
}

func (u *User) VerifyPassword(password string) bool {
    hash := sha256.Sum256([]byte(password))
    hashedPassword := hex.EncodeToString(hash[:])
    return u.Password == hashedPassword
}

func main() {
    user := User{Username: "testuser", Email: "test@example.com"}
    user.SetPassword("password123")
    fmt.Println("Проверка пароля... Пароль:", user.VerifyPassword("password123")) 
    fmt.Println("Проверка пароля... Пароль:", user.VerifyPassword("wrongpassword")) 
}
