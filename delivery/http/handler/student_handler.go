package handler

import (
	"fmt"
	"github.com/Yuideg/firstApp/Student"
	"github.com/Yuideg/firstApp/entity"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// AdminNewsHandler handles category handler admin requests
type StudentHandler struct {
	tmpl           *template.Template
	studentService Student.StudentServices
}

// NewAdminNewsHandler initializes and returns new AdminCateogryHandler
func NewStudentHandler(T *template.Template, NS Student.StudentServices) *StudentHandler {
	return &StudentHandler{tmpl: T, studentService: NS}
}

// AdminNews handle requests on route /admin/newss

func (ach *StudentHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	ach.tmpl.ExecuteTemplate(w, "student.html", nil)
}


func (ach *StudentHandler) StudentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("students handler ")
	Students, err := ach.studentService.Students()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ach.tmpl.ExecuteTemplate(w, "index.html", Students)
}

// AdminNews handle requests on route /admin/newss
func (ach *StudentHandler) StudentHandler(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	news, _ := ach.studentService.Student(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ach.tmpl.ExecuteTemplate(w, "welcome.html", news)
}

// AdminNewsNew hanlde requests on route /admin/new
func (ach *StudentHandler) RegisterStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	st := entity.StudentInfo{}

	if r.Method == http.MethodPost {
		st.FullName = r.PostFormValue("fullname")
		st.Email = r.PostFormValue("email")
		mf, fh, err := r.FormFile("image")
		if err != nil {
			ach.tmpl.ExecuteTemplate(w, "Error.html", nil)
			return
		}
		defer mf.Close()

		st.Image = fh.Filename
		WriteFile(&mf, fh.Filename)
		fmt.Println(st)
		_, err2 := ach.studentService.RegisterStudent(st)
		if len(err2) > 0 {
			fmt.Println("there is error in registering student ")
			ach.tmpl.ExecuteTemplate(w, "Error.html", nil)
			return
		}
		Students, er := ach.studentService.Students()
		if len(er) > 0 {
			fmt.Println("there is error in registering student ")
			ach.tmpl.ExecuteTemplate(w, "Error.html", nil)
			return
		}
		ach.tmpl.ExecuteTemplate(w, "index.html", Students)

	}

}
func (ach *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update start")
	params := mux.Vars(r)
	id, er := strconv.Atoi(params["id"])
	if er != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	student, _ := ach.studentService.Student(id)
	ach.tmpl.ExecuteTemplate(w, "update.html", student)
}
// AdminNewsUpdate handle requests on /admin/events/update
func (ach *StudentHandler) UpdateStudentInfoHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	// Parse the form data
	err=r.ParseMultipartForm(32 << 20)
	fmt.Println("error =",err)
	fmt.Println("update running ")

	id, er := strconv.Atoi(r.FormValue("hid"))
	student, _ := ach.studentService.Student(id)
    fmt.Println("half",id)
	fmt.Println("stu ",student)
	if er != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	student.FullName = r.FormValue("fullname")
	student.Email = r.FormValue("email")
	mf, fh, err := r.FormFile("image")
	if err != nil {
		fmt.Println("129")
		ach.tmpl.ExecuteTemplate(w, "Error.html", nil)
		return
	}
	defer mf.Close()

	student.Image = fh.Filename
	WriteFile(&mf, fh.Filename)
	fmt.Println(student)
	st, err2 := ach.studentService.UpdateStudentInfor(student)
	fmt.Println("new ",st)
	if len(err2) > 0 {
		ach.tmpl.ExecuteTemplate(w, "Error.html", nil)
		return
	}
	http.Redirect(w, r, "/user/students/all", http.StatusSeeOther)

}
// AdminNewsDelete handle requests on route /events/categories/delete
func (ach *StudentHandler) StudentDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete Mode is Running ")
	params := mux.Vars(r)
	id, er := strconv.Atoi(params["id"])
	if er != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//student, _ := ach.studentService.Student(id)
	//ach.tmpl.ExecuteTemplate(w, "update.html", student)

		_, err2 := ach.studentService.DeleteStudent(id)

		if len(err2) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	http.Redirect(w, r, "/user/students/all", http.StatusSeeOther)
}

func WriteFile(mf *multipart.File, fname string) {
	wd := "/home/salemariam/go/src/github.com/Yuideg/firstApp/"
	path := filepath.Join(wd, "ui", "assets", "img", fname)

	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
