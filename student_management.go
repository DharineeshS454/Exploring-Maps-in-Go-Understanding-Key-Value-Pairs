package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Student represents a student with their details
type Student struct {
	Name   string
	Age    int
	Grade  string
	Gender string
}

func main() {
	students := make(map[int]Student)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nStudent Management System")
		fmt.Println("1. Add Student")
		fmt.Println("2. Delete Student")
		fmt.Println("3. Display Students")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			break
		}

		choice := scanner.Text()
		switch choice {
		case "1":
			addStudent(scanner, students)
		case "2":
			deleteStudent(scanner, students)
		case "3":
			displayStudents(students)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Function to add a new student record
func addStudent(scanner *bufio.Scanner, students map[int]Student) {
	fmt.Println("\nAdd Student")
	fmt.Print("Enter student ID: ")
	idStr := getInput(scanner)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid integer.")
		return
	}

	if _, exists := students[id]; exists {
		fmt.Println("Student with ID", id, "already exists.")
		return
	}

	fmt.Print("Enter student name: ")
	name := getInput(scanner)

	fmt.Print("Enter student age: ")
	ageStr := getInput(scanner)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid integer.")
		return
	}

	fmt.Print("Enter student grade: ")
	grade := getInput(scanner)

	fmt.Print("Enter student gender: ")
	gender := getInput(scanner)

	students[id] = Student{Name: name, Age: age, Grade: grade, Gender: gender}
	fmt.Println("Student added successfully.")
}

// Function to delete a student record
func deleteStudent(scanner *bufio.Scanner, students map[int]Student) {
	fmt.Println("\nDelete Student")
	fmt.Print("Enter student ID to delete: ")
	idStr := getInput(scanner)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid integer.")
		return
	}

	if _, exists := students[id]; !exists {
		fmt.Println("Student with ID", id, "does not exist.")
		return
	}

	delete(students, id)
	fmt.Println("Student deleted successfully.")
}

// Function to display all student records
func displayStudents(students map[int]Student) {
	fmt.Println("\nStudent Records:")
	for id, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %s, Gender: %s\n", id, student.Name, student.Age, student.Grade, student.Gender)
	}
}

// Function to get input from the user
func getInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
