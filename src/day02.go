package main

import (
	"errors"
	"fmt"
)

func day02() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["America"] = "纽约"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["China"]
	if ok {
		fmt.Println("China首都是", capital)
	} else {
		fmt.Println("China首都不存在")
	}

	countryCityMap2 := map[string]string{"china": "北京", "japan": "东京"}
	for key, value := range countryCityMap2 {
		fmt.Println("国家", key, "首都", value)
	}
	delete(countryCityMap2, "china")
	fmt.Println(countryCityMap2)
	a1 := countryCapitalMap["UK"]
	if a1 == "" {
		fmt.Println("133")
	} else {
		fmt.Println("456")
	}
	fmt.Println(a)

	iphone := Iphone{}
	iphone.call()

	nokiaPhone := NokiaPhone{1200}
	nokiaPhone.call()
	fmt.Println(nokiaPhone.money)

	result, err := sqrt(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

type Phone interface {
	call()
}

type NokiaPhone struct {
	money int
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am nokia, I can use very long time !")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am Iphone , I looks beautiful !")
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("math: square root of negative number\"")
	}
	a++
	return a, nil
}
