package main

import (
    "fmt"
)

func main() {

	var servers int
	_, _ = fmt.Scan(&servers)
	var KumSumPercent float64 = 0
	mass := make([]float64, 0, servers)

    for i:=0;i<servers;i++{
		var PercentUsingServer, PercentFailureOnServer int
		_, _ = fmt.Scan(&PercentUsingServer, &PercentFailureOnServer)
		tmp := int64((float64(PercentUsingServer)*float64(PercentFailureOnServer)/10000)*1000000000)
		mass = append(mass,float64(tmp)/1000000000)
		KumSumPercent += mass[i]
	}
	for i:=0;i<servers;i++{
		fmt.Printf("%.12f\n",mass[i]/KumSumPercent)
	}

}
