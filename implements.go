package main

import (
	"fmt"
)

// Actions interface

// SignEmployee to sign it someone
func (c *Company) SignEmployee() {
	c.totalEmployees++
}

// FireEmployee to fire someone
func (c *Company) FireEmployee() {
	c.totalEmployees--
}

// Feature interface

// ShowName caca
func (c Company) ShowName() {
	fmt.Printf("Name: %v \n", c.name)
}

// ShowAddress to show an address lol
func (c Company) ShowAddress() {
	fmt.Printf("Address: %v \n", c.address)
}

// ShowTotalEmployee show total employees
func (c Company) ShowTotalEmployee() {
	fmt.Printf("Total Employees: %v \n", c.totalEmployees)
}

func (c Company) ShowCompanyFeature() {
	fmt.Printf("Name: %v - Address: %v - Employees: %v \n", c.name, c.address, c.totalEmployees)
}

func (c Company) ShowSomething() int {
	fmt.Println("can you invoque me?")
	return 1
}
