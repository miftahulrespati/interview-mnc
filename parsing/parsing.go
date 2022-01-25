package parsing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Order struct {
	OrderId       int64       `json:"order_id"`
	OrderDate     string      `json:"order_date"`
	Customer      Customer    `json:"customer"`
	Items         []Items     `json:"items"`
	Discounts     []Discounts `json:"discounts"`
	ShippingPrice float64     `json:"shipping_price"`
}

type Customer struct {
	CustomerId      int64           `json:"customer_id"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	Email           string          `json:"email"`
	Phone           string          `json:"phone"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
	Suburb   string `json:"suburb"`
	State    string `json:"state"`
}

type Items struct {
	Quantity  int64   `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Product   Product `json:"product"`
}

type Product struct {
	ProductId int64    `json:"product"`
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Image     string   `json:"image"`
	Thumbnail string   `json:"thumbnail"`
	Category  []string `json:"category"`
	Url       string   `json:"url"`
	Upc       string   `json:"upc"`
	Gtin14    string   `json:"gtin14"`
	CreatedAt string   `json:"created_at"`
	Brand     Brand    `json:"brand"`
}

type Brand struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Discounts struct {
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
	Priority int64   `json:"priority"`
}

func Parsing() (orders []Order) {
	jsonFile, err := os.Open("test.jsonl")
	log.Println("Opening test.jsonl file ...")

	if err != nil {
		fmt.Println(err)
	}
	log.Println("File opened")
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	decode := json.NewDecoder(jsonFile)

	for {
		var order Order
		err := decode.Decode(&order)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	return orders
}
