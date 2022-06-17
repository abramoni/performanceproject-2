package mypackages

import (
	"fmt"
	"strconv"
	"time"
)

func TimeStampGneerator(timeperiod string, unit string) (startTime string, endTime string) {

	tNow := time.Now()
	minute := tNow.Minute()
	hour := tNow.Hour()
	second := tNow.Second()
	unixdec := second + 60*minute + 3600*hour
	unix := tNow.Unix()
	unixd := int64(unixdec)
	unixend := unix - unixd
	unixstart := unixend - 86400
	
	if timeperiod == "past4hrs"{
		unixend = tNow.Unix()
		unixstart = unixend - 4*3600
	}

	if timeperiod == "day"{
		
	}
	if timeperiod == "week" {
		unixstart = unixstart - 6*86400
	}
	if timeperiod == "month"{
		unixstart = unixstart - 29*86400
	}
	if unit == "M"{
	unixstart = unixstart * 1000
	unixend = unixend * 1000
	}
	if unit == "U"{
		unixstart = unixstart * 1000000
		unixend = unixend * 1000000
	}
	s := strconv.FormatInt(unixstart, 10)
	e := strconv.FormatInt(unixend, 10)
	return s, e
}

func PresentTimeStamp ()(timestamp int64){
	tNow := time.Now()
	timestamp = tNow.Unix()
	return
}

func Range() (Timeperiod string) {

	fmt.Println("Select Range")
	fmt.Println("1 - 24 hours")
	fmt.Println("2 - 7 days")
	fmt.Println("3 - 30 days")

	var option int
	fmt.Scanln(&option)

	var period string

	switch option {
	case 1:
		period = "day"
	case 2:
		period = "week"
	case 3:
		period = "month"
	}
	return period

}


func RollupIntervalSecs() (rollupint string) {

	fmt.Println("Select interval")
	fmt.Println("1 - 3 minutes")
	fmt.Println("2 - 15 minutes")
	fmt.Println("3 - 30 minutes")
	fmt.Println("4 - 1hr")
	fmt.Println("5 - 1day")
	fmt.Println("6 - custom")
	fmt.Println("Enter choic : ")

	var option int
	fmt.Scanln(&option)

	var interval string

	switch option {
	case 1:
		interval = "180"
	case 2:
		interval = "900"
	case 3:
		interval = "1800"
	case 4:
		interval = "3600"
	case 5:
		interval = "86400"
	case 6:
		fmt.Println("Enter interval in seconds : ")
		fmt.Scanln(&interval)
	}
	return interval

}
