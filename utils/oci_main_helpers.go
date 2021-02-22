package utils

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v35/common"
	"github.com/oracle/oci-go-sdk/v35/identity"
	"github.com/oracle/oci-go-sdk/v35/loganalytics"
	"logan/helpers"
)

func ListCompartments(ocid string, ocid_parent string) ([]string, []string) {
	c, _ := identity.NewIdentityClientWithConfigurationProvider(e1)

	var CompartmentIdInSubtree = false

	ListCompartmentsRequest := identity.ListCompartmentsRequest{
		CompartmentId:          &ocid_parent,
		AccessLevel:            "ANY",
		CompartmentIdInSubtree: &CompartmentIdInSubtree,
		LifecycleState:         "ACTIVE",
	}

	list, _ := c.ListCompartments(context.Background(), ListCompartmentsRequest)
	fmt.Printf("Nb of ListCompartments : %d\n", len(list.Items))

	var loc_ocid_arr []string
	var loc_ocid_name []string
	for _, v := range list.Items {
		loc_ocid_arr = append(loc_ocid_arr, *v.Id)
		loc_ocid_name = append(loc_ocid_name, *v.Name)
	}

	return loc_ocid_arr, loc_ocid_name

}

func ExampleListAvailabilityDomains() {
	c, err := identity.NewIdentityClientWithConfigurationProvider(e1)
	helpers.FatalIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helpers.FatalIfError(err)

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	helpers.FatalIfError(err)
	fmt.Printf("Nb of availabilty domains : %d\n", len(r.Items))
	for i, v := range r.Items {
		fmt.Printf("domain[%d]:%s\n", i, *v.Name)
	}

	//log.Printf("list of available domains: %v+", r.Items)
	//fmt.Println("list available domains completed")

	// Output:
	// list available domains completed
}

func GetNamespace() *string {
	client, _ := loganalytics.NewLogAnalyticsClientWithConfigurationProvider(e1)

	var namespace = "frnj6sfkc1ep"
	GetNamespaceRequest := loganalytics.GetNamespaceRequest{
		NamespaceName: &namespace,
	}

	resp, _ := client.GetNamespace(context.Background(), GetNamespaceRequest)

	return resp.Namespace.NamespaceName

}
