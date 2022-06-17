package main

import (
	"fmt"
	"presentation/mypackages"
)


func main(){
	fmt.Println("Collecting the metrics please wait...")
	data := mypackages.GenerateJson()
	fmt.Println()
	mypackages.StoreInElasticDb(data,"longevitycluster")
}