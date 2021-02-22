package utils

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/oracle/oci-go-sdk/v35/common"
	_ "github.com/oracle/oci-go-sdk/v35/identity"
	"github.com/oracle/oci-go-sdk/v35/loganalytics"
	_ "github.com/oracle/oci-go-sdk/v35/loganalytics"
	_ "github.com/oracle/oci-go-sdk/v35/managementdashboard"
	"io/ioutil"
	"os"
)

func ListLogAnalyticsEntities(namespace string, ocid string, ocid_name string) (Id []string, Name []string, M map[string]string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	ListLogAnalyticsEntitiesRequest := loganalytics.ListLogAnalyticsEntitiesRequest{
		NamespaceName:  &namespace,
		CompartmentId:  &ocid,
		LifecycleState: "ACTIVE",
	}

	resp, _ := client.ListLogAnalyticsEntities(context.Background(), ListLogAnalyticsEntitiesRequest)
	var loc_array []string
	var loc_array_name []string
	var m = make(map[string]string)
	for _, v := range resp.Items {
		loc_array = append(loc_array, *v.Id)
		loc_array_name = append(loc_array_name, *v.Name)
		m[*v.Name] = *v.Id
	}
	return loc_array, loc_array_name, m
}

func CreateLogAnalyticsEntity(namespace string, ocid string, ocid_name string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	type Entity struct {
		Directory       string `json:"directory"`
		Name            string `json:"name"`
		Type            string `json:"type"`
		Upload_Name     string `json:"upload_name"`
		LogSourceName   string `json:"logSourceName"`
		Invalidatecache bool   `json:"invalidatecache"`
	}

	type Entities struct {
		Entities []Entity `json:"entities"`
	}

	jsonFile, err := os.Open("entities.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var entities Entities

	json.Unmarshal(byteValue, &entities)

	//fmt.Printf("%d\n", len(entities.Entities))

	for i := 0; i < len(entities.Entities); i++ {
		Entity_name := entities.Entities[i].Name + "-" + ocid_name
		fmt.Printf("%s\n", Entity_name)
		CreateLogAnalyticsEntityRequest := loganalytics.CreateLogAnalyticsEntityRequest{
			NamespaceName: &namespace,
			CreateLogAnalyticsEntityDetails: loganalytics.CreateLogAnalyticsEntityDetails{
				Name:           &Entity_name,
				CompartmentId:  &ocid,
				EntityTypeName: &entities.Entities[i].Type,
			},
		}

		_, err := client.CreateLogAnalyticsEntity(context.Background(), CreateLogAnalyticsEntityRequest)
		if err != nil {
			fmt.Println("Error CreateLogAnalyticsEntity:", err)
			//return
		} else {
			//fmt.Printf("%s \n", *resp.RawResponse)
		}
	}

}

func DeleteLogAnalyticsEntity(namespace string, LogAnalyticsEntityId []string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	for i := 0; i < len(LogAnalyticsEntityId); i++ {
		DeleteLogAnalyticsEntityRequest := loganalytics.DeleteLogAnalyticsEntityRequest{
			NamespaceName:        &namespace,
			LogAnalyticsEntityId: &LogAnalyticsEntityId[i],
		}

		_, err := client.DeleteLogAnalyticsEntity(context.Background(), DeleteLogAnalyticsEntityRequest)
		if err != nil {
			fmt.Println("Error DeleteLogAnalyticsEntity:", err)
			//return
		} else {
			//fmt.Printf("%s \n", *resp.RawResponse)
		}
	}
}
