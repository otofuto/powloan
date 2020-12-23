package main

import (
	"fmt"
	"log"
	"os"
	"net/smtp"
	"net/http"
	"encoding/json"
	"github.com/otofuto/powloan/pkg/database/koes"
)

var port string

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// static
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))));

	//API
	http.HandleFunc("/Signup", SignupHandle)
	http.HandleFunc("/Koes", UploadKoeHandle)

	log.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func SignupHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)
		from := "kyaorensapowroan@gmail.com"
		to := "kyaorensapowroan@gmail.com"

		//kyaorensa1945pow
		//0242nozomi
		auth := smtp.PlainAuth("", from, "bzwvrihiynrgeahr", "smtp.gmail.com")

		msg := []byte("" +
			"From: システム\r\n" +
			"To: 広報部\r\n" +
			"Subject: 新規ご入信がありました\r\n\r\n" +
			"おなまえ: \"" + r.FormValue("name") + "\"\r\n" +
			"性別: " + GetSex(r.FormValue("sex")) + "\r\n" +
			"メールアドレス: \"" + r.FormValue("email") + "\"\r\n" +
			"自己紹介: " + r.FormValue("description") + "\r\n")

		err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "true")
	} else {
		http.Error(w, "method not allowed.", 405)
	}
}

func GetSex(sex string) string {
	if sex == "0" {
		return "男性"
	} else if sex == "1" {
		return "女性"
	} else if sex == "2" {
		return "その他"
	}
	return "不明"
}

func UploadKoeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)
		k := koes.Koes {
			Who: r.FormValue("who"),
			Comment: r.FormValue("comment"),
		}
		if k.Insert() {
			fmt.Fprintf(w, "1")
		} else {
			http.Error(w, "insert failed.", 500)
		}
	} else if r.Method == http.MethodGet {
		koes, err := koes.All()
		if err != nil {
			log.Println(err);
			http.Error(w, "failed to get all koes.", 500)
			return
		}
		bytes, err := json.Marshal(koes)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to marshal json koes.", 500)
			return
		}
		fmt.Fprintf(w, string(bytes))
	} else {
		http.Error(w, "method not allowed.", 405)
	}
}