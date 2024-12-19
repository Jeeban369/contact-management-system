package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}
type ContactManager struct {
	contacts []Contact
}

func (cm *ContactManager) AddContact(name, email, phone string) {
	contact := Contact{Name: name, Email: email, Phone: phone}
	cm.contacts = append(cm.contacts, contact)
	fmt.Printf("Added contact: %s\n", name)
}

func (cm *ContactManager) Viewcotact() {
	if len(cm.contacts) == 0 {
		fmt.Println("No contacts available.")
		return
	}
	fmt.Println()
	for i, contact := range cm.contacts {
		fmt.Printf("(%d) Name: %s, Phone: %s, Email: %s\n", i, contact.Name, contact.Phone, contact.Email)
	}
}

func (cm *ContactManager) SearchContact(name string) {
	fmt.Println()
	fmt.Println("Searching for contact-", name)
	found := false

	for _, contact := range cm.contacts {
		if strings.EqualFold(contact.Name, name) {
			fmt.Printf("Found: Name: %s, Phone: %s, Email: %s\n", contact.Name,contact.Phone,contact.Email)
			found = true
		}
	}

	if !found{
		fmt.Println("No contact found with that name.")
		fmt.Println()
	}
}

func (cm *ContactManager) DeleteContact(name string){
	for i, contact := range cm.contacts{
		if strings.EqualFold(contact.Name, name){
			cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
			fmt.Println("Deleted contact:",name)
			return
		}
	}
	fmt.Println("No contact found with that name.")
}

func main() {
	cm := ContactManager{}

	cm.AddContact("jeeban", "93487", "papu@yahoo.com")
	cm.AddContact("Alice", "1234567890", "alice@example.com")
	cm.AddContact("Bob", "9876543210", "bob@example.com")

	cm.Viewcotact()

	cm.SearchContact("jeeban")

	cm.DeleteContact("alice")
	cm.SearchContact("alice")
	cm.SearchContact("bob")

}
