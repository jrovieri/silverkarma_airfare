package main

import (
	"errors"
	"fmt"
)

func getCityFromCode(lookupCode string) (City, error) {

	for _, item := range cities {
		if item.code == lookupCode {
			return item, nil
		}
	}
	message := fmt.Sprintf("%s is not a valid city code", lookupCode)
	return City{"", "", 0, 0}, errors.New(message)
}

func getCabinClassFromCode(lookupCode string) (CabinClass, error) {

	for _, item := range cabinClasses {
		if item.code == lookupCode {
			return item, nil
		}
	}
	message := fmt.Sprintf("%s is not a valid cabin class code", lookupCode)
	return CabinClass{"", "", 0}, errors.New(message)
}
