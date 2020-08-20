package utils

import (
	"ecommerce/model"

	"github.com/fatih/set"
)

func ArrSetProduct(arr *[]model.Product) ([]int, []int, []int) {
	a := set.New(set.ThreadSafe)
	for i := 0; i < len(*arr); i++ {
		a.Add((*arr)[i].MitraID)
	}
	mitraSet := set.IntSlice(a) 
	var elementProduct []int
	var limit []int
	var counter int
	for j := 0; j < len(mitraSet); j++ {
		for k := 0; k < len(*arr); k++ {
			if uint(mitraSet[j]) == (*arr)[k].MitraID {
				elementProduct = append(elementProduct, int((*arr)[k].ID))
				counter++
			}
		}
		limit = append(limit, counter)
		counter = 0
	} 
	return mitraSet, elementProduct, limit
}

func ArrSetInt(arr []int) ([]int, []int, []int) {
	a := set.New(set.ThreadSafe)
	for i := 0; i < len(arr); i++ {
		a.Add(arr[i])
	}
	arrSet := set.IntSlice(a) 
	var element []int
	var length []int
	var counter int
	for j := 0; j < len(arrSet); j++ {
		for k := 0; k < len(arr); k++ {
			if arrSet[j] == arr[k] {
				element = append(element, arr[k])
				counter++
			}
		}
		length = append(length, counter)
		counter = 0
	} 
	return arrSet, element, length
}
