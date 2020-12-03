package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var dataDir string

func main() {
	addr := flag.String("p", "8888", "http server address")
	dir := flag.String("d", "./", "data file address")

	flag.Parse()

	localPort := *addr
	dataDir = *dir

	http.Handle("/trace1.data", http.HandlerFunc(SendTrace1))
	http.Handle("/trace2.data", http.HandlerFunc(SendTrace2))
	http.Handle("/api/finished", http.HandlerFunc(GetCheckSum))

	log.Printf("listening on %s, data dir: %s", localPort, dataDir)
	log.Printf("try curl http://localhost:%s/trace1.data", localPort)
	log.Printf("try curl http://localhost:%s/trace2.data", localPort)
	err := http.ListenAndServe(":"+localPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func GetCheckSum(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	sum, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln("failed to get checksum")
	}

	ioutil.WriteFile("result.txt", sum, 0644)
	log.Println("checksum done")

	writer.WriteHeader(200)
}

func SendTrace1(writer http.ResponseWriter, request *http.Request) {
	data, err := os.Open(dataDir + "/trace1.data")
	if err != nil {
		log.Fatalln("failed to load trace1.data")
	}

	defer data.Close()
	log.Println("sending trace1...")
	io.Copy(writer, bufio.NewReader(data))

	log.Println("trace1 done")
}

func SendTrace2(writer http.ResponseWriter, request *http.Request) {
	data, err := os.Open(dataDir + "/trace2.data")
	if err != nil {
		log.Fatalln("failed to load trace2.data")
	}

	defer data.Close()
	log.Println("sending trace2...")
	io.Copy(writer, bufio.NewReader(data))
	log.Println("trace2 done")
}
