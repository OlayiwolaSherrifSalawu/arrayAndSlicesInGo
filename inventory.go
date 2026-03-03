package main

import "fmt"

type Inventory map[string]int

func (f Inventory) AddItem(item string, amount int) {
	f[item] = amount
}
func (f Inventory) UpdateItem(item string, amount int) error {

	_, ok := f[item]
	if !ok {

		return fmt.Errorf("item is not in the map")
	}
	f[item] += amount
	fmt.Printf("item %s has been updated and now we have %d  \n", item, f[item])
	return nil
}
func (f Inventory) DeleteItem(item string) error {
	// check  item
	_, ok := f[item]
	if !ok {
		return fmt.Errorf("cannot delete item that does not exist \n")
	}
	if f == nil {
		return fmt.Errorf("cannot delete an empty map")
	}
	delete(f, item)
	return nil
}

func (f Inventory) DisplayInventory() {
	for key, value := range f {
		fmt.Printf(" %s : %d ", key, value)
	}
}
