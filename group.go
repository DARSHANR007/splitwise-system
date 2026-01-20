package main

type Group struct {
	Name        string
	Members     []*User
	Id          int
	total       float64
	splitamount map[int]float64
	splittype   Splittype
}

type Splittype int

const (
	price Splittype = iota
	percentage
)
