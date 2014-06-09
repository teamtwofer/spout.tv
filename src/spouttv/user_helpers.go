package spouttv

import (
  // "labix.org/v2/mgo/bson"
  "labix.org/v2/mgo"
)

func GetUserCollection() (*mgo.Session, *mgo.Collection, error) {
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  // defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("spouttv").C("users")

  return session, c, err
}

// func (this *User) Follow(userId bson.ObjectId) {
//   session, c, err := GetUserCollection()
//   defer session.Close()

//   // user := GetUserById(this.Id, c)
//   // PushFollower(user.following, userId)
//   // user.Save()
// }