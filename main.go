package main

import (
	"encoding/json"
	"fmt"
	"interview-mnc/parsing"
	"io/ioutil"
	"log"
)

// type Orders struct {
// 	Orders []Order
// }

type Order struct {
	OrderId             int      `json:"order_id"`
	OrderDate           string   `json:"order_date"`
	CustomerData        Customer `json:"customer_data"`
	ItemsQuantity       int      `json:"items_quantity"`
	PriceBeforeDiscount float64  `json:"price_before_discount"`
	Discounts           float64  `json:"discounts"`
	PriceAfterDiscount  float64  `json:"price_after_discount"`
	PriceWithShipping   float64  `json:"price_with_shipping"`
}

type Customer struct {
	CustomerId      int    `json:"customer_id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	ShippingAddress string `json:"shipping_address"`
}

func main() {
	parse := parsing.Parsing()
	orders := make([]Order, len(parse))

	log.Println("Parsing jsonl file into expected struct ...")

	for i := 0; i < len(parse); i++ {
		orders[i].OrderId = parse[i].OrderId
		orders[i].OrderDate = parse[i].OrderDate
		orders[i].CustomerData.CustomerId = parse[i].Customer.CustomerId
		orders[i].CustomerData.Name = parse[i].Customer.FirstName + " " + parse[i].Customer.LastName
		orders[i].CustomerData.Email = parse[i].Customer.Email
		orders[i].CustomerData.Phone = parse[i].Customer.Phone
		orders[i].CustomerData.ShippingAddress = parse[i].Customer.ShippingAddress.Street + ", " + parse[i].Customer.ShippingAddress.Suburb + ", " + parse[i].Customer.ShippingAddress.State + ", Postcode: " + parse[i].Customer.ShippingAddress.Postcode
		for j := 0; j < len(parse[i].Items); j++ {
			orders[i].ItemsQuantity += parse[i].Items[j].Quantity
			orders[i].PriceBeforeDiscount += (float64(parse[i].Items[j].Quantity) * (parse[i].Items[j].UnitPrice))
		}
		for k := 0; k < len(parse[i].Discounts); k++ {
			if parse[i].Discounts[k].Type == "DOLLAR" {
				orders[i].Discounts += parse[i].Discounts[k].Value
			} else if parse[i].Discounts[k].Type == "PERCENTAGE" {
				orders[i].Discounts += ((parse[i].Discounts[k].Value / 100) * orders[i].PriceBeforeDiscount)
			}
		}
		orders[i].PriceAfterDiscount = orders[i].PriceBeforeDiscount - orders[i].Discounts
		orders[i].PriceWithShipping = orders[i].PriceAfterDiscount + parse[i].ShippingPrice
	}

	json, err := json.MarshalIndent(orders, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Creating new json file from struct ...")

	if ioutil.WriteFile("result.json", json, 0644) != nil {
		fmt.Println(err)
		return
	}
	log.Println("New file: result.json created! Check your directory")
}
