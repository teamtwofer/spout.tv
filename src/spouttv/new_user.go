package spouttv

type NewUser struct {
  Name         string `json:"name"`
  Email        string `json:"email"`
  Password     string `json:"password"`
  Username     string `json:"username"`
  StayLoggedIn string `json:"stay_logged_in"`
}

type LoginUser struct {
  Email        string `json:"email"`
  Username     string `json:"username"`
  Password     string `json:"password"`
  StayLoggedIn string `json:"stay_logged_in"`
}