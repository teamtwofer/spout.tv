package spouttv

import(
  "testing"
  "fmt"
)

var validEmails = []string {
  "ben@benbayard.com",
  "b.j.bayard@gmail.com",
  "b.j.bayard+potato@gmail.com",
  "b.j.ba_yard+pota_to@gma_il.com",
}

var invalidEmails = []string {
  "ben",
  "ben@ben",
  "ben$ben@ben.com",
}

func TestValidateEmail(t *testing.T) {
  for _, email := range validEmails {
    v, err := ValidateEmail(email)
    if v != true && err == nil {
      t.Error(
        "For", email, 
        "expected", true,
        "got", false,
      )
    }
  }
  for _, email := range invalidEmails {
    v, err := ValidateEmail(email)
    if v != false && err == nil {
      t.Error(
        "For", email, 
        "expected", false,
        "got", true,
      )
    }
  }
}

func TestCreateUser(t *testing.T) {
  nu := &NewUser{
    Name:     "Ben Bayard",
    Email:    "b.j.bayard@gmail.com",
    Password: "potato",
  }

  user, err := CreateUser(*nu)

  if err == nil && user.Email != "" && user.Name != "" && user.EncryptedPassword != nil {
    fmt.Println(user.LastSignInAt)
    fmt.Println(user.CreatedAt)
    fmt.Println(user.UpdatedAt)
    if (user.LastSignInAt == 0 || user.CreatedAt == 0 || user.UpdatedAt == 0) {
      t.Error("Signin, Createdat or updatedat were not set")  
    }
  } else {  
    t.Error("User, name or password was not set")  
  }
}

func TestCreateUser_2(t *testing.T) {
  // invalid email!
  nu := &NewUser{
    Name:     "Ben Bayard",
    Email:    "b.j.bayardgmail.com",
    Password: "potato",
  }

  user, err := CreateUser(*nu)

  if err == nil {
    t.Error(
      "For", user.Email, 
      "expected error to be nil",
      "but it wasn't dawg",
    )
  }
}

func BenchmarkEncryptPassword(b *testing.B) {
  pass := []byte("password")
  for i := 0; i < b.N; i++ {
    ep, _ := EncryptPassword(pass)
    if ep != nil {

    }
  }
}

func BenchmarkRandString(b *testing.B) {
  for i := 0; i < b.N; i++ {
    p := RandString(30)
  }
}