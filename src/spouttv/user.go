package spouttv

// new User HTTP Benchmark
// for i in {1..10};curl -X POST "http://localhost:3000/user" -d $'{"name": "Ben Bayard","email": "bjbayard@gmail.com","password": "potato"}' -m 30 -s -w "%{time_total}\n" -o /dev/null

import (
  "fmt"
  "time" 
  "regexp"
  "code.google.com/p/go.crypto/bcrypt"
  "crypto/rand"
  "math/big"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type User struct {
  Id                bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Name              string        `json:"name"`
  Email             string        `json:"email"`
  Username          string        `json:"email"`
  EncryptedPassword []byte        `json:"password"`
  LastSignInAt      int64         `json:"last_signin_at"`
  CreatedAt         int64         `json:"created_at"`
  UpdatedAt         int64         `json:"updated_at"`
  LoginToken        string        `json:"token"`
}

// you use clear because it removed the password in plain text from memory. 
func clear(b []byte) {
  for i := 0; i < len(b); i++ {
    b[i] = 0;
  }
}

func FindUserById(id string) User {
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("spouttv").C("users")

  result := User{}
  
  c.FindId(bson.ObjectIdHex(id)).One(&result)

  fmt.Printf("%#v \n", result)


  if err != nil{
    return User{}
  } else {
    return result
  }
}

func CreateUser(newUser NewUser) (user *User, userErr error) {
  defer clear([]byte(newUser.Password))
  user = &User{}
  user.Name = newUser.Name
  user.Username = newUser.Username
  validEmail, err := ValidateEmail(newUser.Email)
  if err != nil {
    // TODO: Create a UserError
  } 
  if validEmail == true {
    user.Email = newUser.Email
    encryptedPassword, err := EncryptPassword([]byte(newUser.Password))
    if err != nil {
      // TODO: Create a UserError
    } else {
      user.EncryptedPassword = encryptedPassword
    }

    if err == nil {
      user.LastSignInAt = time.Now().Unix()
      user.CreatedAt    = time.Now().Unix()
      user.UpdatedAt    = time.Now().Unix()
      user.LoginToken   = RandString(30)

      session, err := mgo.Dial("localhost:27017")
      if err != nil {
        panic(err)
      }
      defer session.Close()

      // Optional. Switch the session to a monotonic behavior.
      session.SetMode(mgo.Monotonic, true)

      c := session.DB("spouttv").C("users")

      // Index
      index := mgo.Index{
        Key:        []string{"email", "username"},
        Unique:     true,
        DropDups:   true,
        Background: true,
        Sparse:     true,
      }
     
      err = c.EnsureIndex(index)


      if err != nil {
        panic(err)
      }

      err = c.Insert(user)
      if err != nil {
        panic(err)
        userErr = fmt.Errorf("Duplicate Field")
      }
    }
  } else {
    userErr = fmt.Errorf("Invalid Email %s", user.Email)
  }

  return
}

func ValidateEmail(email string) (status bool, err error) {
  status, err = regexp.MatchString(`^([\w\+\.]+@\w+\.\w+)$`, email)
  return 
}

func EncryptPassword(password []byte) ([]byte, error) {
  defer clear(password)
  return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func LoginToken(tmpUser LoginUser) (User, error) {
  defer clear(password)
  defer session.Close()

  // ep, _:= bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
  err := bcrypt.CompareHashAndPassword(user.EncryptedPassword, password)

  if tmpUser.Email != "" {
    // find user by email
    
  }

  if err != nil {

    return fmt.Errorf("Invalid Password")
  } else {
    token := RandString(30)
    user.LoginToken = token;

    return nil
  }
  // *user.LastSignInAt = time.Now().Unix()
}

func RandString(n int) string {
  const alphanum = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_+-=[]\{}|;':",./<>?`
  symbols := big.NewInt(int64(len(alphanum)))
  states := big.NewInt(0)
  states.Exp(symbols, big.NewInt(int64(n)), nil)
  r, err := rand.Int(rand.Reader, states)
  if err != nil {
    panic(err)
  }
  var bytes = make([]byte, n)
  r2 := big.NewInt(0)
  symbol := big.NewInt(0)
  for i := range bytes {
    r2.DivMod(r, symbols, symbol)
    r, r2 = r2, r
    bytes[i] = alphanum[symbol.Int64()]
  }
  return string(bytes)
}
