package main

import "fmt"

func main() {
	//arrF := []float64{2.12, 2.22, 2.32, 2.42, 2.53}
	//arrF = arrF[:3]
	//total := sum(arrF...)
	//fmt.Println(total)

	ints := []int{1, 2, 3}
	//floats := []float32{4.5, 6.7}
	total, avg := SumAndAverage(ints...)
	fmt.Println("Total:", total)
	fmt.Println("Average:", avg)
}

//func SumAndAverage(dataI ...int, dataF ...float32) (total float32, avg float32) {
//	var intTotal int
//	for _, value := range dataI {
//		intTotal += value
//	}
//	var floatTotal float32
//	for _, value := range dataF {
//		floatTotal += value
//	}
//	return float32(intTotal) + floatTotal, (float32(intTotal) + floatTotal) / 2
//}

func SumAndAverage(dataI ...int) (total float32, avg float32) {
	sumI := 0
	for _, num := range dataI {
		sumI += num
	}

	//sumF := float32(0)
	//for _, num := range dataF {
	//	sumF += num
	//}
	//
	//total = float32(sumI) + sumF
	//avg = total / float32(len(dataI)+len(dataF))
	return total, avg
}

//	func SumAndAverage(data ...interface{}) (int, int, float32, float32) {
//		sumI := 0
//		countI := 0
//		sumF := float32(0)
//		countF := 0
//
//		for _, value := range data {
//			switch v := value.(type) {
//			case int:
//				sumI += v
//				countI++
//			case float32:
//				sumF += v
//				countF++
//			}
//		}
//
//		avgI := float32(sumI) / float32(countI)
//		avgF := sumF / float32(countF)
//
//		return sumI, countI, avgI, avgF
//	}

func sum(data ...float64) float64 {
	var total float64
	for _, value := range data {
		total += value
	}
	return total
}
