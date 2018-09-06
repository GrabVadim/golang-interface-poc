package main

import (
	"fmt"
)

var companies = []*Company{}

func ActionsAndFeatures(af ActionFeature) {
	af.ShowCompanyFeature()
	af.SignEmployee()
}

// ShowFeatures Display stuff from a company
func ShowFeatures(f Feature) {
	f.ShowCompanyFeature()
}

// RunActions run actions
func RunActions(a Action) {
	a.SignEmployee()
}

func CanIInvoque() {
	fmt.Printf("Of course!! \n")
}

func main() {
	companies = append(companies, &Company{"google", "colorado", 100, CanIInvoque})
	companies = append(companies, &Company{"microsoft", "silicon valley", 100, CanIInvoque})

	for _, company := range companies {
		var af ActionFeature
		af = company
		af.ShowCompanyFeature()
		af.SignEmployee()
		af.ShowCompanyFeature()
		// var a Action
		// var f Feature
		// a = company
		// f = company
		// f.ShowCompanyFeature()
		// a.SignEmployee()
		// f.ShowCompanyFeature()
		// i := company.ShowSomething()
		// fmt.Printf("%d \n", i)
		// company.myFunction()
		// ShowFeatures(company)
		// RunActions(company)
		// ShowFeatures(company)
		// ActionsAndFeatures(company)
	}
}
