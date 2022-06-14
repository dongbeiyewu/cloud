package main

import (
	"cloud/model"
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

// 添加

func addpeople(w http.ResponseWriter, r *http.Request) {
	people := &model.People{
		Name:     r.FormValue("name"),
		Num:      r.FormValue("num"),
		Picture:  r.FormValue("pic"),
		Keywords: r.FormValue("keywords"),
	}
	people.AddPeople(people)
	people2 := &model.People{}
	p, _ := people2.GetUsers()
	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))
	t.ExecuteTemplate(w, "index.html", "")
	for _, v := range p {
		//fmt.Printf("第%v个用户是%v：\n", k+1, v)
		t.ExecuteTemplate(w, "people.html", v)
		t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+v.Keywords)
	}
}

// 处理 /upload  逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))
	t.ExecuteTemplate(w, "index.html", "")
	people := &model.People{}
	p, _ := people.GetUsers()
	for _, v := range p {
		//fmt.Printf("第%v个用户是%v：\n", k+1, v)
		t.ExecuteTemplate(w, "people.html", v)
		t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+v.Keywords)
	}
}
func update(w http.ResponseWriter, r *http.Request) {
	csvfile, err := os.Open("test/data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	read := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := read.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		people := &model.People{
			Name:     record[0],
			Num:      record[1],
			Picture:  record[2],
			Keywords: record[3],
		}
		people.AddPeople(people)
	}
	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))
	t.ExecuteTemplate(w, "index.html", "")
	people2 := &model.People{}
	p, _ := people2.GetUsers()
	for _, v := range p {
		//fmt.Printf("第%v个用户是%v：\n", k+1, v)
		t.ExecuteTemplate(w, "people.html", v)
		t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+v.Keywords)
	}
}
func showindex(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("web/people.html")
	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))
	people := &model.People{}
	p, _ := people.GetUsers()
	//json, _ := json.Marshal(p[0])
	t.ExecuteTemplate(w, "index.html", "")
	for _, v := range p {
		//fmt.Printf("第%v个用户是%v：\n", k+1, v)
		t.ExecuteTemplate(w, "people.html", v)
		t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+v.Keywords)
	}
}
func searchpeople(w http.ResponseWriter, r *http.Request) {
	people := &model.People{
		Name: r.FormValue("username"),
	}
	p, _ := people.GetUserByname()
	//json, _ := json.Marshal(p)

	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))

	t.ExecuteTemplate(w, "people.html", p)
	t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+p.Keywords)

}
func deletepeople(w http.ResponseWriter, r *http.Request) {
	people := &model.People{
		Name: r.FormValue("username"),
	}
	people.DeletePeople(people)
	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))
	t.ExecuteTemplate(w, "index.html", "")
	people2 := &model.People{}
	p, _ := people2.GetUsers()
	for _, v := range p {
		//fmt.Printf("第%v个用户是%v：\n", k+1, v)
		t.ExecuteTemplate(w, "people.html", v)
		t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+v.Keywords)
	}
}

// 上传图像接口
func uploadpicHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./uploadedpic/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
	t := template.Must(template.ParseFiles("web/index.html", "web/people.html", "web/pic.html"))
	t.ExecuteTemplate(w, "index.html", "")
	people := &model.People{}
	p, _ := people.GetUsers()
	for _, v := range p {
		t.ExecuteTemplate(w, "people.html", v)
		t.ExecuteTemplate(w, "pic.html", "http://localhost:4396/static/"+v.Keywords)
	}
}

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func main() {
	fs := http.FileServer(http.Dir("uploadedpic/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// 调用处理器处理请求
	http.HandleFunc("/main", showindex)
	http.HandleFunc("/uploadpic", uploadpicHandle) // 上传
	http.HandleFunc("/update", update)
	http.HandleFunc("/searchpeople", searchpeople)
	http.HandleFunc("/deletepeople", deletepeople)
	http.HandleFunc("/addpeople", addpeople)
	http.HandleFunc("/uploadcsv", upload)
	http.ListenAndServe(":4396", nil)
}
