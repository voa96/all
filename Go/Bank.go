package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

type Speaker interface {
	Speak() string
}

type BankAccount struct {
	Name         string
	Age          int
	balance      float64
	transactions []float64
	cardnumber   int
	password     string
}

var accounts = make(map[int]BankAccount)

func (b BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сумма должна быть положительной")
	}
	b.balance += amount
	b.transactions = append(b.transactions, amount)
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if b.balance-amount < 0 {
		return fmt.Errorf("недостаточно средстр баланс: %v нехватает %v", b.balance, amount-b.balance)
	}
	b.balance -= amount
	b.transactions = append(b.transactions, -amount)
	return nil
}

func (b BankAccount) Speak() string {
	return fmt.Sprintf("Я %s и я пользователь банка ****", b.Name)
}

func Speck(s Speaker) {
	fmt.Println(s.Speak())
}

func (b *BankAccount) GetTransactions() []float64 {
	return b.transactions
}

func accountMenu(b *BankAccount, cardnumber int) {
	for {
		fmt.Println("\n--- МЕНЮ АККАУНТА ---")
		fmt.Println("1. Проверить баланс")
		fmt.Println("2. Пополнить счет")
		fmt.Println("3. Снять деньги")
		fmt.Println("4. История транзакций")
		fmt.Println("5. Выйти")
		fmt.Print("Выберите действие: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Баланс : %.2f\n", b.GetBalance())
		case 2:
			var amount float64
			fmt.Print("На сколько хотите пополнить счёт : ")
			fmt.Scan(&amount)
			if err := b.Deposit(amount); err != nil {
				fmt.Println("ошибка", err)
			} else {
				fmt.Println("счет пополнен")
				accounts[cardnumber] = *b
			}
		case 3:
			var amount float64
			fmt.Print("сколько хотите снять : ")
			fmt.Scan(&amount)
			if err := b.Withdraw(amount); err != nil {
				fmt.Println("Ошибка", err)
			} else {
				fmt.Println("Баланс -", amount)
				accounts[cardnumber] = *b
			}
		case 4:
			fmt.Println("История транзакций : \n", b.GetTransactions())
		case 5:
			fmt.Print("выполняем выходы")
			return
		default:
			fmt.Print("такого варианта нет")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Print("Выберите действие: \n")
		fmt.Print("1. Регистрация \n")
		fmt.Print("2. Войти \n")
		fmt.Print("--- Ваш выбор: ")
		var choice2 int
		fmt.Scan(&choice2)
		switch choice2 { 
		case 1:
			registerAccount()
		case 2:
			loginAccount()
		default:
			fmt.Println("такого варианта нет")
		}
	}

}

func registerAccount() {
	var name string
	var age int
	var password string
	cardnumber := rand.Intn(9000) + 1000
	for {
		if _, exists := accounts[cardnumber]; exists {
			cardnumber = rand.Intn(9000) + 1000
		} else {
			break
		}
	}

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	fmt.Print("Придумайте пароль не менее 4 символов: ")
	for {
		fmt.Scan(&password)
		if utf8.RuneCountInString(password) >= 4 {
			break
		}
		fmt.Print("пароль должен содержать не менее 4 символов ")
	}
	fmt.Printf("Ваш номер карты: %d \n", cardnumber)

	account := BankAccount{
		Name:       name,
		Age:        age,
		cardnumber: cardnumber,
		balance:    0,
		password:   password,
	}
	accounts[cardnumber] = account
	fmt.Printf("Счет создан, номер счета: %d\n", cardnumber)
}

func loginAccount() {
	var cartNumber int
	var account BankAccount
	var exists bool

	for {
		fmt.Print("Введите номер счета: ")
		fmt.Scan(&cartNumber)

		account, exists = accounts[cartNumber]
		if exists {
			break
		}
		fmt.Println("Счет не найден! Попробуйте снова.\n")
	}
	for {
		var password string
		fmt.Print("Введите пароль: ")
		fmt.Scan(&password)

		if account.password == password {
			break
		}
		fmt.Println("Неверный пароль! Попробуйте снова.\n")
	}
	accountMenu(&account, cartNumber)
}
