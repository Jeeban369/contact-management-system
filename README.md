# Contact Management System

A simple **Contact Management System** built with Go. This program lets you manage contacts by adding, viewing, searching, and deleting them. It’s beginner-friendly and demonstrates basic programming concepts in Go.

---

## Features

1. **Add Contact**: Create a new contact with a name, phone number, and email address.
2. **View Contacts**: Display all saved contacts.
3. **Search Contact**: Search for a contact by name.
4. **Delete Contact**: Remove a contact from the list using its name.

---

## Code Overview

### Contact Struct
The `Contact` struct represents an individual contact with the following fields:
- `Name` (string): The contact’s name.
- `Phone` (string): The contact’s phone number.
- `Email` (string): The contact’s email address.

### ContactManager Struct
The `ContactManager` struct handles the operations related to managing contacts and contains:
- `contacts` ([]Contact): A slice to store multiple contacts.

### Methods

#### `AddContact`
Adds a new contact to the manager.
```go
func (cm *ContactManager) AddContact(name, phone, email string)
```

#### `ViewContacts`
Displays all the saved contacts.
```go
func (cm *ContactManager) ViewContacts()
```

#### `SearchContact`
Finds a contact by its name (case-insensitive).
```go
func (cm *ContactManager) SearchContact(name string)
```

#### `DeleteContact`
Removes a contact by its name.
```go
func (cm *ContactManager) DeleteContact(name string)
```

---

## How to Run

1. **Install Go**: Ensure Go is installed on your system. Download it from [Go's official website](https://golang.org/dl/).

2. **Clone the Code**: Copy the code into a `.go` file, for example, `contact_manager.go`.

3. **Run the Program**:
```bash
go run contact_manager.go
```

---

## Example Output

```
Added contact: Alice
Added contact: Bob

Contacts:
Name: Alice, Phone: 1234567890, Email: alice@example.com
Name: Bob, Phone: 9876543210, Email: bob@example.com

Searching for contact: Alice
Found: Name: Alice, Phone: 1234567890, Email: alice@example.com

Deleted contact: Bob

Contacts:
Name: Alice, Phone: 1234567890, Email: alice@example.com
```

---

## Why This Project?

- **Beginner-Friendly**: Ideal for those starting with Go.
- **Core Concepts**: Demonstrates structs, slices, and basic methods.
- **Expandable**: Can be extended with additional features like updating a contact or saving data to a file.


