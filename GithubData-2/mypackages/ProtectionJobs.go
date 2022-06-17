package mypackages

import (
	"encoding/json"
	//"fmt"
	"strconv"
)

var ProtectionJobsList []string
var ProtectionJobKeys = make(map[string]string)
var Protectiongrouptype = make(map[string]string)

type metrics struct {
	Runs      []run `json:"runs"`
	TotalRuns int   `json:"totalRuns"`
}
type run struct {
	Id              string          `json:"id"`
	BackupInfo      info            `json:"localBackupInfo"`
	ReplicationInfo replicationInfo `json:"replicationInfo"`
	ArchivalInfo    archivalInfo    `json:"archivalInfo"`
}

type info struct {
	StartTimeUsecs  int64  `json:"startTimeUsecs"`
	EndTimeUsecs    int64  `json:"endTimeUsecs"`
	QueuedTimeUsecs int64  `json:"queuedTimeUsecs"`
	Status          string `json:"status"`
	IsSlaViolated   bool   `json:"isSlaViolated"`
}

type replicationInfo struct {
	TargetResults []info `json:"replicationTargetResults"`
}
type archivalInfo struct {
	TargetResults []info `json:"archivalTargetResults"`
}
type jobs struct {
	ProtectionJobs []protectionjob `json:"protectionGroups"`
}

type protectionjob struct {
	JobName string `json:"name"`
	Id      string `json:"id"`
}

func ProtectionJobInfo(jobname string) (jobbackuptime string, jobreplicationtime string, jobarchivaltime string, sla int, runs int) {
	FillProtectionJobKeys()

	jobid := ProtectionJobKeys[jobname]
	timePeriod := "week"
	starttime, endtime := TimeStampGneerator(timePeriod, "U")

	newUrl := GenerateNewURLforProtectionJobs(starttime, endtime, jobid)

	response := PostRequestForAccessToken()

	data := GetRequestForData(response, newUrl)

	var responsedata metrics

	json.Unmarshal(data, &responsedata)

	jobbackuptime = JobBackup(responsedata)
	jobreplicationtime = JobReplication(responsedata)
	jobarchivaltime = JobArchival(responsedata)
	sla = SlaTimes(responsedata)
	runs = responsedata.TotalRuns

	return
}

func FillProtectionJobKeys() {

	response := PostRequestForAccessToken()

	url := "https://10.14.19.226/v2/data-protect/protection-groups?useCachedData=false&pruneSourceIds=true&isDeleted=false&includeTenants=true&includeLastRunInfo=true"

	data := GetRequestForData(response, url)

	var responsedata jobs

	json.Unmarshal(data, &responsedata)

	l := len(responsedata.ProtectionJobs)

	for i := 0; i < l; i++ {
		ProtectionJobsList = append(ProtectionJobsList, responsedata.ProtectionJobs[i].JobName)
		id := responsedata.ProtectionJobs[i].Id
		key := GenerateKey(id)

		ProtectionJobKeys[ProtectionJobsList[i]] = key
	}
}

func GenerateKey(s string) (result string) {

	id := s
	i := len(id) - 1
	var key string

	for id[i] != 58 {
		key = key + string(id[i])
		i -= 1
	}

	for _, v := range key {
		result = string(v) + result
	}
	return
}

func JobBackup(responsedata metrics) (avgbackuptime string) {
	var s []int64
	var e []int64

	for i := 0; i < len(responsedata.Runs); i++ {

		status := responsedata.Runs[i].BackupInfo.Status

		if status == "Succeeded" || status == "SucceededWithWarning" || status == "Failed" {

			s = append(s, responsedata.Runs[i].BackupInfo.StartTimeUsecs)
			e = append(e, responsedata.Runs[i].BackupInfo.EndTimeUsecs)

		} else if status == "Running" {
	}
}
	l := len(s)
	avg := Average(l, s, e)
	convtime := Timeconvert(avg, "M")

	return convtime
}

