package models

type Info struct {
	Patient_Id       int
	Name             string
	Address          string
	PhoneNo          int
	Disease          string
	Modifier         string
	Payer            string
	Bill_no          int
	Bill_Amount      int
	Paid_by_Payer    int
	Remaining_Amount int
}
