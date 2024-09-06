package programs

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName, LastName string
	Age                 int
}

type ByAge []User
type ByFirstName []User

func (users ByAge) Len() int {
	return len(users)
}

func (users ByAge) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func (users ByAge) Less(i, j int) bool {
	return users[i].Age < users[j].Age
}

type ByName []User

func (users ByName) Len() int {
	return len(users)
}

func (users ByName) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func (users ByName) Less(i, j int) bool {

	if users[i].FirstName == users[j].FirstName {
		return users[i].LastName < users[j].LastName
	}
	return users[i].FirstName < users[j].FirstName
}

func SortByFields() {

	userList := []User{
		{
			FirstName: "Vallabh",
			LastName:  "Lakade",
			Age:       33,
		},
		{
			FirstName: "Ashish",
			LastName:  "Kulkarni",
			Age:       32,
		},
		{
			FirstName: "Somnath",
			LastName:  "Lawand",
			Age:       34,
		},
	}

	sort.Sort(ByAge(userList))
	fmt.Println("Sorting by age: ", ByAge(userList))

	sort.Sort(ByName(userList))
	fmt.Println("Sorting by name:", userList)

}
