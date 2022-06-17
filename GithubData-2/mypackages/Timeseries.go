package mypackages

import "encoding/json"

type metric struct {
	Name    string `json:"metricName"`
	Datavec []rate `json:"dataPointVec"`
}

type rate struct {
	Time  int64     `json:"timestampMsecs"`
	Value ratevalue `json:"data"`
}

type ratevalue struct {
	Data int64 `json:"int64Value"`
}

var endtime, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, starttime string

func Filecreaterate()(avgfilecreaterate int64) {

	metricName = "kCreateFileOps"
	schemaName = "kBridgeViewPerfStats"
	metricUnitType = "5"
	rollupFunction = "rate"
	//entityId = SelectEntity()
	timePeriod = "day"
	starttime, endtime = TimeStampGneerator(timePeriod, "M")
	rollupIntervalSecs = "180"

	avgfilecreaterate = Decodejson()
	return

}

func UtilizationChangeRate()(avgutilrate int64){

	metricName = "kSystemUsageBytes"
	schemaName = "kBridgeClusterStats"	
	metricUnitType ="0"
	rollupFunction = "rate"
	entityId = "2790138600742128"
	timePeriod = "day"
	starttime, endtime = TimeStampGneerator(timePeriod, "M")
	rollupIntervalSecs = "180"
	avgutilrate = Decodejson()
	return
}

func GarbageCollection()(avggc int64){

	metricName = "EstimatedGarbageBytes"
	schemaName = "ApolloV2ClusterStats"	
	metricUnitType ="0"
	rollupFunction = "average"
	entityId = "st-longevity+(ID+2790138600742128)"
	timePeriod = "day"
	starttime, endtime = TimeStampGneerator(timePeriod, "M")
	rollupIntervalSecs = "180"
	avggc = Decodejson()
	return
}



func Decodejson()(avgrate int64) {

	newURL := GenerateNewURL(endtime, entityId, metricName, metricUnitType, timePeriod, rollupFunction, rollupIntervalSecs, schemaName, starttime)

	response := PostRequestForAccessToken()

	data := GetRequestForData(response, newURL)

	var responsedata metric

	json.Unmarshal(data, &responsedata)
	l := len(responsedata.Datavec)
	for i := 0; i < len(responsedata.Datavec); i++ {
		avgrate += responsedata.Datavec[i].Value.Data
	}
	if l>0{
		avgrate /= int64(l)
	}else{
		avgrate = 0
	}
	
	return
}
