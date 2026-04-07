package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Inventory struct {
	Items []Item
}

// add function

func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
}

// remove function

func (inv *Inventory) RemoveItem(id int) bool {
	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return true
		}
	}
	return false
}

// Update function
func (inv *Inventory) UpdateItem(id int, name string, qty int, price float64) bool {
	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items[i].Name = name
			inv.Items[i].Quantity = qty
			inv.Items[i].Price = price
			return true
		}
	}
	return false
}

// search function

func (inv *Inventory) SearchItem(id int) (*Item, bool) {
	for i, item := range inv.Items {
		if item.ID == id {
			return &inv.Items[i], true
		}
	}
	return nil, false
}

// json loading

func (inv *Inventory) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(inv.Items, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

//unloading

func (inv *Inventory) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil // ignore if file doesn't exist
	}
	return json.Unmarshal(data, &inv.Items)
}

func main() {
	inv := Inventory{}
	file := "inventory.json"

	// Load existing data
	inv.LoadFromFile(file)

	for {
		fmt.Println("\n--- Inventory Menu ---")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. Update Item")
		fmt.Println("4. Search Item")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var id, qty int
			var name string
			var price float64

			fmt.Print("Enter ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter Name: ")
			fmt.Scan(&name)

			fmt.Print("Enter Quantity: ")
			fmt.Scan(&qty)

			fmt.Print("Enter Price: ")
			fmt.Scan(&price)

			item := Item{id, name, qty, price}

			inv.AddItem(item)
			fmt.Println("Item added!")

		case 2:
			var id int
			fmt.Print("Enter ID to remove: ")
			fmt.Scan(&id)

			if inv.RemoveItem(id) {
				fmt.Println("Item removed!")
			} else {
				fmt.Println("Item not found!")
			}

		case 3:
			var id, qty int
			var name string
			var price float64

			fmt.Print("Enter ID to update: ")
			fmt.Scan(&id)

			fmt.Print("Enter new Name: ")
			fmt.Scan(&name)

			fmt.Print("Enter new Quantity: ")
			fmt.Scan(&qty)

			fmt.Print("Enter new Price: ")
			fmt.Scan(&price)

			if inv.UpdateItem(id, name, qty, price) {
				fmt.Println("Item updated!")
			} else {
				fmt.Println("Item not found!")
			}

		case 4:
			var id int
			fmt.Print("Enter ID to search: ")
			fmt.Scan(&id)

			item, found := inv.SearchItem(id)
			if found {
				fmt.Printf("Found: %+v\n", *item)
			} else {
				fmt.Println("Item not found!")
			}

		case 5:
			inv.SaveToFile(file)
			fmt.Println("Data saved. Exiting...")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
