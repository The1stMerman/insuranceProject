package main

import "fmt"

type insurance interface {
	insurancePrice() float64
}

type carInsurance struct {
	brand string
	model string
}

type travelInsurance struct {
	country string
	day     int
}

type healthInsurance struct {
	sex string
	age int
	job string
}

func main() {
	car := carInsurance{"Toyota", "Altis"}
	travel := travelInsurance{"Italy", 7}
	health := healthInsurance{"Male", 60, "Driver"}

	fmt.Println("ประกันรถ :", printInsurancePrice(car))
	fmt.Println("ประกันการเดินทาง :", printInsurancePrice(travel))
	fmt.Println("ประกันอุบัติเหตุ :", printInsurancePrice(health))
}

func (c carInsurance) insurancePrice() float64 {
	carRates := map[string]map[string]float64{
		"Toyota": {
			"Yaris": 1000.00,
			"Altis": 1500.00,
			"Camry": 2000.00,
		},
		"Honda": {
			"Accord": 1000.00,
			"City":   1500.00,
			"Civic":  2000.00,
		},
	}
	return carRates[c.brand][c.model]
}

func (t travelInsurance) insurancePrice() float64 {
	countryRates := map[string]float64{
		"Japan":   1500.00,
		"England": 2000.00,
		"US":      2500.00,
		"Italy":   3000.00,
	}
	return countryRates[t.country] * float64(t.day)
}

func (h healthInsurance) insurancePrice() float64 {
	sexRates := map[string]float64{
		"Male":   2,
		"Female": 1.5,
	}
	jobRates := map[string]float64{
		"Driver": 300.00,
		"SE":     200.00,
		"HR":     100.00,
	}
	return sexRates[h.sex]*jobRates[h.job] + float64(h.age)
}

func printInsurancePrice(i insurance) float64 {
	return i.insurancePrice()
}
