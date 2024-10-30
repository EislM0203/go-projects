package main

import "fmt"

func main() {
    // 1) Create a new array containing three hobbies
    hobbies := [3]string{"Homelabing", "Cooking", "Hiking"}
    fmt.Println("Hobbies array:", hobbies)
    
    // 2) Output more data about that array
    fmt.Println("First hobby:", hobbies[0])
    secondAndThird := hobbies[1:3]
    fmt.Println("Second and third hobbies:", secondAndThird)
    
    // 3) Create a slice based on the first element that contains the first and second elements
    slice1 := hobbies[:2]
    fmt.Println("Slice 2 (First and second elements):", slice1)
    
    // 4) Re-slice the slice to contain the first and second element
    resliced := slice1[1:3]
    fmt.Println("Re-sliced to second and last element:", resliced)
    
    // 5) Create a dynamic array with your course goals
    goals := []string{"Learn Go", "Build a project"}
    fmt.Println("Initial goals:", goals)
    
    // 6) Set the second goal to a different one and add a third goal
    goals[1] = "Master concurrency"
    goals = append(goals, "Contribute to open source")
    fmt.Println("Updated goals:", goals)
    
    // 7) Bonus: Create a "Product" struct and dynamic list of products
    type Product struct {
        Title string
        ID    int
        Price float64
    }
    
    products := []Product{
        {"Laptop", 1, 999.99},
        {"Smartphone", 2, 499.99},
    }
    
    // Adding a third product
    products = append(products, Product{"Tablet", 3, 299.99})
    fmt.Println("Products:", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.