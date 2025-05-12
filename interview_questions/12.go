package main

import "fmt"

// “Design an interface called Notifier with a Notify(message string) method. Implement it for both EmailSender and SMSSender, then write a function that takes a Notifier and calls Notify.”

type Notifier interface {
	Notify(message string)
}

type EmailSender struct {
	a string
}

type SMSSender struct {
	b string
}

func (e *EmailSender) Notify(message string) {
	fmt.Printf("%v received from %v\n", message, e.a)
}

func (s *SMSSender) Notify(message string) {
	fmt.Printf("%v received from %v\n", message, s.b)
}

// func main() {

// 	i := []Notifier{
// 		&EmailSender{
// 			a: "Email",
// 		},
// 		&SMSSender{
// 			b: "SMS",
// 		},
// 	}

// 	for _, r := range i {
// 		r.Notify("message")
// 	}

// }
