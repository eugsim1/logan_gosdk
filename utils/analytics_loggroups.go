package utils

import (
    "fmt"
	"context"	
	_"github.com/oracle/oci-go-sdk/v35/common"	
	_"github.com/oracle/oci-go-sdk/v35/identity"
	_"github.com/oracle/oci-go-sdk/v35/loganalytics"
	_"github.com/oracle/oci-go-sdk/v35/managementdashboard"
	"github.com/oracle/oci-go-sdk/v35/loganalytics"	
)




func ListLogAnalyticsLogGroups(ocid string, ocid_name string) []string {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)
	var namespace = "frnj6sfkc1ep"

	ListLogAnalyticsLogGroupsRequest := loganalytics.ListLogAnalyticsLogGroupsRequest{
		NamespaceName: &namespace,
		CompartmentId: &ocid,
	}
	var loc_array []string
	resp, _ := client.ListLogAnalyticsLogGroups(context.Background(), ListLogAnalyticsLogGroupsRequest)
	for _, v := range resp.Items {
		loc_array = append(loc_array, *v.Id)
	}
	return loc_array
}

func CreateLogAnalyticsLogGroup(namespace  string,ocid string, ocid_name string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)


		var DisplayName string = ocid_name + "LogGroup"
		var Description string = "test description"
		LogAnalyticsLogGroupDetails := loganalytics.CreateLogAnalyticsLogGroupRequest{
			NamespaceName: &namespace,
			CreateLogAnalyticsLogGroupDetails: loganalytics.CreateLogAnalyticsLogGroupDetails{
				DisplayName:   &DisplayName,
				CompartmentId: &ocid,
				Description:   &Description,
				FreeformTags: map[string]string{
					"Project":     "log_analytics",
					"Role":        "log_analytics for HOL ",
					"Comment":     "log_analytics setup for HOL ",
					"Version":     "0.0.0.0",
					"Responsible": "Eugene Simos",
					"agent":       "oci sdk go"},
			},
		}


	_, err := client.CreateLogAnalyticsLogGroup(context.Background(), LogAnalyticsLogGroupDetails)
	if err != nil {
		fmt.Println("Error CreateLogAnalyticsLogGroupDetails:", err)
	}

}


func DeleteLogAnalyticsLogGroup(LogAnalyticsLogGroupId []string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)
	var namespace = "frnj6sfkc1ep"
	for i := 0; i < len(LogAnalyticsLogGroupId); i++ {
		DeleteLogAnalyticsLogGroupRequest := loganalytics.DeleteLogAnalyticsLogGroupRequest{
			NamespaceName:          &namespace,
			LogAnalyticsLogGroupId: &LogAnalyticsLogGroupId[i],
		}

		_, err := client.DeleteLogAnalyticsLogGroup(context.Background(), DeleteLogAnalyticsLogGroupRequest)
		if err != nil {
			fmt.Println("Error DeleteLogAnalyticsLogGroup:", err)
		}
	}
}
