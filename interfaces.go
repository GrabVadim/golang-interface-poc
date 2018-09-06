package main

type ActionFeature interface {
	Action
	Feature
}

// Action 's Compnay
type Action interface {
	SignEmployee()
	FireEmployee()
}

// Feature of a company
type Feature interface {
	ShowName()
	ShowAddress()
	ShowTotalEmployee()
	ShowCompanyFeature()
}
