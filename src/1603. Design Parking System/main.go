package main

import "fmt"

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {
	obj := Constructor(1, 1, 0)
	param_1 := obj.AddCar(2)
	param_2 := obj.AddCar(2)
	fmt.Println(param_1)
	fmt.Println(param_2)
}

type ParkingSystem struct {
	cards []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{cards: []int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cards[carType-1] > 0 {
		this.cards[carType-1]--
		return true
	}
	return false
}

//
//type ParkingSystem struct {
//	big, medium, small int
//}
//
//func Constructor(big int, medium int, small int) ParkingSystem {
//	return ParkingSystem{big, medium, small}
//}
//
//func (this *ParkingSystem) AddCar(carType int) bool {
//	if carType == 1 {
//		if this.big != 0{
//			this.big--
//			return true
//		}
//	}
//	if carType == 2 {
//		if this.medium != 0{
//			this.medium--
//			return true
//		}
//	}
//	if carType == 3 {
//		if this.small != 0{
//			this.small--
//			return true
//		}
//	}
//	return false
//}
