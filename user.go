package main

import "fmt"

type User struct {
	Name        string
	Email       string
	Phone       string
	Id          int
	savings     int
	GroupJoined map[int]*Group
	balance     map[int]float64
}

func createUser(name string, email string, phone string, id int, savings int) *User {
	return &User{
		Name:        name,
		Email:       email,
		Phone:       phone,
		Id:          id,
		savings:     savings,
		GroupJoined: make(map[int]*Group),
		balance:     make(map[int]float64), //group id - amount
	}
}

func (u *User) joinGroup(g *Group) {
	u.GroupJoined[g.Id] = g
	u.balance[g.Id] = 0.0
}

func (u *User) addMoneytoSavings(amount int) {
	u.savings += amount
	fmt.Println("successfully added money to savings, current savings:", u.savings)
}

func (u *User) payExpense(g *Group, amount float64) {
	group, ok := u.GroupJoined[g.Id]
	if !ok || group == nil {
		fmt.Println("User not part of the group")
		return
	}
	if u.balance[g.Id] == 0 {
		fmt.Println("all the expenses have been paid off")
		return
	}
	u.balance[g.Id] -= amount
	fmt.Println("remaining amount to be paid")
}
func (u *User) createGroup(
	name string,
	members []*User,
	total float64,
	splittype Splittype,
) *Group {

	id := len(u.GroupJoined) + 1

	splitMap := make(map[int]float64)

	// initialize split amounts to 0
	for _, member := range members {
		splitMap[member.Id] = 0.0
	}

	newgroup := &Group{
		Id:          id,
		Name:        name,
		Members:     members,
		total:       total,
		splitamount: splitMap,
		splittype:   splittype,
	}

	u.GroupJoined[id] = newgroup
	u.balance[id] = 0

	return newgroup
}
