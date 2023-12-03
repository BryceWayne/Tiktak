package models

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"` // Store as a hash
    Profile  string `json:"profile"`  // URL to the profile picture
    Bio      string `json:"bio"`
}
