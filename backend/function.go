package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/k-washi/jwt-decode/jwtdecode"
)

//ParseJSON function is Parse Data To Json
func ParseJSON(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return json
}

func DecodeToken(token string) string {
	hCS, err := jwtdecode.JwtDecode.DecomposeFB(token)
	if err != nil {
		log.Fatalln("Error : ", err)
	}
	payload, err := jwtdecode.JwtDecode.DecodeClaimFB(hCS[1])
	if err != nil {
		log.Fatalln("Error :", err)
	}
	return payload.Email
}

//CreateToken is credential Json Request
func CreateToken(user User) (string, error) {
	var err error

	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   user.Name,
	})

	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println("-----------------------------")
	fmt.Println("tokenString:", tokenString)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

// CORS is middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT,OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CreateFile is function whitch create file in /images/
func CreateFile(file multipart.File, name string) (string, error) {
	extention := SplitExtention(name)
	filename := fmt.Sprintf("./images/uploaded_%d.%s", time.Now().UnixNano(), extention)
	fmt.Println(filename)
	saveFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer saveFile.Close()
	_, err = io.Copy(saveFile, file)
	if err != nil {
		log.Fatal(err)
	}
	return filename, err
}

// RemoveFile is function whitch delete file by /images/
func RemoveFile(filename string) error {
	err := os.Remove(filename)
	return err
}

// SplitExtention is function
func SplitExtention(str string) string {
	slice := strings.Split(str, ".")
	if len(slice) == 0 {
		return ""
	}
	return slice[len(slice)-1]
}
