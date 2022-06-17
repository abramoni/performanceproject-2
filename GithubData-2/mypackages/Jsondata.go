package mypackages

import (
	"encoding/json"
	"fmt"
)

type jsondata struct {
	Id                    string                `json:"id"`
	TimeStamp             int64                 `json:"timeStamp"`
	FileCreateRate        int64                `json:"fileCreateRate"`
	//FileCreateRateUnit string  	`json:"fileCreateRateUnit`
	UtilizationChangeRate int64                `json:"utilizationChangeRate"`
	//UtilizationChangeRateUnit string `json:"utilizationChangeRateUnit"`
	GarbageCollection     int64                `json:"garbageCollection"`
	//GarbageCollectionUnit string `json:"garbageCollectionUnit"`

	ProtectionJobsInfo    []protectiongrouptype `json:"protectionJobsInfo"`
}

type protectiongrouptype struct {
	ProtectionGroupTypeName string   `json:"protectionGroupTypeName"`
	ProtectionGroups        []groups `json:"groups"`
}

type groups struct {
	ProtectionGroupName string `json:"protectionGroupName"`
	JobBackuptime       string  `json:"jobBackuptime"`
	JobReplicationTime  string  `json:"jobReplicationTime"`
	JobArchivalTime     string  `json:"jobArchivalTime"`
	SlaTime             int    `json:"slaTimes"`
	Runs                int    `json:"runs"`
}

func GenerateJson()(data []byte){

	x := [5]int{6, 4, 4, 3, 1}
	y := [5]string{"NAS","VirtualMachines","PhysicalServers","DataBases","CohesityViews"}
	z := [18]string{"GenericNAS_Large_Files","GenericNAS_Small_Files","netapp_NAS_Small_Files", "netappNAS_Large_Files","PureNAS_Large_Files","PureNAS_Small_Files","HyperV_job18","VMWare_Linux","VMware_Windows","VMWare_Windows1","Physical_Linux_BlockBased","Physical_Linux_FileBased","Physical_Windows_BlockBased","Physical_Windows_FileBased ","Oracle","SQL_Volume_Based","SQLFileBased","smartfiles" }
	count := 0
	var protectionjobs []protectiongrouptype

	for i :=0;i<5;i++{
		var mygroups []groups
	for j:=0; j<x[i] ;j++{
		jbt,jrt,jat,salt,run := ProtectionJobInfo(z[count])
		group := groups{
		ProtectionGroupName: z[count],
		JobBackuptime:       jbt,
		JobReplicationTime:  jrt,
		JobArchivalTime:     jat,
		SlaTime:             salt,
		Runs:                run,
		}
		count++

		mygroups = append(mygroups, group)
	}

	mygProtectionGrouptype := protectiongrouptype{
		ProtectionGroupTypeName: y[i],
		ProtectionGroups: mygroups,
	}
	
	protectionjobs = append(protectionjobs, mygProtectionGrouptype)
}

//fcr := Filecreaterate()
ucr := UtilizationChangeRate()
gc := GarbageCollection()
ts := PresentTimeStamp()
myjsondata := jsondata{
	Id:                    "longevitycluster",
	TimeStamp:             ts,
	FileCreateRate:        1234567890,
	//FileCreateRateUnit: "nil",
	UtilizationChangeRate: ucr,
	//UtilizationChangeRateUnit: "GiB/sec",
	GarbageCollection:     gc,
	//GarbageCollectionUnit: "GiB",
	ProtectionJobsInfo:    protectionjobs,
}

p, _ := json.Marshal(myjsondata)
fmt.Printf("%s", p)
	return p
}