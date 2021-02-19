// Eugene Simos
// utility to load data /clean up from data a logging analytics OCI compartement or a series of loggan compartments under a "main compartment"

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v35/common"
	_"github.com/oracle/oci-go-sdk/v35/example/helpers"
	_"github.com/oracle/oci-go-sdk/v35/identity"
	"github.com/oracle/oci-go-sdk/v35/loganalytics"
	_"github.com/oracle/oci-go-sdk/v35/managementdashboard"
	"io/ioutil"
	_ "log"
	"logan/utils"
	"os"
 	"sync"
)

var (
	e1 = common.CustomProfileConfigProvider("/home/oracle/.oci/config", "DBSECN")
	namespace = "frnj6sfkc1ep"
)


// the main configuration tasks are runnning here
// If you want to clean the compartment or the tenany then execute 
// utils.DeleteUpload(utils.ListUploads())
// utils.DeleteLogAnalyticsLogGroup(utils.ListLogAnalyticsLogGroups(ocid, ocid_name))	
// utils.DeleteLogAnalyticsLogGroup(utils.ListLogAnalyticsLogGroups(ocid, ocid_name))	
// loc_list_Entity_Id, _ , _:= utils.ListLogAnalyticsEntities(namespace, ocid, ocid_name)  
// utils.DeleteLogAnalyticsEntity(namespace, loc_list_Entity_Id)
// utils.DeleteManagementSavedSearch(utils.ListManagementSavedSearches(ocid, ocid_name))	
// 
// If you want to recreate all the configurations :
// 1/ update the entities.json
// Execute 
// utils.DeleteUpload(utils.ListUploads())
// utils.DeleteLogAnalyticsLogGroup(utils.ListLogAnalyticsLogGroups(ocid, ocid_name))	
// utils.DeleteLogAnalyticsLogGroup(utils.ListLogAnalyticsLogGroups(ocid, ocid_name))	
// loc_list_Entity_Id, _ , _:= utils.ListLogAnalyticsEntities(namespace, ocid, ocid_name)  
// utils.DeleteLogAnalyticsEntity(namespace, loc_list_Entity_Id)
// utils.DeleteManagementSavedSearch(utils.ListManagementSavedSearches(ocid, ocid_name))
// utils.CreateLogAnalyticsLogGroup(namespace,ocid, ocid_name)
// utils.CreateLogAnalyticsEntity(namespace, ocid, ocid_name)
// UploadLogFile(namespace, ocid, utils.ListLogAnalyticsLogGroups(ocid, ocid_name)[0], ocid_name)
// The below setting are recreating a ready to use LogAn series of compartments with LogGroup, Entitties, Files
// The Logan Infra is setup with a tf Scrip ( see ref in github)
//


func process_ocids(namespace string, ocid string, ocid_name string, wg *sync.WaitGroup) {

	utils.DeleteLogAnalyticsLogGroup(utils.ListLogAnalyticsLogGroups(ocid, ocid_name))	
	loc_list_Entity_Id, _ , _:= utils.ListLogAnalyticsEntities(namespace, ocid, ocid_name)  
	utils.DeleteLogAnalyticsEntity(namespace, loc_list_Entity_Id)
    utils.DeleteManagementSavedSearch(utils.ListManagementSavedSearches(ocid, ocid_name))	
//	
 	utils.CreateLogAnalyticsLogGroup(namespace,ocid, ocid_name)
 	utils.CreateLogAnalyticsEntity(namespace, ocid, ocid_name)
//
 	UploadLogFile(namespace, ocid, utils.ListLogAnalyticsLogGroups(ocid, ocid_name)[0], ocid_name)

	wg.Done()
	return
}

/// Eugene Simos
/// main entry point to the programm
/// the ocid of the tenancy is not needed now 
/// provide only the ocid to the second argument of ListCompartments to get a list of compartments to work with
/// The utility can launch several parallel tasks to execute different jobs from the process_ocids function
/// The same task can be executed in a serial way ( commented code below)
/// The entites / files to be uploaded are configured in the entities.json fileName
/// For every type of file to upload an entry has to be created in this files 
/// If at the time to use there is not associated entity with the log file in LogAn the last parameter should be false
/// The upload in paraller way is extremely efficied 500of Mb of day are loaded in 14 compartments in less that 1 minute from a VBox
/// this code will be improved ...
/// You need to create an config file for the oci SDK ( instruction are in OCI site) an example is given below
/// 
/// [DBSECN]
/// region=eu-frankfurt-1
/// tenancy=XXX
/// user=XXX
/// fingerprint=XXX
/// compartment-id=XXX
/// key_file=XXX
/// install go, then run => go run logan.go
/// to make this programm executable 
/// cd utils go build
/// cd logan go install then execute logan
///


func main() {
	ocid, _ := e1.TenancyOCID()
	var list_ocid []string
	var list_ocid_name []string

	list_ocid, list_ocid_name = utils.ListCompartments(ocid, "ocid1.compartment.oc1..aaaaaaaalatb5qnxqrqh7c3fnuj5q7k4mndh3zw7ctq3hkjdwljl2nlbojga")
	
	utils.DeleteUpload(utils.ListUploads())

	wg := &sync.WaitGroup{}
	wg.Add(len(list_ocid))

	for i := 0; i < len(list_ocid); i++ {
		go process_ocids(namespace, list_ocid[i], list_ocid_name[i],  wg)
	}

//   	for i := 0; i < len(list_ocid); i++ {
//  	UploadLogFile(namespace, list_ocid[i], utils.ListLogAnalyticsLogGroups(list_ocid[i], list_ocid_name[i])[0], list_ocid_name[i])
//   	}
 		wg.Wait()
}
