package moneySaver

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	//"github.com/sqs/goreturns/returns"
	"google.golang.org/api/option"
)

var (
	ctx = context.Background()
)

type moneySaver struct {
	client *firestore.Client
}

func MoneySaver() moneySaver{
	var m moneySaver
	return m
}

//Connect is a function that use to connect to firestore database by using projectID and credentialsFilePath
func (m *moneySaver)Connect(projectID string, credentialsFile string) error{

	conf := &firebase.Config{ProjectID: projectID}
	sa := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		return err
	}

	m.client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	return nil
}

//Close is a function that use to close connection to firebase
func (m *moneySaver)Close() error{
	return m.client.Close()
}

//Test is a function that use to test  connection to firebase by add new record.
func (m *moneySaver)Test() error{
	_, _, err := m.client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "JJ",
		"last":  "TOMSON",
		"born":  1998,
	})
	if err != nil {
		return err
	}
	return nil
}
