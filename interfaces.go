package main

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
