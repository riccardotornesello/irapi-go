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

- update documentation with new version
- document the scraping tool and update it to work with the new oauth flow
- generate more endpoints with the scraping tool
- add tests to endpoints and client generation
- parse csv endpoints and add types
- automatically refresh tokens
