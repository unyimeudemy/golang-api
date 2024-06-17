package user

import (
	"Ecom/types"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// This test focuses on testing the the handler and thus the database repository
// is mocked
func TestUserServiceHandler(t *testing.T){

	// Create a mock user store to simulate database interactions. Implementation of
	// of mock return type are done below
	userStore := &mockUserStore{}

	// Initialize the needed handler or controller with the mock user store
	handler := NewHandler(userStore)

	t.Run(
		"should fail if the user payload is invalid",
		func(t *testing.T){

			// create a test payload
			payload := types.RegisterUserPayload{
				Firstname: "unyime",
				Lastname: "udoh",
				Email: "INVALID_EMAIL",
				Password: "123456",
			}
		
			//json.Marshal converts the payload into its JSON representation 
			marshalled, _ := json.Marshal(payload)
		
			// define what you want your test request to look like using the payload
			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			//	rr = response recorder ( create a variable that will store the
			//  response from the request )
			rr := httptest.NewRecorder()

			// create a new router for the testing purpose
			router := mux.NewRouter()

			// specify which controller or handler you want to test and its entry point (route)
			router.HandleFunc("/register", handler.handleRegister)

			// Make actual test http call and store the response in rr
			router.ServeHTTP(rr, req)

			// assert that the status code from the test response is what you expected
			if rr.Code != http.StatusBadRequest {
				t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
			}
		},
	)

	t.Run(
		"Should correctly register the user",
		func (t *testing.T){
						// create a test payload
						payload := types.RegisterUserPayload{
							Firstname: "unyime",
							Lastname: "udoh",
							Email: "a@gmail.com",
							Password: "123456",
						}
					
						//json.Marshal converts the payload into its JSON representation 
						marshalled, _ := json.Marshal(payload)
					
						// define what you want your test request to look like using the payload
						req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
						if err != nil {
							t.Fatal(err)
						}
			
						//	rr = response recorder ( create a variable that will store the
						//  response from the request )
						rr := httptest.NewRecorder()
			
						// create a new router for the testing purpose
						router := mux.NewRouter()
			
						// specify which controller or handler you want to test and its entry point (route)
						router.HandleFunc("/register", handler.handleRegister)
			
						// Make actual test http call and store the response in rr
						router.ServeHTTP(rr, req)
			
						// assert that the status code from the test response is what you expected
						if rr.Code != http.StatusCreated {
							t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
						}
		},
	)
}


// Mocked version of the repository
type mockUserStore struct {}

func (m *mockUserStore)GetUserByEmail(email string)(*types.User, error){
	return nil, fmt.Errorf("user not found")

}

func (m *mockUserStore)GetUserByID(id int)(*types.User, error){
	return nil, nil
}


func (m *mockUserStore)CreateUser(types.User) error {
	return nil
}

