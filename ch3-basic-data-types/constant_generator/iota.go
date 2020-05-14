package main

import "fmt"

/*
	A constant declaration may use the constant generator
	iota, which is used to create a sequence of related values
	without having to spell out each one explicitly.
	A typical example is Enums
*/

type Weekday int

const (
	Sunday Weekday = iota //Start out from 0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
	/*
		As iota increments, each constant is assigned the value of 1 << iota
	*/
)

func main() {

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	/* Output
	0
	1
	2
	3
	4
	5
	6
	*/

	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)
	fmt.Println(FlagPointToPoint)
	fmt.Println(FlagMulticast)

	/* Output
	1
	2
	4
	8
	16
	*/

}
