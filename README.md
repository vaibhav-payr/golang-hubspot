[![Build Status](https://travis-ci.org/bold-commerce/go-hubspot.png)](https://travis-ci.org/bold-commerce/go-hubspot)

# go-hubspot
[Go](https://golang.org/) client for [HubSpot](https://app.hubspot.com)

*Note: This currently does not implement all HubSpot API endpoints, however pull requests are welcome*

## Install
```
go get github.com/bold-commerce/go-hubspot
```

## Unit Tests
To run the unit tests, install [ginkgo](https://onsi.github.io/ginkgo) and [gomega](https://onsi.github.io/gomega/) and run:

```
ginkgo -r
```

## Send Email Usage
#### Remove Custom or Contact properties if not required 
```go
package main

import (
	"log"

	"github.com/vaibhav-walke/golang-hubspot"
)

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
```
## Create or Update Usage

```go
package main

import (
	"log"

	"github.com/vaibhav-walke/golang-hubspot"
)

func main() {
 client := hubspot.NewClient("https://api.hubapi.com", "api-key")
 
 	// send single email
 	emailAdd := "vaibhav.walke181@gmail.com"
 	contactRequest := hubspot.CreateOrUpdateContactRequest{
 		Properties: []hubspot.ContactField{
 			{Name: "Country", Value: "Norway"},
 
 		},
 	}
 
 	_, err:= client.CreateOrUpdateContact(contactRequest, emailAdd)
 
 	if err != nil {
 		log.Fatalf("hubspot error: %s", err.Error())
 	}
 }
```

