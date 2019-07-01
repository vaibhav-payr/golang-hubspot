
//This file contains example main functions for multiple operations
package main

import (
	"log"

	//"github.com/vaibhav-walke/golang-hubspot"
)

//function for send email
func main() {
	client := hubspot.NewClient("https://api.hubapi.com", "api-key")

	// send single email
	emailId := 12345678910
	emailRequest := hubspot.SendEmailRequest{
		EmailID:emailId,
		Message:hubspot.Message{
			To:"vaibhav.walke181@gmail.com",
		},
		CustomProperties: []hubspot.MergeField{
			{Name: "code", Value: "1234"},
		},
		ContactProperties:[]hubspot.MergeField{
		{Name:"keyname", Value:"value"},
		},
	}

	err := client.SingleEmail(emailRequest)

	if err != nil {
		log.Fatalf("hubspot error: %s", err.Error())
	}
}

//function for send or update contact
func main() {
	client := hubspot.NewClient("https://api.hubapi.com", "api-key")

	// send single email
	emailAdd := "vaibhav.walke181@gmail.com"
	emailRequest := hubspot.CreateOrUpdateContactRequest{
		Properties: []hubspot.ContactField{
			{Name: "Country", Value: "Norway"},

		},
	}

	_, err:= client.CreateOrUpdateContact(emailRequest, emailAdd)

	if err != nil {
		log.Fatalf("hubspot error: %s", err.Error())
	}
}

//Function to add contacts in List - WIP
func main() {
	client := hubspot.NewClient("https://api.hubapi.com", "api-key")

	// send single email
	listId:=1
	addListRequest := []string{
		"vaibhav@payr.no","vaibhav@walke181@gmail.com",
	}

	_, err:= client.AddContactsToList(addListRequest, listId)

	if err != nil {
		log.Fatalf("hubspot error: %s", err.Error())
	}
}
