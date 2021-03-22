package moneySaver

import (
	"time"

	"google.golang.org/grpc"
    "google.golang.org/grpc/codes"
)

type Expense struct{
	Pay int `json:"get"`
	To string `json:"from"`
	Type string `json:"type"`
	Date string `json:"date"`
}

//NewIncome is a function that use to create new account to store money.
func (m *moneySaver)NewExpense(accountName string, pay int, to string, expenseType string) error{

	format := "2006-01-02"

	e := Expense{
		Pay: pay,
		To: to,
		Type: expenseType,
		Date: time.Now().Format(format),
	}

	_, _, err := m.client.Collection("expense").Add(ctx, map[string]interface{}{
        "get":  e.Pay,
        "from": e.To,
		"type": e.Type,
		"Date": e.Date,
	})
	if err != nil {
		return err
	}

	err = m.addExpenseType(e.Type)
	if err != nil{
		return err
	}

	err = m.Withdraw(accountName, e.Pay)
	if err != nil {
		return err
	}
	return nil
}

func (m *moneySaver)addExpenseType(expenseType string) error{

	type AllExpenseType struct{
		Data []string
	}
	var all AllExpenseType

	dsnap, err := m.client.Collection("expenseType").Doc("allType").Get(ctx)
	if grpc.Code(err) == codes.NotFound{
		data := []string{expenseType}
		_, err = m.client.Collection("expenseType").Doc("allType").Set(ctx, map[string]interface{}{
			"data": data,
		})
		if err != nil {
			return err
		}
	}else if err != nil{
		return err
	}
	
	dsnap.DataTo(&all)

	check := true
	for _, d := range all.Data {
		if d == expenseType{
			check = false
		}
	}
	if check{
		all.Data = append(all.Data, expenseType)
	}

	_, err = m.client.Collection("expenseType").Doc("allType").Set(ctx, map[string]interface{}{
        "data": all.Data,
	})
	if err != nil {
		return err
	}
	return nil
}