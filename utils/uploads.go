package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v35/loganalytics"
	"io/ioutil"
	"os"
)



func ListUploadFiles(namespace string, namefile string, UploadReference string) loganalytics.UploadFileCollection {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	// Create a request and dependent object(s).
	req := loganalytics.ListUploadFilesRequest{
	    SearchStr: &namefile,
		NamespaceName:  &namespace,
		UploadReference: &UploadReference,
       }

	// Send the request using the service client
	resp, err := client.ListUploadFiles(context.Background(), req)
	if err != nil {
	fmt.Println("Error ListUploadFiles:", err)
    }
    return  resp.UploadFileCollection 
}

func ListUploads(namespace string, namecontains string) loganalytics.UploadCollection {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	var ListUploadsRequest loganalytics.ListUploadsRequest

	ListUploadsRequest = loganalytics.ListUploadsRequest{
			NameContains: &namecontains,
     		NamespaceName: &namespace,
	}
	
	fmt.Printf("%+v\n",ListUploadsRequest )
	
	resp, _ := client.ListUploads(context.Background(), ListUploadsRequest)
	//	var loc_array []string
	//	for _, v := range resp.Items {
	//		fmt.Printf("%s %s \n", *v.Reference, *v.Name)
	//		loc_array = append(loc_array, *v.Reference)
	//	}
	return resp.UploadCollection
}

// Deletes an Upload by its reference
func DeleteUpload(namespace string, UploadCollection loganalytics.UploadCollection) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)
	for i := 0; i < len(UploadCollection.Items); i++ {
		DeleteUploadRequest := loganalytics.DeleteUploadRequest{
			NamespaceName:   &namespace,
			UploadReference: UploadCollection.Items[i].Reference ,
		}
		resp, err := client.DeleteUpload(context.Background(), DeleteUploadRequest)
		if err != nil {
			fmt.Println("Error DeleteUpload:", err)
			//return
		} else {
			fmt.Printf("%d %d \n", *resp.OpcDeletedLogfileCount, *resp.OpcDeletedLogCount)
		}
	}
}

func ValidateFile(namespace string, ocid string, OpcMetaLoggrpid string, ocid_name string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)
	type Entity struct {
		Directory           string `json:"directory"`
		Name                string `json:"name"`
		Type                string `json:"type"`
		Upload_Name         string `json:"upload_name"`
		LogSourceName       string `json:"logSourceName"`
		Invalidatecache     bool   `json:"invalidatecache"`
		Associate_with_type bool   `json:"associate_with_type"`
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
	if err := json.Unmarshal(byteValue, &entities); err != nil {
		panic(err)
	}

	//fmt.Printf("ListLogAnalyticsEntities=>%s %s\n", ocid  ,ocid_name)
	// _, _, _ := ListLogAnalyticsEntities(namespace,ocid,ocid_name)

	// Map[name] => id

	//    fmt.Printf("Entity_ocid=>%s\n", Entity_ocid )

	for i := 0; i < len(entities.Entities); i++ {
		files, err := ioutil.ReadDir(entities.Entities[i].Directory)
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range files {
			if f.Name() != "config.properties" {
				fmt.Println(f.Name())
				var fileName string = f.Name()
				var full_path string = entities.Entities[i].Directory + "/" + fileName
				ValidateFileRequest := loganalytics.ValidateFileRequest{
					NamespaceName:  &namespace,
					ObjectLocation: &full_path,
					Filename:       &fileName,
				}
				_, err := client.ValidateFile(context.Background(), ValidateFileRequest)
				if err != nil {
					fmt.Printf("Validation File Error on file %s located to directory %s %s\n", fileName, entities.Entities[i].Directory, err)
				}
			}
		}

	}
}

func UploadLogFile(namespace string, ocid string, OpcMetaLoggrpid string, ocid_name string) {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	type Entity struct {
		Directory           string `json:"directory"`
		Name                string `json:"name"`
		Type                string `json:"type"`
		Upload_Name         string `json:"upload_name"`
		LogSourceName       string `json:"logSourceName"`
		Invalidatecache     bool   `json:"invalidatecache"`
		Associate_with_type bool   `json:"associate_with_type"`
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
	if err := json.Unmarshal(byteValue, &entities); err != nil {
		panic(err)
	}

	//fmt.Printf("ListLogAnalyticsEntities=>%s %s\n", ocid  ,ocid_name)
	_, _, Map := ListLogAnalyticsEntities(namespace, ocid, ocid_name)

	// Map[name] => id

	//    fmt.Printf("Entity_ocid=>%s\n", Entity_ocid )

	for i := 0; i < len(entities.Entities); i++ {
		files, err := ioutil.ReadDir(entities.Entities[i].Directory)
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range files {
			if f.Name() != "config.properties" {
				fmt.Println(f.Name())

				var fileName string = f.Name()

				content, err := os.Open(entities.Entities[i].Directory + "/" + f.Name())
				if err != nil {
					fmt.Println(err)
				}

				fmt.Printf("%s %s %s %s %s %s\n", namespace, entities.Entities[i].Upload_Name, entities.Entities[i].LogSourceName, fileName, OpcMetaLoggrpid, entities.Entities[i].Type)
				var UploadName = entities.Entities[i].Upload_Name + "-" + ocid_name
				var Entity_Key = entities.Entities[i].Name + "-" + ocid_name
				var Entity_Ocid string
				var UploadLogFileRequest loganalytics.UploadLogFileRequest

				if entities.Entities[i].Associate_with_type == true {
					Entity_Ocid = Map[Entity_Key]

					UploadLogFileRequest = loganalytics.UploadLogFileRequest{
						EntityId: &Entity_Ocid,
					}

				}

				if entities.Entities[i].LogSourceName == "OCI VCN Flow Logs" && entities.Entities[i].Type == "oci_vcn" {

				}

				UploadLogFileRequest = loganalytics.UploadLogFileRequest{
					NamespaceName:     &namespace,
					UploadName:        &UploadName,
					LogSourceName:     &entities.Entities[i].LogSourceName,
					Filename:          &fileName,
					OpcMetaLoggrpid:   &OpcMetaLoggrpid,
					InvalidateCache:   &entities.Entities[i].Invalidatecache,
					UploadLogFileBody: ioutil.NopCloser(content),
				}

				_, err = client.UploadLogFile(context.Background(), UploadLogFileRequest)
				if err != nil {
					fmt.Println("Error UploadLogFile:", err)
					return
				}

			}
		}
	}
}
