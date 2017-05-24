package main

//go:generate go build  -o bin/web-log_linux-amd64 web-log.go

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var port *string
var logDir string

func main() {
	port = flag.String("p", "8010", "port to listen")
	logDirArg := flag.String("d", "./", "dir to log")
	flag.Parse()
	log.Println("listen on :" + *port)
	logDir, _ = filepath.Abs(*logDirArg)

	http.HandleFunc("/", handle)               //设置访问的路由
	err := http.ListenAndServe(":"+*port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func logToFile(logFile string, content string) {
	fd, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fdTime := time.Now().Format("2006-01-02 15:04:05")
	fdContent := strings.Join([]string{"======", fdTime, "=====\n", content, "\n"}, "")
	buf := []byte(fdContent)
	fd.Write(buf)
	fd.Close()
}
func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	uri := r.RequestURI
	uri = uri[1:]
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", uri)
	if !matched || len(uri) > 30 {
		return
	}
	month := time.Now().Format("0000-00")
	ext := ".log"
	path := logDir
	filename := uri
	logFile := path + "/" + filename + "-" + month + ext
	//log.Println("log to file : [" + logFile + "]")
	defer r.Body.Close()
	result, _ := ioutil.ReadAll(r.Body)
	logToFile(logFile, string(result[0:]))
}
