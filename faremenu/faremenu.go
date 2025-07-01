package main

import "fmt"

func main() {

	var origin, dest, cabin string
	var isCityValid bool = false
	var isDestValid bool = false
	var isCabinValid bool = false

	fmt.Println("SilverKarma Airfare Calculator")

	for !isCityValid {

		fmt.Print("Enter origin code: ")
		fmt.Scanln(&origin)

		city, err := getCityFromCode(origin)
		if err == nil {
			fmt.Println("You've entered: " + city.cityName)
			isCityValid = true
		} else {
			fmt.Println(err)
		}

	}

	for !isDestValid {

		fmt.Print("Enter destination code: ")
		fmt.Scanln(&dest)

		city, err := getCityFromCode(dest)
		if err == nil {
			fmt.Println("You've entered: " + city.cityName)
			isDestValid = true
		} else {
			fmt.Println(err)
		}
	}

	for !isCabinValid {

		fmt.Print("Enter class code: ")
		fmt.Scanln(&cabin)

		cabinClass, err := getCabinClassFromCode(cabin)
		if err == nil {
			fmt.Println("You've entered: " + cabinClass.className)
			isCabinValid = true
		} else {
			fmt.Println(err)
		}
	}

}
