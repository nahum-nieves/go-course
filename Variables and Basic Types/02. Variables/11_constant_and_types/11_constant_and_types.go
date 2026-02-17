package main
import "fmt"

func main() {
	//Untyped constants are constants that do not have a specific
	//  type and can be used in different contexts without requiring explicit type conversion
	const rewardPoints = 10
	//The output of the following line will show that rewardPoints is an int
	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints)
	var totalRewardPoints float64 = 150.5
	//Constants can adapt to the type of the context in which they are used,
	//  allowing for flexibility in calculations and comparisons
	totalRewardPoints += rewardPoints
	fmt.Printf("Updated loyalty points %.2f\n", totalRewardPoints)
}