func JobReplication(responsedata metrics) (avgreptime string) {
	var s []int64
	var e []int64

	for i := 0; i < len(responsedata.Runs); i++ {

		if 0 < len(responsedata.Runs[i].ReplicationInfo.TargetResults) {
			for j := 0; j < (len(responsedata.Runs[i].ReplicationInfo.TargetResults)); j++ {

				status := responsedata.Runs[i].ReplicationInfo.TargetResults[j].Status

				if status == "Succeeded" || status == "SucceededWithWarning" || status == "Failed" {
					s = append(s, responsedata.Runs[i].ReplicationInfo.TargetResults[j].StartTimeUsecs)
					e = append(e, responsedata.Runs[i].ReplicationInfo.TargetResults[j].EndTimeUsecs)
				} else if status == "Running" {}
			}
		} else {}
	}
	l := len(s)
	avg := Average(l, s, e)
	convtime := Timeconvert(avg, "M")

	return convtime
}

func JobArchival(responsedata metrics) (avgarchtime string) {

	//fmt.Println("Runs: ", len(responsedata.Runs))
	var s []int64
	var e []int64

	for i := 0; i < len(responsedata.Runs); i++ {

		//fmt.Println("id = ", responsedata.Runs[i].Id)

		if 0 < len(responsedata.Runs[i].ArchivalInfo.TargetResults) {

			for j := 0; j < (len(responsedata.Runs[i].ArchivalInfo.TargetResults)); j++ {

				status := responsedata.Runs[i].ArchivalInfo.TargetResults[j].Status
				//fmt.Println("Archival Status: ", status)

				if status == "Succeeded" || status == "SucceededWithWarning" {
					//fmt.Println("Archival Queued Time: ", responsedata.Runs[i].ArchivalInfo.TargetResults[j].QueuedTimeUsecs)
					s = append(s, responsedata.Runs[i].ArchivalInfo.TargetResults[j].StartTimeUsecs)
					//fmt.Println("Archival End Time: ", responsedata.Runs[i].ArchivalInfo.TargetResults[j].EndTimeUsecs)
					e = append(e, responsedata.Runs[i].ArchivalInfo.TargetResults[j].EndTimeUsecs)
				} else if status == "Running" {

					//fmt.Println("Archival Start Time: ", responsedata.Runs[i].ArchivalInfo.TargetResults[j].StartTimeUsecs)
					//fmt.Println("Archival Queued Time: ", responsedata.Runs[i].ArchivalInfo.TargetResults[j].QueuedTimeUsecs)

				} else if status == "Failed" {

					//fmt.Println("Archival Queued Time: ", responsedata.Runs[i].ArchivalInfo.TargetResults[j].QueuedTimeUsecs)
				}
			}
		} else {
			//fmt.Println("No Replication data available")
		}
		//fmt.Println()
	}
	l := len(s)
	avg := Average(l, s, e)
	//fmt.Println("average job backup time : ", avg)
	convtime := Timeconvert(avg, "M")

	//fmt.Println("average job backup time after conversion: ", convtime)

	return convtime

}

func SlaTimes(responsedata metrics) (slacount int) {

	slacount = 0

	for i := 0; i < len(responsedata.Runs); i++ {

		slaViolation := responsedata.Runs[i].BackupInfo.IsSlaViolated

		if !slaViolation {
			slacount++
		}
	}
	return slacount
}

func Average(l int, s []int64, e []int64) (avg int64) {

	var k int64
	for i := 0; i < l; i++ {
		k += (e[i] - s[i])
	}

	avg = k / int64(l)

	return
}

func Timeconvert(initialtime int64, unit string) (finaltime string) {

	initialtime /= 1000000
	hrs := initialtime / 3600
	initialtime = initialtime - hrs*3600
	minutes := initialtime / 60
	seconds := initialtime - minutes*60

	hrsstr := strconv.FormatInt(hrs, 10)
	minstr := strconv.FormatInt(minutes, 10)
	secstr := strconv.FormatInt(seconds, 10)

	finaltime = hrsstr + "hr " + minstr + "min " + secstr + "sec"
	return
}
