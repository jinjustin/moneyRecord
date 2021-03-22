package moneySaver

import (
	//"errors"
	"google.golang.org/api/iterator"
)

type Account struct{
	Total int `json:"total"`
	Order int `json:"order"`
}

//CreateAccount is a function that use to create new account to store money.
func (m *moneySaver)CreateAccount(accountName string) error{

	a := Account{
		Total: 0,
		Order: 0,
	}

	_, err := m.client.Collection("account").Doc(accountName).Set(ctx, map[string]interface{}{
        "Total": a.Total,
		"Order": a.Order,
	})
	if err != nil {
		return err
	}
	return nil
}

//GetAllAccount is a function that use to get all account.
func (m *moneySaver)GetAllAccount(accountName string) ([]Account, error){

	var accounts []Account

	var a Account

	iter := m.client.Collection("account").Documents(ctx)
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			doc.DataTo(&a)
			accounts = append(accounts, a)
	}
	return accounts, nil
}

//Deposit is a function use to deposit money to account.
/*func (m *moneySaver)Deposit(accountName string, value int) error{

	dsnap, err := m.client.Collection("account").Doc(accountName).Get(ctx)
	if err != nil{
		return err
	}

	var a Account
	dsnap.DataTo(&a)

	_, err = m.client.Collection("account").Doc(accountName).Set(ctx, map[string]interface{}{
        "Total": a.Total + value,
		"Order": a.Order + 1,
	})
	if err != nil {
		return err
	}
	return nil
}

//Withdraw is a function use to withdraw money from account.
func (m *moneySaver)Withdraw(accountName string, value int) error{

	dsnap, err := m.client.Collection("account").Doc(accountName).Get(ctx)
	if err != nil{
		return err
	}

	var a Account
	dsnap.DataTo(&a)
	if a.Total - value >= 0 {
		_, err = m.client.Collection("account").Doc(accountName).Set(ctx, map[string]interface{}{
			"Total": a.Total - value,
			"Order": a.Order + 1,
		})
		if err != nil {
			return err
		}
	}else{
		return errors.New("value: too much withdraw value.")
	}
	return nil
}*/

func (m *moneySaver)updateAccount(accountName string, total int, order int) error{

	_, err := m.client.Collection("account").Doc(accountName).Set(ctx, map[string]interface{}{
		"Total": total,
		"Order": order,
	})
	if err != nil {
		return err
	}
	return nil
}