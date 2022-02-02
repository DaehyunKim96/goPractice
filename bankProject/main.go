package main

import (
	"bankProject/accounts"
	"bankProject/mydict"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	errRequestFailed = errors.New("Request Failed")
)

type result struct {
	url    string
	status string
}

func main() {
	//accountTest()
	//dictTest()
	urlTest()
	//chanTest()
}
func urlTest() {
	var results = map[string]string{}
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	//for _, val := range urls {
	//	if err := hitUrl(val); err != nil {
	//		results[val] = "Failed"
	//	} else {
	//		results[val] = "Success"
	//	}
	//}
	//for key, val := range results {
	//	fmt.Println(key, val)
	//}
	for _, url := range urls {
		go hitUrlGo(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for key, val := range results {
		fmt.Println(key, val)
	}
}

// send only
func hitUrlGo(url string, c chan<- result) {
	//fmt.Println("checking", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, status: "Failed"}
	} else {
		c <- result{url: url, status: "Success"}
	}
}
func hitUrl(url string) error {
	fmt.Println("checking", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
func chanTest() {
	c := make(chan string)
	people := [5]string{"daniel", "hyuns", "one", "two", "three"}
	for _, name := range people {
		go isGood(name, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}
func isGood(name string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- name + " is good"

}
func dictTest() {
	dictionary := mydict.Dictionary{}
	dictionary["first"] = "word1"
	if val, err := dictionary.Search("second"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	err := dictionary.Add("hello", "greeting")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)
	errUpdate := dictionary.Update("hello", "greeting2")
	if errUpdate != nil {
		fmt.Println(errUpdate)
	}
	fmt.Println(dictionary)
	errDelete := dictionary.Delete("hello")
	if errDelete != nil {
		fmt.Println(errDelete)
	}
	fmt.Println(dictionary)
}
func accountTest() {
	account := accounts.NewAccount("daniel")
	//not copy, account object
	account.Deposit(50)
	fmt.Println(account.Balance())
	err := account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	fmt.Println(account.Owner())
	account.ChangeOwner("daniel.hyuns")
	fmt.Println(account.Owner())
	fmt.Println(account)
}
