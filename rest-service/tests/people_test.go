package tests

import (
	"net/http"
	"testing"
)

func TestPeople(t *testing.T) {
	testCases := []APITestCase{
		{
			Name:         "test_GetPeople_All",
			Method:       "GET",
			URL:          "/people",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `[{"id":"81eb745b-3aae-400b-959f-748fcafafd81","first_name":"John","last_name":"Doe","phone_number":"+1 (800) 555-1212"},{"id":"5b81b629-9026-450d-8e46-da4f8c7bd513","first_name":"Jane","last_name":"Doe","phone_number":"+1 (800) 555-1313"},{"id":"df12ce76-767b-4bf0-bccb-816745df9e70","first_name":"Brian","last_name":"Smith","phone_number":"+44 7700 900077"},{"id":"135af595-aa86-4bb5-a8f7-df17e6148e63","first_name":"John","last_name":"Doe","phone_number":"+1 (800) 555-1414"},{"id":"000ebe58-b659-422b-ab48-a0d0d40bd8f9","first_name":"Jenny","last_name":"Smith","phone_number":"+44 7700 900077"}]`,
		},
		{
			Name:         "test_GetPeople_ByName_DataExist",
			Method:       "GET",
			URL:          "/people?first_name=Brian&last_name=Smith",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `[{"id":"df12ce76-767b-4bf0-bccb-816745df9e70","first_name":"Brian","last_name":"Smith","phone_number":"+44 7700 900077"}]`,
		},
		{
			Name:         "test_GetPeople_ByName_DataEmpty",
			Method:       "GET",
			URL:          "/people?first_name=&last_name=",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `[]`,
		},
		{
			Name:         "test_GetPeople_ByPhoneNumber_DataExist",
			Method:       "GET",
			URL:          "/people?phone_number=%2B1+%28800%29+555-1313",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `[{"id":"5b81b629-9026-450d-8e46-da4f8c7bd513","first_name":"Jane","last_name":"Doe","phone_number":"+1 (800) 555-1313"}]`,
		},
		{
			Name:         "test_GetPeople_ByPhoneNumber_DataEmpty",
			Method:       "GET",
			URL:          "/people?phone_number=112233",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `[]`,
		},
		{
			Name:         "test_GetPeopleByID_DataExist",
			Method:       "GET",
			URL:          "/people/5b81b629-9026-450d-8e46-da4f8c7bd513",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `{"id":"5b81b629-9026-450d-8e46-da4f8c7bd513","first_name":"Jane","last_name":"Doe","phone_number":"+1 (800) 555-1313"}`,
		},
		{
			Name:         "test_GetPeopleByID_404NotFound",
			Method:       "GET",
			URL:          "/people/112233",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusNotFound,
			WantResponse: `{"message":"data not found"}`,
		},
	}

	router := SetUpRouter()
	for _, tc := range testCases {
		Endpoint(t, router, tc)
	}
}
