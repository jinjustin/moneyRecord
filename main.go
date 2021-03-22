package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/jinjustin/moneyRecord/moneySaver"
	"log"
	"fmt"	
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

	/*err = m.NewIncome("Test", 2000, "พ่อให้มา", "เงินพิเศษ")
	if err != nil{
		log.Fatalf("Failed to add new income: %v\n", err)
	}

	err = m.NewIncome("Test", 3000, "ค่าจ้างเขียน BotLine", "เงินพิเศษ")
	if err != nil{
		log.Fatalf("Failed to add new income: %v\n", err)
	}

	err = m.NewExpense("Test", 600, "ค่าโรงแรม", "ค่าเที่ยว")
	if err != nil{
		log.Fatalf("Failed to add new income: %v\n", err)
	}*/

	records ,err := m.GetRecordMonthly(3, 2021)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(records)
}