package fileserver

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var uploadFolder *string

func StartFileServer() {
	uploadFolder = flag.String("uploadfolder", "/home/work/data/stress", "folder to store the upload files") 
	flag.Parse()
	fmt.Println(*uploadFolder)
	if _, err := os.Stat(*uploadFolder); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.MkdirAll(*uploadFolder, 0777)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/exist", ensureFileExist)
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir(*uploadFolder))))
	err := http.ListenAndServe(":8060", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func ensureFileExist(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var filename, category string
	if len(r.Form["filename"]) > 0 {
		filename = r.Form["filename"][0]
	} else {
		w.WriteHeader(400)
		fmt.Fprintln(w, "filename si MUST!")
		return
	}
	if len(r.Form["category"]) > 0 {
		category = r.Form["category"][0]
	}
	targetPath := *uploadFolder
	if len(strings.TrimSpace(category)) != 0 {
		targetPath += "/" + category
	}
	targetPath += "/" + filename
	fmt.Println("Targetpath: ", targetPath)
	_, err := os.Stat(targetPath)
	if os.IsNotExist(err) {
		fmt.Fprintln(w, "file Not exist!")
	} else {
		fmt.Fprintln(w, "file exist!")
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		println(r)
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("filename")
		category := r.FormValue("category")
		fmt.Println("category: ", category)
		if err != nil {
			fmt.Println(err)
			return
		}
		targetFolder := *uploadFolder
		if len(strings.TrimSpace(category)) != 0 {
			targetFolder += "/" + category
		}
		if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
			// path/to/whatever does not exist
			os.MkdirAll(targetFolder, 0777)
		}
		defer file.Close()
		filenamesep :=  strings.Split(handler.Filename, "/")
		filename := filenamesep[len(filenamesep)-1]
		fmt.Println(filename)
		f, err := os.OpenFile(targetFolder + "/" + filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "upload ok!")
	} else {
		index(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}

const tpl = `<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
 <p>
 <input type="file" name="filename" />
 </p>
 <p>
 上传文件所属category:
 <input type="text" name="category" size="50">
 </p>
 <input type="hidden" name="token" value="{...{.}...}"/>
 <input type="submit" value="upload" />
</form>
</body>
</html>`
