package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

// Gets

// GetWorkByID 1@works@ url, name, description, image_url, \
// user{id, name, image_url}, \
// []work_items[{id, image_url, body}]
// work_items SQL BY WHERE workItemID
func GetWorkByID(w http.ResponseWriter, r *http.Request) {
	workID := chi.URLParam(r, "workID")
	type ResultWork struct {
		ID          uint
		Name        string
		ImageURL    string
		Description string
		URL         string
		UserID      uint
		UserName    string
		WorkItems   []*WorkItem
	}
	var rw ResultWork
	id := cast.ToUint(workID)
	work, err := FetchWorkByID(DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user, err := FetchUserByID(DB, work.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	items, err := FetchWorkItemsByWorkID(DB, work.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	rw.ID = work.ID
	rw.Name = work.Name
	rw.Description = work.Description
	rw.ImageURL = work.ImageURL
	rw.URL = work.URL
	rw.UserID = work.UserID
	rw.UserName = user.Name
	rw.WorkItems = items
	w.Write(ParseJSON(rw))
}

// GetWorks []@works@ id, name, image_url, description, url \
// user{id, name, image_url}
// work_items SQL BY WHERE workItemID
func GetWorks(w http.ResponseWriter, r *http.Request) {
	type ResultWork struct {
		ID           uint
		Name         string
		ImageURL     string
		UserID       uint
		UserName     string
		UserImageURL string
	}
	type Result struct {
		Works []*ResultWork
	}
	var res Result
	works, err := FetchWorks(DB, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, work := range works {
		var rw ResultWork
		rw.ID = work.ID
		rw.Name = work.Name
		rw.ImageURL = work.ImageURL
		rw.UserID = work.UserID
		user, err := FetchUserByID(DB, work.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		rw.UserName = user.Name
		rw.UserImageURL = user.ImageURL
		res.Works = append(res.Works, &rw)
	}
	w.Write(ParseJSON(res))
}

// GetUsers []@users@ id, name, url, image_url, works_count
func GetUsers(w http.ResponseWriter, r *http.Request) {
	type ResultUser struct {
		ID         uint
		Name       string
		ImageURL   string
		WorksCount int
	}
	type Result struct {
		Users []*ResultUser
	}
	var res Result
	users, err := FetchUsers(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	for _, user := range users {
		var ru ResultUser
		ru.ID = user.ID
		ru.Name = user.Name
		ru.ImageURL = user.ImageURL
		works, err := FetchWorks(DB, user.ID)
		if err != nil {
			log.Fatal(err)
		}
		ru.WorksCount = len(works)
		res.Users = append(res.Users, &ru)
	}
	w.Write(ParseJSON(res))
}

// GetUserByID 1@user id, password, email
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	type ResultWork struct {
		ID           uint
		Name         string
		ImageURL     string
		UserID       uint
		UserName     string
		UserImageURL string
	}
	type Result struct {
		ID           uint
		Name         string
		ImageURL     string
		URL          string
		Introduction string
		Works        []*ResultWork
	}
	var res Result
	id := cast.ToUint(userID)
	user, err := FetchUserByID(DB, id)
	works, err := FetchWorks(DB, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	for _, work := range works {
		var rw ResultWork
		rw.ID = work.ID
		rw.Name = work.Name
		rw.ImageURL = work.ImageURL
		rw.UserID = work.UserID
		rw.UserName = user.Name
		rw.UserImageURL = user.ImageURL
		res.Works = append(res.Works, &rw)
	}
	res.ID = user.ID
	res.ImageURL = user.ImageURL
	res.Name = user.Name
	res.Introduction = user.Introduction
	res.URL = user.URL
	w.Write(ParseJSON(res))
}

// GetPrivateInfo is function
func GetPrivateInfo(w http.ResponseWriter, r *http.Request) {
	type ResultUser struct {
		ID           uint
		Name         string
		Email        string
		ImageURL     string
		Introduction string
		URL          string
	}
	var rw ResultUser
	// TODO
	token := r.Header.Get("Authorization")
	token = strings.Trim(token, "Bearer%20")
	email := DecodeToken(token)
	user, err := FetchUserByEmail(DB, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	rw.ID = user.ID
	rw.Name = user.Name
	rw.Email = user.Email
	rw.ImageURL = user.ImageURL
	rw.Introduction = user.Introduction
	rw.URL = user.URL
	w.Write(ParseJSON(rw))
}

// Posts

// CreateUser 1@jwt token
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	password := r.FormValue("password")
	email := r.FormValue("email")
	user.Password = password
	user.Email = email
	user.Name = r.FormValue("name")

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = string(hash)
	// クエリ発行してDBに保存
	err = DB.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = ""
	w.Write(ParseJSON(user))
}

// RequestSession 1@jwt token
func RequestSession(w http.ResponseWriter, r *http.Request) {
	type ResultWork struct {
		Token string
	}
	// JSON 解析
	type Request struct {
		Email    string
		Password string
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var req Request
	json.Unmarshal(body, &req)

	// req.EmailからUser取得
	user, err := FetchUserByEmail(DB, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	hasedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(req.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := CreateToken(*user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var rw ResultWork
	rw.Token = token
	user.Token = rw.Token
	err = DB.Save(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(rw))
}

// CreateWorks is
func CreateWorks(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID uint
	}
	var work Work
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	work.Name = r.FormValue("name")
	work.Description = r.FormValue("description")
	work.URL = r.FormValue("url")
	work.UserID = cast.ToUint(r.FormValue("user_id"))

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := UploadFileToBucket(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(result)
	// // AWS S3のURLをwork.ImageURLに挿入
	// work.ImageURL = string(result)

	err = DB.Create(&work).Error
	if err != nil {
		log.Fatal(err)
	}
	var res Result
	res.ID = work.ID
	w.Write(ParseJSON(res))
}

// Puts

// UpdateUser is function
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID uint
	}
	userID := chi.URLParam(r, "userID")
	id := cast.ToUint(userID)
	user, err := FetchUserByID(DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var res Result
	////////////
	if user.ID == cast.ToUint(r.FormValue("user_id")) {
		if cast.ToBool(r.FormValue("windowOpt")) {
			user.Name = r.FormValue("name")
			user.URL = r.FormValue("url")
			user.Introduction = r.FormValue("introduction")
			// user.ImageURL
		} else {
			user.Email = r.FormValue("email")
			pass := r.FormValue("password")
			hashedPassword := user.Password
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			userPass := r.FormValue("new_password")
			hash, err := bcrypt.GenerateFromPassword([]byte(userPass), 10)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			user.Password = string(hash)
		}
	}
	////////////
	err = DB.Save(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res.ID = user.ID
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// UpdateWorks is handler
func UpdateWorks(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID uint
	}
	var res Result
	workID := chi.URLParam(r, "workID")
	id := cast.ToUint(workID)
	work, err := FetchWorkByID(DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	//////
	work.Name = r.FormValue("name")
	work.Description = r.FormValue("description")
	work.URL = r.FormValue("url")
	// work.ImageURL = r.FormValue("file")
	//////
	err = DB.Save(&work).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res.ID = work.ID
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// UpdateWorkItems is handler with PUT
func UpdateWorkItems(w http.ResponseWriter, r *http.Request) {
	workID := chi.URLParam(r, "workID")
	type Result struct {
		ID uint
	}
	var res Result
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	// for _, data := range req.DataSets {
	// 	if cast.ToUint(data.ID) > 0 {
	// 		item, err := FetchWorkItemByID(DB, data.ID)
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusNotFound)
	// 		}
	// 		item.Body = data.Body
	// 		item.ImageURL = data.ImageURL
	// 		err = DB.Save(&item).Error
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusBadRequest)
	// 		}
	// 	} else {
	// 		var item WorkItem
	// 		item.WorkID = cast.ToUint(workID)
	// 		item.Body = data.Body
	// 		item.ImageURL = data.ImageURL
	// 		err = DB.Create(&item).Error
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusBadRequest)
	// 		}
	// 	}
	// }
	////
	res.ID = cast.ToUint(workID)
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// Deletes
