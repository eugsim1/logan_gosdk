//++
// Eugene Simos
// utility to load data /clean up from data a logging analytics OCI compartement or a series of loggan compartments under a "main compartment"

package main

import (
	_ "context"
	_ "encoding/json"
	 "fmt"
	"github.com/oracle/oci-go-sdk/v35/common"
	_ "github.com/oracle/oci-go-sdk/v35/example/helpers"
	_ "github.com/oracle/oci-go-sdk/v35/identity"
	_ "github.com/oracle/oci-go-sdk/v35/loganalytics"
	_ "github.com/oracle/oci-go-sdk/v35/managementdashboard"
	_ "io/ioutil"
	_ "log"
	"logan/utils"
	_ "os"
	"sync"
)

var (
	// use a config file the same that you use for the OCI CLI DBSECN is a particular sub section with the details of the target tenancy
	e1        = common.CustomProfileConfigProvider("/home/oracle/.oci/config", "DBSECN")
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
// loc_list_Entity_Id, _ , _:= utils.ListLogAnalyticsEntities(namespace, ocid, ocid_name)
// utils.DeleteLogAnalyticsEntity(namespace, loc_list_Entity_Id)
// utils.DeleteManagementSavedSearch(utils.ListManagementSavedSearches(ocid, ocid_name))
// utils.CreateLogAnalyticsLogGroup(namespace,ocid, ocid_name)
// utils.CreateLogAnalyticsEntity(namespace, ocid, ocid_name)
// UploadLogFile(namespace, ocid, utils.ListLogAnalyticsLogGroups(ocid, ocid_name)[0], ocid_name)

// The below settings are use to recreate a ready to use LogAn series of compartments with LogGroup, Entitties, Files
// The Logan Infra is setup with a tf Script ( see this repo https://github.com/eugsim1/logging-analytics-tf-infra )
//

func process_ocids(namespace string, ocid string, ocid_name string, wg *sync.WaitGroup) {

	//	utils.DeleteLogAnalyticsLogGroup(utils.ListLogAnalyticsLogGroups(ocid, ocid_name))
	// 	loc_list_Entity_Id, _ , _:= utils.ListLogAnalyticsEntities(namespace, ocid, ocid_name)
	//	utils.DeleteLogAnalyticsEntity(namespace, loc_list_Entity_Id)
	//        utils.DeleteManagementSavedSearch(utils.ListManagementSavedSearches(ocid, ocid_name))
	//
	// 	utils.CreateLogAnalyticsLogGroup(namespace,ocid, ocid_name)
	// 	utils.CreateLogAnalyticsEntity(namespace, ocid, ocid_name)
	//
	// utils.ValidateFile(namespace, ocid, utils.ListLogAnalyticsLogGroups(ocid, ocid_name)[0], ocid_name)
	// 	utils.UploadLogFile(namespace, ocid, utils.ListLogAnalyticsLogGroups(ocid, ocid_name)[0], ocid_name)

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
/// If at the time to run the scripts  there is not associated entity with the log file in LogAn the last parameter should be false
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
///
/// install go, then run => go run logan.go
/// to make this programm executable
/// cd utils go build
/// cd logan go install then execute logan
///

func main() {

	ocid, _ := e1.TenancyOCID()

	var list_ocid []string
	var list_ocid_name []string

	// the second parameter is your "main compartement of Log Analytics
	// if below there are many others the scrip will take care of all of them
	//
	list_ocid, list_ocid_name = utils.ListCompartments(ocid, "ocid1.compartment.oc1..aaaaaaaalatb5qnxqrqh7c3fnuj5q7k4mndh3zw7ctq3hkjdwljl2nlbojga")
	
	var UploadCollection = utils.ListUploads(namespace,"")
	fmt.Printf("%v",UploadCollection)

	//utils.DeleteUpload(namespace,utils.ListUploads(namespace))

	// go routine to execute in parallel the config tasks
	// this makes the whole script to run several conigurations tasks for every compartment
	// if you dont like this comment the code and execute all the functions on the commented loop below

	wg := &sync.WaitGroup{}
	wg.Add(len(list_ocid))

	// for every compartment execute a batch of configs
	//
	for i := 0; i < len(list_ocid); i++ {
		go process_ocids(namespace, list_ocid[i], list_ocid_name[i], wg)
	}

	//   	for i := 0; i < len(list_ocid); i++ {
	//  	UploadLogFile(namespace, list_ocid[i], utils.ListLogAnalyticsLogGroups(list_ocid[i], list_ocid_name[i])[0], list_ocid_name[i])
	//   	}
	wg.Wait()
}
