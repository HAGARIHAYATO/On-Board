package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

// Gets

// GetWorksPerDay is
func GetWorksPerDay(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Today     int
		Yesterday int
		DBY       int
		TDA       int
		FDA       int
	}
	works, err := FetchWorks(DB, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var list []time.Time
	for _, el := range works {
		list = append(list, el.CreatedAt)
	}
	m := CreateDateObject(list)
	var res Result
	res.Today = m["today"]
	res.Yesterday = m["yesterday"]
	res.DBY = m["dby"]
	res.TDA = m["tda"]
	res.FDA = m["fda"]
	w.Write(ParseJSON(res))
}

// GetSkills is
func GetSkills(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Skills map[string]int
	}
	skills, err := FetchSkills(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var list []string
	for _, el := range skills {
		list = append(list, el.Name)
	}
	m := RemoveDuplicationName(list)
	var res Result
	res.Skills = m
	w.Write(ParseJSON(res))
}

// GetWorkByID 1@works@ url, name, description, image_url, \
// user{id, name, image_url}, \
// []work_items[{id, image_url, body}]
// work_items SQL BY WHERE workItemID
func GetWorkByID(w http.ResponseWriter, r *http.Request) {
	workID := chi.URLParam(r, "workID")
	type AssessmentUser struct {
		ImageURL string
		UserID   uint
		Name     string
	}
	type ResultWork struct {
		ID                  uint
		Name                string
		ImageURL            string
		Description         string
		URL                 string
		CacooURL            string
		IsPublished         bool
		UserID              uint
		UserName            string
		UserImageURL        string
		UserGithubToken     string
		GHR                 string
		Assessment          map[string]int
		AssessmentUsers     []AssessmentUser
		AssessmentUserCount int
		WorkItems           []*WorkItem
		Skills              []*Skill
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
	skills, _ := FetchSkillsByWorkID(DB, work.ID)
	assessments, _ := FetchAssessmentsByWorkID(DB, work.ID)
	assess := AverageAssessment(assessments)

	for _, assessment := range assessments {
		var assessmentUser AssessmentUser
		assessmentUser.UserID = assessment.UserID
		asUser, _ := FetchUserByID(DB, assessment.UserID)
		assessmentUser.ImageURL = asUser.ImageURL
		assessmentUser.Name = asUser.Name
		rw.AssessmentUsers = append(rw.AssessmentUsers, assessmentUser)
	}

	rw.ID = work.ID
	rw.Name = work.Name
	rw.Description = work.Description
	rw.ImageURL = work.ImageURL
	rw.URL = work.URL
	rw.UserID = work.UserID
	rw.UserName = user.Name
	rw.UserImageURL = user.ImageURL
	rw.WorkItems = items
	rw.Skills = skills
	rw.Assessment = assess
	rw.AssessmentUserCount = len(assessments)
	rw.CacooURL = work.CacooURL
	rw.GHR = work.GHR
	rw.UserGithubToken = user.GitHubToken
	rw.IsPublished = cast.ToBool(work.IsPublished)
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
		Skills       []*Skill
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
		skills, _ := FetchSkillsByWorkID(DB, work.ID)
		rw.Skills = skills
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

// GetWorksIDs is
func GetWorksIDs(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		IDs []uint
	}
	works, err := FetchWorks(DB, 0)
	if err != nil {
		log.Fatal(err)
	}
	var array []uint
	for _, work := range works {
		array = append(array, work.ID)
	}
	var res Result
	res.IDs = array
	w.Write(ParseJSON(res))
}

// GetAssessmentWithLoginUser is
func GetAssessmentWithLoginUser(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Assessment *Assessment
		Status     int
	}
	wid := cast.ToUint(chi.URLParam(r, "workID"))
	uid := cast.ToUint(r.FormValue("user_id"))
	assessment, err := FetchAnAssessment(DB, uid, wid)
	var res Result
	res.Assessment = assessment
	if err != nil {
		res.Status = http.StatusNotFound
	} else {
		res.Status = http.StatusOK
	}
	w.Write(ParseJSON(res))
}

