package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContactsFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var contacts []Contact

	err := loadContactsFile(&contacts)

	if err != nil {
		fmt.Println("Error loading contacts file:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("=== Contact Manager ===\n",
			"1. Add contact\n",
			"2. List contacts\n",
			"3. Exit\n",
			"Choose an option: ")

		var option int
		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Invalid option")
			continue
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			err = saveContactsFile(contacts)

			if err != nil {
				fmt.Println("Error saving contacts file:", err)
				return
			}
		case 2:
			fmt.Println("=== Contacts ===")
			for i, c := range contacts {
				fmt.Printf("%d. Name: %s - Email: %s - Phone: %s\n", i+1, c.Name, c.Email, c.Phone)
			}
			fmt.Println("=== End of contacts ===")
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}

	}

}
