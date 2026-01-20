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
	fmt.Println("remaining amount to be paid %d", u.balance[g.Id])
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

func (u *User) splitbill(id int, splitAmounts map[int]float64) bool {
	group, ok := u.GroupJoined[id]
	if !ok || group == nil {
		fmt.Println("User not part of the group")
		return false
	}
	if group.total == 0 {
		fmt.Println("No expense to split")
		return false
	}

	switch group.splittype {
	case price:
		perPersonAmount := group.total / float64(len(group.Members))
		for _, member := range group.Members {
			group.splitamount[member.Id] = perPersonAmount
			member.balance[id] += perPersonAmount
		}
	case percentage:
		if len(splitAmounts) == 0 {
			fmt.Println("Percent split requires input percentages")
			return false
		}
		totalPct := 0.0
		for _, member := range group.Members {
			pct, ok := splitAmounts[member.Id]
			if !ok {
				fmt.Println("Missing percentage for member", member.Id)
				return false
			}
			group.splitamount[member.Id] = pct
			totalPct += pct
		}
		if totalPct != 100.0 {
			fmt.Println("Percentages do not sum to 100")
			return false
		}
		for _, member := range group.Members {
			amount := (group.splitamount[member.Id] / 100.0) * group.total
			group.splitamount[member.Id] = amount
			member.balance[id] += amount
		}
	}

	fmt.Println("Bill split successfully")
	return true
}
