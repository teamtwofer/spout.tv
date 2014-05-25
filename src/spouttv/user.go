package spouttv

import (
  // "encoding/json"
  // "fmt"
  // "io/ioutil"
  // "net/http"
  "regexp"
  "code.google.com/p/go.crypto/bcrypt"
)

type User struct {
  Name              string `json:"name"`
  Email             string `json:"email"`
  EncryptedPassword []byte
  LastSignInAt      int64  `json:"last_sign_in_at"`
  CreatedAt         int64  `json:"created_at"`
  UpdatedAt         int64  `json:"updated_at"`
}

// type UserError struct {
//   Error string ``
// }

func NewUser(name string, email string, password []byte) (user *User) {
  user = &User{}
  user.Name = name
  validEmail, err := ValidateEmail(email)
  if err != nil {
    // TODO: Create a UserError
  } 
  if validEmail == true {
    user.Email = email
  }
  encryptedPassword, err := encryptPassword([]byte(password))
  if err != nil {
    // TODO: Create a UserError
  } else {
    user.EncryptedPassword = encryptedPassword
  }
  return
}

func ValidateEmail(email string) (status bool, err error) {
  status, err = regexp.MatchString(`^([\w\+\.]+@\w+\.\w+)$`, email)
  return 
}

// you use clear because it removed the password in plain text from memory. 
func clear(b []byte) {
  for i := 0; i < len(b); i++ {
    b[i] = 0;
  }
}

func encryptPassword(password []byte) ([]byte, error) {
  defer clear(password)
  return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}
