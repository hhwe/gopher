package main

import (
	"net/http"
	"fmt"
	"log"
)

func init() {
	http.HandleFunc("/api/signin", signInHandle)
}

func main() {


	http.ListenAndServe(":8080", nil)
}

type SignInResponse struct {
	UserID string `json:"user_id"`
	SessionToken string `json:"session_token"`
}

func signInHandle(w http.ResponseWriter, r *http.Request) {
	//ctx := appengine.New
	token := r.Header.Get("Authorization")
	userID, err := verifyToken(ctx, token)
	if err != nil {
		log.Printf(ctx, "%v", err)

	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Sign in ")
}
