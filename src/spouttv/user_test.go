package spouttv

import(
  "testing"
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