// GetUsersIDs is
func GetUsersIDs(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		IDs []uint
	}
	users, err := FetchUsers(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var array []uint
	for _, user := range users {
		array = append(array, user.ID)
	}
	var res Result
	res.IDs = array
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
		Skills       []*Skill
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
		Email        string
		GitHubToken  string
		QiitaName    string
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
		skills, err := FetchSkillsByWorkID(DB, work.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		rw.Skills = skills
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
	res.Email = user.Email
	res.GitHubToken = user.GitHubToken
	res.QiitaName = user.QiitaName
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
		GitHubToken  string
		QiitaName    string
		IsAdmin      bool
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
	rw.GitHubToken = user.GitHubToken
	rw.QiitaName = user.QiitaName
	rw.IsAdmin = false
	if user.IsAdmin == true {
		rw.IsAdmin = true
	}
	w.Write(ParseJSON(rw))
}

// Posts

// CreateUser 1@jwt token
func CreateUser(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID    uint
		Token string
	}
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
	token, err := CreateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var res Result
	res.ID = user.ID
	res.Token = token
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
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
	work.CacooURL = r.FormValue("cacoo_url")
	work.IsPublished = cast.ToBool(r.FormValue("is_published"))
	work.UserID = cast.ToUint(r.FormValue("user_id"))
	work.GHR = r.FormValue("gh_repo")
	file, fileHeader, _ := r.FormFile("file")
	if file != nil {
		filename, err := CreateFile(file, fileHeader.Filename, work.UserID)
		result, err := UploadFileToBucket(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		work.ImageURL = string(result.Location)
		err = RemoveFile(work.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	err = DB.Create(&work).Error
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 6; i++ {
		var item WorkItem
		item.WorkID = work.ID
		err = DB.Create(&item).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	array := r.FormValue("array")
	CreateSkills(array, work.ID)
	var res Result
	res.ID = work.ID
	w.Write(ParseJSON(res))
}

// PostAssessment is
func PostAssessment(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Status int
	}
	wid := cast.ToUint(chi.URLParam(r, "workID"))
	uid := cast.ToUint(r.FormValue("user_id"))
	var assessment Assessment
	DB.Where(Assessment{UserID: uid, WorkID: wid}).Attrs(Assessment{
		UserID:   uid,
		WorkID:   wid,
		Function: cast.ToInt(r.FormValue("function")),
		UIX:      cast.ToInt(r.FormValue("uix")),
		BugSafe:  cast.ToInt(r.FormValue("bug_safe")),
		Content:  cast.ToInt(r.FormValue("content")),
		MDN:      cast.ToInt(r.FormValue("mdn")),
	}).FirstOrCreate(&assessment)
	assessment.Function = cast.ToInt(r.FormValue("function"))
	assessment.UIX = cast.ToInt(r.FormValue("uix"))
	assessment.BugSafe = cast.ToInt(r.FormValue("bug_safe"))
	assessment.Content = cast.ToInt(r.FormValue("content"))
	assessment.MDN = cast.ToInt(r.FormValue("mdn"))
	err = DB.Save(&assessment).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var res Result
	if err != nil {
		res.Status = http.StatusBadRequest
	} else {
		res.Status = http.StatusOK
	}
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
	if user.ID == cast.ToUint(r.FormValue("user_id")) {
		if cast.ToBool(r.FormValue("windowOpt")) {
			user.Name = r.FormValue("name")
			user.Email = r.FormValue("email")
			user.GitHubToken = r.FormValue("github")
			user.QiitaName = r.FormValue("qiita")
			user.URL = r.FormValue("url")
			user.Introduction = r.FormValue("introduction")
			preURL := user.ImageURL
			file, fileHeader, _ := r.FormFile("file")
			if file != nil {
				filename, err := CreateFile(file, fileHeader.Filename, user.ID)
				result, err := UploadFileToBucket(filename)
				if err != nil {
					fmt.Println("--------%v--------", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				user.ImageURL = string(result.Location)
				err = RemoveFile(user.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// TODO eliminate ... aws sdk config about iam role
				if preURL != "" {
					err = DeleteFileByBucket(preURL)
					if err != nil {
						fmt.Println("--------%v--------", err)
					}
				}
			}
		} else {
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
	preURL := work.ImageURL
	work.Name = r.FormValue("name")
	work.CacooURL = r.FormValue("cacoo_url")
	work.Description = r.FormValue("description")
	work.IsPublished = cast.ToBool(r.FormValue("is_published"))
	work.URL = r.FormValue("url")
	work.GHR = r.FormValue("gh_repo")
	file, fileHeader, _ := r.FormFile("file")
	if file != nil {
		filename, err := CreateFile(file, fileHeader.Filename, work.UserID)
		result, err := UploadFileToBucket(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		work.ImageURL = string(result.Location)
		err = RemoveFile(work.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if preURL != "" {
			err = DeleteFileByBucket(preURL)
			if err != nil {
				fmt.Println("--------%v--------", err)
			}
		}
	}
	err = DB.Save(&work).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	RemoveSkills(work.ID)
	array := r.FormValue("array")
	CreateSkills(array, work.ID)
	res.ID = work.ID
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// UpdateWorkItem is handler with PUT
func UpdateWorkItem(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID uint
	}
	var res Result
	workID := cast.ToUint(chi.URLParam(r, "workID"))
	itemid := r.FormValue("id")
	uid := cast.ToUint(r.FormValue("uid"))
	item, err := FetchWorkItemByID(DB, cast.ToUint(itemid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	preURL := item.ImageURL
	item.Body = r.FormValue("body")
	file, fileHeader, _ := r.FormFile("file")
	if file != nil {
		filename, err := CreateFile(file, fileHeader.Filename, uid)
		result, err := UploadFileToBucket(filename)
		if err != nil {
			fmt.Println("--------%v--------", err)
		}
		item.ImageURL = string(result.Location)
		err = RemoveFile(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if preURL != "" {
			err = DeleteFileByBucket(preURL)
			if err != nil {
				fmt.Println("--------%v--------", err)
			}
		}
	}
	err = DB.Save(&item).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.ID = workID
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// Deletes

// DeleteWorks is handler
func DeleteWorks(w http.ResponseWriter, r *http.Request) {
	workID := chi.URLParam(r, "workID")
	type Result struct {
		ID     uint
		Status int
	}
	work, err := FetchWorkByID(DB, cast.ToUint(workID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	RemoveSkills(work.ID)
	err = DB.Unscoped().Delete(&work).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = DB.Unscoped().Where("work_id=?", cast.ToUint(workID)).Delete(WorkItem{}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var res Result
	res.Status = http.StatusOK
	res.ID = cast.ToUint(workID)
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// DeleteUser is handler
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID     uint
		Status int
	}
	token := r.Header.Get("Authorization")
	token = strings.Trim(token, "Bearer%20")
	email := DecodeToken(token)
	user, err := FetchUserByEmail(DB, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	works, _ := FetchWorks(DB, user.ID)
	for _, work := range works {
		RemoveSkills(work.ID)
		err = DB.Delete(&work).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	var res Result
	res.ID = user.ID

	err = DB.Unscoped().Delete(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	DeleteDependent(res.ID)
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}

// ExecutedUser is
func ExecutedUser(w http.ResponseWriter, r *http.Request) {
	// json => uid
	type Request struct {
		UID uint
	}
	type Result struct {
		ID     uint
		Status int
	}
	var res Result
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var req Request
	json.Unmarshal(body, &req)
	user, err := FetchUserByID(DB, req.UID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	works, _ := FetchWorks(DB, user.ID)
	for _, work := range works {
		RemoveSkills(work.ID)
		DB.Unscoped().Where("work_id=?", work.ID).Delete(WorkItem{})
		err = DB.Unscoped().Delete(&work).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	res.ID = user.ID
	err = DB.Unscoped().Delete(&user).Error
	if err != nil {
		res.Status = http.StatusBadRequest
	} else {
		res.Status = http.StatusOK
	}
	w.Write(ParseJSON(res))
}

// PostInformation is
func PostInformation(w http.ResponseWriter, r *http.Request) {
	// json => { if } uid == 0 => all__user { else } user(uid)
	// json => message
	// json => title
	type Request struct {
		UID     uint
		Message string
		Title   string
	}
	type Result struct {
		Status int
	}
	var res Result
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var req Request
	json.Unmarshal(body, &req)
	var info Info
	info.UserID = req.UID
	info.Title = req.Title
	info.Message = req.Message
	err = DB.Create(&info).Error
	if err != nil {
		res.Status = http.StatusBadRequest
	} else {
		res.Status = http.StatusOK
	}
	w.Write(ParseJSON(res))
}

// GetInformation is
func GetInformation(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Info []*Info
	}
	userID := chi.URLParam(r, "userID")
	id := cast.ToUint(userID)
	aLotOfInfo, err := FetchInfoByUID(DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var res Result
	res.Info = aLotOfInfo
	w.WriteHeader(http.StatusOK)
	w.Write(ParseJSON(res))
}
