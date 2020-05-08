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
	"github.com/spf13/cast"
)

//ParseJSON function is Parse Data To Json
func ParseJSON(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return json
}

// DecodeToken is analyze JWT token
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
func CreateFile(file multipart.File, name string, userID uint) (string, error) {
	dirname := fmt.Sprintf("./images/%d", userID)
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		os.Mkdir(dirname, 0777)
	}
	extention := SplitExtention(name)
	filename := fmt.Sprintf("./images/%d/uploaded_%d.%s", userID, time.Now().UnixNano(), extention)
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
func RemoveFile(uid uint) error {
	dirname := fmt.Sprintf("./images/%d", uid)
	err := os.RemoveAll(dirname)
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

// DeleteDependent is function
func DeleteDependent(uid uint) {
	err = DB.Unscoped().Where("user_id=?", cast.ToUint(uid)).Delete(Work{}).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	//Userのbucketも削除する
}

// CreateSkills is function
func CreateSkills(arraySTR string, wid uint) error {
	var err error
	if arraySTR != "" {
		strs := strings.Split(arraySTR, ",")
		for _, str := range strs {
			var skill Skill
			skill.Name = str
			skill.WorkID = wid
			err = DB.Create(&skill).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return err
}

// RemoveSkills is function
func RemoveSkills(wid uint) error {
	DB.Unscoped().Where("work_id=?", wid).Delete(Skill{})
	var err error
	return err
}
