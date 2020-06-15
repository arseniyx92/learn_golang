`"golang.org/x/crypto/bcrypt"`

`type user struct {
 	UserName string
 	Password []byte
 	First    string
 	Last     string
 }`
 
 `bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)`
 
 to check (login)
 
 `err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))`
 
 cleaning sessions:
 https://www.udemy.com/course/go-programming-language/learn/lecture/6162682#overview
 