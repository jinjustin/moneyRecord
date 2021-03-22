package moneySaver

import (
	"time"

	"google.golang.org/grpc"
    "google.golang.org/grpc/codes"
)

type Income struct{
	Get int `json:"get"`
	From string `json:"from"`
	Type string `json:"type"`
	Date string `json:"date"`
}

//NewIncome is a function that use to create new account to store money.
func (m *moneySaver)NewIncome(accountName string, get int, from string, incomeType string) error{

	format := "2006-01-02"

	i := Income{
		Get: get,
		From: from,
		Type: incomeType,
		Date: time.Now().Format(format),
	}

	_, _, err := m.client.Collection("income").Add(ctx, map[string]interface{}{
        "get":    i.Get,
        "from": i.From,
		"type": i.Type,
		"Date": i.Date,
	})
	if err != nil {
		return err
	}

	err = m.addIncomeType(i.Type)
	if err != nil{
		return err
	}

	err = m.Deposit(accountName, i.Get)
	if err != nil {
		return err
	}
	return nil
}

func (m *moneySaver)addIncomeType(incomeType string) error{

	type AllIncomeType struct{
		Data []string
	}
	var all AllIncomeType

	dsnap, err := m.client.Collection("incomeType").Doc("allType").Get(ctx)
	if grpc.Code(err) == codes.NotFound{
		data := []string{incomeType}
		_, err = m.client.Collection("incomeType").Doc("allType").Set(ctx, map[string]interface{}{
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
		if d == incomeType{
			check = false
		}
	}
	if check{
		all.Data = append(all.Data, incomeType)
	}

	_, err = m.client.Collection("incomeType").Doc("allType").Set(ctx, map[string]interface{}{
        "data": all.Data,
	})
	if err != nil {
		return err
	}
	return nil
}