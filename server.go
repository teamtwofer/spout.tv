package main

import (
  "fmt"
  "net/http"
  "spouttv"
  "encoding/json"
  "io/ioutil"
)

func main() {
  fmt.Println("Hello, world")
  http.HandleFunc("/user", handleUser)
  http.HandleFunc("/user/login", handleLogin)
  http.HandleFunc("/streamer/", handleUserEvents)
  http.HandleFunc("/", handleIndex)
  // user := &spouttv.User{}
  // user.Name = "potato"
  // fmt.Println(user.Name)

  // _ = spouttv.FindUserById("538382a73495963937a20eb2")

  // fmt.Printf("%#v \n", u)

  fmt.Println("Starting Server...")
  http.ListenAndServe(":3000", nil)

  // 538304353495963937a20eaf
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
  writer.Header().Set("Content-Type", "text/html")
  fileName := ""
  fmt.Println(request.URL.Path)
  switch request.URL.Path {
  case "/":
    fileName = "index.html"
  case "/application.js":
    fileName = "application.js"
  case "/style.css":
    fileName = "style.css"
  }

  fmt.Println(fileName)

  if fileName == "" {
    http.Error(writer, "Page Not Found", 404)
  } else {
    http.ServeFile(writer, request, "assets/"+fileName)
  }
}

func handleUser(writer http.ResponseWriter, request *http.Request) {
  // so from here basically all a person can do is:
  // 1) Edit their account.
  // 2) Create an account
  // This means no GET allowed!
  writer.Header().Set("Content-Type", "application/json")
  switch request.Method {
  case "GET":
    http.Error(writer, "Page Not Found", http.StatusNotFound)
  case "POST":
    handleCreateUser(writer, request)
  case "PUT":
    handleEditUser(writer, request)
  }
}

func handleCreateUser(writer http.ResponseWriter, request *http.Request) { 
  // need to handle errors like what if the post is all kinds of wrong?
  tmpUser := &spouttv.NewUser{}

  body, err := ioutil.ReadAll(request.Body)
  if err != nil {
    panic(err.Error())
  }

  json.Unmarshal(body, tmpUser)

  fmt.Println(tmpUser.Name)
  fmt.Println(tmpUser.Email)
  fmt.Println(len(tmpUser.Password))
  fmt.Println(tmpUser.Password)

  if tmpUser.Username == "" || tmpUser.Name == "" || tmpUser.Email == "" || len(tmpUser.Password) == 0 {
    http.Error(writer, "Missing Parameters", 403)
  }


  user, err := spouttv.CreateUser(*tmpUser)

  if err != nil {
    http.Error(writer, "Invalid or Missing Parameters", 403)
  } else {
    writer.WriteHeader(http.StatusCreated)
    data, err := json.Marshal(user)
    if err != nil {
      panic(err.Error())
    }
    fmt.Fprintf(writer, string(data))
  }
}

func handleLogin(writer http.ResponseWriter, request *http.Request) {
  writer.Header().Set("Content-Type", "application/json")
  
  tmpUser := &spouttv.LoginUser{}

  body, err := ioutil.ReadAll(request.Body)
  if err != nil {
    panic(err.Error())
    // http.Error(writer, "Invalid or Missing Parameters", 403)
  }

  json.Unmarshal(body, tmpUser)

  if tmpUser.Email == "" && tmpUser.Username == "" {
    http.Error(writer, "Invalid or Missing Parameters", 403)
  } else {
    if tmpUser.Password == "" {
      http.Error(writer, "Invalid or Missing Parameters", 403)
    } else {
      user, err := spouttv.LoginToken(*tmpUser)
      if err != nil {
        // panic(err)
        http.Error(writer, "Invalid Username or Password", 403)
      } else {
        data, err := json.Marshal(user)
        if err != nil {
          panic(err.Error())
        }
        fmt.Fprintf(writer, string(data))        
      }
    }
  }
}

func handleEditUser(writer http.ResponseWriter, request *http.Request) {

}

func handleUserEvents(writer http.ResponseWriter, request *http.Request) {
  // Acceptable paths:
  // /streamer/login
  // /streamer/:user_name
  // The first is a POST, the latter is a GET. 

}