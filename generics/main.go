package main

import "fmt"
type Number interface {
    ~int | ~int64 | ~float64 | ~float32
}
func main() {
result := addUsingFirstApproachBasic(1, 2)
fmt.Println(result)
result2 := addUsingSecondApproachUsingGenericBasic(1, 2)
fmt.Println(result2)
result3 := addUsingThirdApproachBest(1, 2)
fmt.Println(result3)

}
func addUsingFirstApproachBasic(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int);
	bInt, bIsInt := b.(int);
	if aIsInt && bIsInt {
		return aInt + bInt
	}
	return nil
	// similarly for other types we can do this
}
func addUsingSecondApproachUsingGenericBasic[T, any](a,b T) T {
	return a + b
}
func addUsingThirdApproachBest[T int | float64 | string | bool](a,b T) T {
	return a + b;
}
func addUsingFinalApproach[T Number](a,b T) T {
	return a + b;
}