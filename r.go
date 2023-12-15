package main

import (
	"fmt"
	"strings"
)

// MenuItem represents an item in the menu with its corresponding cook time
type MenuItem struct {
	Name     string
	CookTime int
}

// FoodDeliveryTime calculates the time to prepare the order and returns the total time or an error
func FoodDeliveryTime(order string) (int, error) {
	menu := []MenuItem{
		{"burger", 15},
		{"chips", 10},
		{"nuggets", 12},
	}

	orders := strings.Split(order, ",")
	totalTime := 0
	found := false

	for _, item := range orders {
		item = strings.TrimSpace(item)
		found = false

		for _, menuItem := range menu {
			if strings.ToLower(item) == menuItem.Name {
				totalTime += menuItem.CookTime
				found = true
				break
			}
		}

		if !found {
			fmt.Errorf("404 - Item '%s' not found in the menu", item)
		}
	}

	return totalTime, nil
}

func main() {
	order1 := "burger, chips, nuggets"
	order2 := "burger, fries, nuggets"

	time1, err1 := FoodDeliveryTime(order1)
	time2, err2 := FoodDeliveryTime(order2)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("Total time for order 1: %d minutes\n", time1)
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Total time for order 2: %d minutes\n", time2)
	}
}
