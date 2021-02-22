package utils

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v35/common"
	_ "github.com/oracle/oci-go-sdk/v35/identity"
	_ "github.com/oracle/oci-go-sdk/v35/loganalytics"
	"github.com/oracle/oci-go-sdk/v35/managementdashboard"
)

var (
	e1 = common.CustomProfileConfigProvider("/home/oracle/.oci/config", "DBSECN")
)

func DeleteManagementSavedSearch(ManagementSavedSearchId []string) {
	client, _ := managementdashboard.NewDashxApisClientWithConfigurationProvider(e1)

	for i := 0; i < len(ManagementSavedSearchId); i++ {
		DeleteManagementSavedSearchRequest := managementdashboard.DeleteManagementSavedSearchRequest{
			ManagementSavedSearchId: &ManagementSavedSearchId[i],
		}
		_, err := client.DeleteManagementSavedSearch(context.Background(), DeleteManagementSavedSearchRequest)
		if err != nil {
			fmt.Println("Error DeleteLogAnalyticsEntity:", err)
			//return
		}
	}
}

func ListManagementSavedSearches(ocid string, ocid_name string) []string {
	client, _ := managementdashboard.NewDashxApisClientWithConfigurationProvider(e1)
	ListManagementSavedSearchesRequest := managementdashboard.ListManagementSavedSearchesRequest{
		CompartmentId: &ocid,
	}
	var loc_array []string
	resp, _ := client.ListManagementSavedSearches(context.Background(), ListManagementSavedSearchesRequest)
	for _, v := range resp.Items {
		loc_array = append(loc_array, *v.Id)
	}
	return loc_array

}
