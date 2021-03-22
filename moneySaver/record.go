package moneySaver

import (
	"time"

	"google.golang.org/grpc"
    "google.golang.org/grpc/codes"
	"google.golang.org/api/iterator"
	"errors"
)

type Record struct{
	RecordType string `json:"recordType"`
	Amount int `json:"amount"`
	From int `json:"from"`
	To int `json:"to"`
	Name string `json:"name"`
	ChangeType string `json:"changeType"`
	Day int `json:"day"`
	Month int `json:"month"`
	Year int `json:"year"`
	Order int `json:"order"`
}

//NewIncome is a function that use to create new account to store money.
func (m *moneySaver)NewIncome(accountName string, amount int, name string, changeType string) error{

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	dsnap, err := m.client.Collection("account").Doc(accountName).Get(ctx)
	if err != nil {
		return err
	}
	var a Account
	dsnap.DataTo(&a)

	r := Record{
		RecordType: "Income",
		From: a.Total,
		To: a.Total+amount,
		Amount: amount,
		Name: name,
		ChangeType: changeType,
		Day : day,
		Month: month,
		Year: year,
		Order: a.Order+1,
	}

	_, _, err = m.client.Collection("record").Add(ctx, map[string]interface{}{
        "recordType": r.RecordType,
		"from": r.From,
		"to": r.To,
        "amount": r.Amount,
		"name": r.Name,
		"changeType": r.ChangeType,
		"day": r.Day,
		"month": r.Month,
		"year": r.Year,
		"order": r.Order,
	})

	if err != nil {
		return err
	}

	err = m.addIncomeType(r.ChangeType)
	if err != nil{
		return err
	}

	err = m.updateAccount(accountName, r.To, r.Order)
	if err != nil {
		return err
	}
	return nil
}

//NewExpense is a function that use to create new account to store money.
func (m *moneySaver)NewExpense(accountName string, amount int, name string, changeType string) error{

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	dsnap, err := m.client.Collection("account").Doc(accountName).Get(ctx)
	if err != nil {
		return err
	}
	var a Account
	dsnap.DataTo(&a)

	if a.Total - amount < 0{
		return errors.New("moneySaver: Expense value greater than money in account.")
	}

	r := Record{
		RecordType: "Expense",
		From: a.Total,
		To: a.Total - amount,
		Amount: amount,
		Name: name,
		ChangeType: changeType,
		Day : day,
		Month: month,
		Year: year,
		Order: a.Order+1,
	}

	_, _, err = m.client.Collection("record").Add(ctx, map[string]interface{}{
        "recordType": r.RecordType,
		"from": r.From,
		"to": r.To,
        "amount": r.Amount,
		"name": r.Name,
		"changeType": r.ChangeType,
		"day": r.Day,
		"month": r.Month,
		"year": r.Year,
		"order": r.Order,
	})

	if err != nil {
		return err
	}

	err = m.addIncomeType(r.ChangeType)
	if err != nil{
		return err
	}

	err = m.updateAccount(accountName, r.To, r.Order)
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

//GetRecordMonthly is a function that use to get all record in input time.
func (m *moneySaver)GetRecordMonthly(month int, year int) ([]Record, error){

	var r Record

	var records []Record

	iter := m.client.Collection("record").Where("month", "==", month).Where("year", "==", year).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		doc.DataTo(&r)
		records = append(records, r)
	}

	for num1, i := range records{
		for num2, j := range records{
			if i.Order < j.Order{
				records[num1], records[num2] = records[num2], records[num1]
			}
		}
	}
	
	return records, nil
}

//GetRecordYearly is a function that use to get all record in input time.
func (m *moneySaver)GetRecordYearly(month int, year int) ([]Record, error){

	var r Record

	var records []Record

	iter := m.client.Collection("record").Where("month", "==", month).Where("year", "==", year).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		doc.DataTo(&r)
		records = append(records, r)
	}

	for num1, i := range records{
		for num2, j := range records{
			if i.Order < j.Order{
				records[num1], records[num2] = records[num2], records[num1]
			}
		}
	}
	
	return records, nil
}