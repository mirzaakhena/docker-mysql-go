package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
	"github.com/spf13/viper"
)

// Product is
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	// wait 20 second
	fmt.Printf("Wait 20 second before mysql is ready to run\n")
	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d second\n", i+1)
	}

	// config
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	cs := viper.GetString("database.connection_string")
	fmt.Printf("Using connection string %v\n", cs)

	// database setup
	db, err := gorm.Open("mysql", cs)
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()
	fmt.Println("Success connect")

	// database insert
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "L1212", Price: 1000})
	var products []Product
	db.Find(&products)
	bytes, _ := json.Marshal(products)
	fmt.Printf("Data %v\n", string(bytes))

	// rest API
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"message":  "hello",
			"products": products,
		})
	})
	r.Run()

}
