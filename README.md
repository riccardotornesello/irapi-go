# iRacing API - Go Client

This is a Go client for the iRacing API.

It is capable of connecting to most of the API endpoints and parsing the responses into Go structs.

## Usage

```go
package main

import (
	"github.com/riccardotornesello/irapi-go"
	"github.com/riccardotornesello/irapi-go/api/member"
)

func main() {
	irClient, err := irapi.NewIRacingApiClient("myemail", "mypassword")
	if err != nil {
		panic(err)
	}

	// Get the user info
	params := member.MemberGetParams{
		CustIds: []int{911231},
	}
	userInfo, err := irClient.Member.GetMember(params)
	if err != nil {
		panic(err)
	}

	// Print the user info
	for _, member := range userInfo.Members {
		println(member.DisplayName)
	}
}
```

## Next steps

This project is in its early stages and will grow a lot in the coming period.

Here are the highest priority tasks:

- better document the use of the library and the generation tool (tools folder), perhaps by having the tool generate the documentation automatically
- improve the API generation tool to make it clearer to use and modify
- add output in OpenAPI format to the tool
- add missing endpoints and for those that are present add unknown field types, using different parameters in example calls
- add comments to parameters in go structs
- adjust tests for endpoints that require parameters since they don't currently receive them
- if skip s3 is true, the get function should skip the s3 call
- some parameters are lists, check if they're passed correctly in the sdk and in the tools
- add the structs for the csv files
