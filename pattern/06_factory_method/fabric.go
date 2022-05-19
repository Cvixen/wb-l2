package main

import "fmt"

func getTransport(tt string) (iTransport, error) {
	if tt == "scooter" {
		return newElectricScooter(), nil
	}
	if tt == "quadcopter" {
		return newQuadcopter(), nil
	}
	return nil, fmt.Errorf("Wrong type")
}
