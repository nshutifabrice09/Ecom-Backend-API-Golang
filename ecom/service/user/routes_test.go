package user

import (
	"bytes"
	"ecom/types"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	//first test
	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "Tanto",
				LastName: "Boom",
				Email: "tanto@gmail.com",
				Password: "12345",
				Phone: "+250780484427".
				Address: "kg102st",

			}
			marshalled, _ := json.Marshal(payload)
			
			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
			if err != nil{
				t.Fail(err)
			}

			rr := httptest.NewRecorder()
			router := mux.Router()

			router.HandleFunc("/register", handler.handleRegister())
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusBadRequest {
				t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
			}
		})

}

type mockUserStore struct {
	// Add any fields you need for your tests
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById (id int) (*types.User, error){
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
