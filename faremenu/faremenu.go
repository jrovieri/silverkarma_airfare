package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)

	var srcCity, dstCity City
	var err error
	var cabin CabinClass

	var isCityValid bool = false
	var isDestValid bool = false
	var isCabinValid bool = false

	fmt.Println("SilverKarma Airfare Calculator")

	for !isCityValid {

		var code string
		fmt.Print("Enter origin code: ")
		fmt.Scanln(&code)

		srcCity, err = getCityFromCode(strings.ToUpper(code))
		if err == nil {
			fmt.Println("You've entered: " + srcCity.cityName)
			isCityValid = true
		} else {
			fmt.Println(err)
		}

	}

	for !isDestValid {

		var code string
		fmt.Print("Enter destination code: ")
		fmt.Scanln(&code)

		dstCity, err = getCityFromCode(strings.ToUpper(code))
		if err == nil {
			fmt.Println("You've entered: " + dstCity.cityName)
			isDestValid = true
		} else {
			fmt.Println(err)
		}
	}

	for !isCabinValid {

		fmt.Println("The available classes are: ")
		fmt.Fprintln(w, "Code\tClass\tRate")

		for _, c := range cabinClasses {
			line := fmt.Sprintf("%s\t%s\t%d\n", c.code, c.className, c.rate)
			fmt.Fprint(w, line)
		}
		w.Flush()

		var code string
		fmt.Print("Enter class code: ")
		fmt.Scanln(&code)

		cabin, err = getCabinClassFromCode(strings.ToUpper(code))
		if err == nil {
			fmt.Println("You've entered: " + cabin.className)
			isCabinValid = true
		} else {
			fmt.Println(err)
		}
	}

	srcLng := float64(srcCity.longitude) / 10000
	srcLat := float64(srcCity.latitude) / 10000
	dstLng := float64(dstCity.longitude) / 10000
	dstLat := float64(dstCity.latitude) / 10000

	distance := CalculateDistance(dstLng, dstLat, srcLng, srcLat)
	fmt.Printf("\nDistance  : %.1f km", distance)

	rate := float32(cabin.rate) / 100
	fmt.Printf("\n$ per km  : %.2f", rate)

	total := rate * float32(distance)
	fmt.Printf("\nTotal fare: %.2f\n", total)
}
