package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip,omitempty"`
}

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
	IsActive bool    `json:"is_active,omitempty"`
}

func employeeToJSON(e Employee) (string, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func employeesToJSON(employees []Employee) (string, error) {
	data, err := json.Marshal(employees)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func JSONToEmployee(data string) (Employee, error) {
	var e Employee
	err := json.Unmarshal([]byte(data), &e)
	if err != nil {
		return e, err
	}
	return e, nil
}

func JSONToEmployees(data string) ([]Employee, error) {
	var employees []Employee
	err := json.Unmarshal([]byte(data), &employees)
	if err != nil {
		return employees, err
	}
	return employees, nil
}

func main() {
	worker := Employee{
		ID:   1832,
		Name: "Sergey Shillelagh",
		Age:  20,
		Address: Address{
			Street: "ul. Solnechnaya, d. 23",
			City:   "Krasnodar",
		},
		IsActive: true,
	}
	data, err := employeeToJSON(worker)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

	jsonString := `{"id":1833,"name":"Alice Fox","age":23,"address":{"street":"ul. Solnechnaya, d. 33","city":"Krasnodar","zip":"350000"}}`
	worker2, err := JSONToEmployee(jsonString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(worker2)

	workersSlice := []Employee{
		{
			ID:   1834,
			Name: "Alex Kazakh",
			Age:  20,
			Address: Address{
				Street: "ul. Severnaya, d. 13",
				City:   "Moscow",
				Zip:    "101000",
			},
			IsActive: true,
		},
		{
			ID:   1835,
			Name: "Ivan Bogatyr",
			Age:  32,
			Address: Address{
				Street: "ul. Lenina, d. 42",
				City:   "Tula",
			},
			IsActive: false,
		},
	}
	sliceData, err := employeesToJSON(workersSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sliceData)

	employeesSlice, err := JSONToEmployees(sliceData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employeesSlice)

	anotherJsonString := `{"id":1488,"name":"Nick Nick","age":43,"address":{"street":"ul. Solnechnaya, d. 33","city":"Krasnodar","zip":"350000"},"department":"NYPD"}`
	employee, err := JSONToEmployee(anotherJsonString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employee)
}
