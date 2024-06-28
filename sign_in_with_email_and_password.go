package firebase_auth

import (
	"fmt"
	"net/http"
)

type SignInWithEmailPasswordRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type SignInWithEmailPasswordResponse struct {
	IDToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
	Registered   bool   `json:"registered"`
}

// SignUpWithEmailPassword creates a new email and password user by issuing an HTTP POST request to the Auth
// `signupNewUser` endpoint.
//
//	Reference: https://firebase.google.com/docs/reference/rest/auth#section-create-email-password
func (c *client) SignInWithEmailAndPassword(req *SignInWithEmailPasswordRequest) (*SignInWithEmailPasswordResponse, *ErrorResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", c.apiKey)
	return request[SignInWithEmailPasswordRequest, SignInWithEmailPasswordResponse](c.httpClient, http.MethodPost, url, req)
}
