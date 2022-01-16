package source

import "testing"

var mockJSON = `
	{
		"common": {
			"lineNumber": "Номер строки"
		},
		"person" : {
			"lastName": "Фамилия",
			"firstName": "Имя",
			"contact": {
				"phone": "88005553535",
				"countryCode": "030",
			}
		}

	}
`

var mockTOML = `
[common]
  lineNumber = "Номер строки"
[person]
  lastName = "Фамилия"
  firstName = "Имя"
	[contact]
		phone = "88005553535"
		countryCode = "030"
`

var mockYAML = `

`

// TestGet - func to test value from dict
func TestGet(t *testing.T) {

}

// TestGetList - func to test values from dict
func TestGetList(t *testing.T) {

}
