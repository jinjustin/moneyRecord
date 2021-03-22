package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/jinjustin/moneyRecord/moneySaver"
	"log"	
)

var (
	projectID = "moneyrecord-7ef16"
	credentialsFile = "./moneySaver/moneyrecord-7ef16-firebase-adminsdk-ymi9q-56a433392f.json"
)

func main() {
	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()*/

	m := moneySaver.MoneySaver()

	err := m.Connect(projectID, credentialsFile)
	if err != nil{
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer m.Close()

	err = m.NewIncome("Test Account", 2000,"พ่อให้มา", "ค่าขนม")
	if err != nil{
		log.Fatalf("Failed to add new income: %v\n", err)
	}

	err = m.NewIncome("Test Account", 3000,"เงินเดือน", "เงินเดือน")
	if err != nil{
		log.Fatalf("Failed to add new income: %v\n", err)
	}
}