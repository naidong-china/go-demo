package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"time"
)

func main() {

	const addr = "127.0.0.1:8080"
	log.Println("listen at ", addr)

	http.HandleFunc("/", fileHandler)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Disposition", "attachment; filename=test.csv")
	w.Header().Add("Content-Type", "text/csv")
	w.(http.Flusher).Flush()

	// 返回空文件
	writer := csv.NewWriter(w)
	writer.Write([]string{"序号", "国家", "数字"})
	w.(http.Flusher).Flush()

	time.Sleep(time.Duration(1) * time.Second)

	// 后续慢慢写入数据
	for i := 0; i < 20; i++ {
		data := [][]string{
			{"1", "中国", "23"},
			{"2", "美国", "23"},
		}
		writer.WriteAll(data)
		w.(http.Flusher).Flush()
		time.Sleep(time.Duration(3) * time.Second)
	}
}
