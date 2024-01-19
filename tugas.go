package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//blueprint
type About struct {
	Tentang string
}

type Gemscool struct {
	Name string
	Release int
	Platform string
}

type Hoyoverse struct {
	Name string
	Release int
	Platform string
}

type EA struct {
	Name string
	Release int
	Platform string
}

type Rockstar struct {
	Name string
	Release int
	Platform string
}

type Valve struct {
	Name string
	Release int
	Platform string
}

func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	//validasi method
	if r.Method == "GET" { //jika method GET maka eksekusi kode dibawah
		//implementasi blueprint
		obj:= About {
			Tentang: "ini adalah website list game dari semua developer games oleh Yohanes Hubert,silahkan di cek!",
		}
		//konversi ke json
		res, err := json.Marshal(obj) // mengubah struct menjadi json
		//validasi konversi json
		if err != nil { //jika gagal konversi ke json
			http.Error(w, "Gagal konversi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)// tulis res
		//set header
		w.Header().Set("Content-Type","application/json")
	} else { // kalau menggunakan metod yang lain maka error
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Handlergm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		arrayObjects := []Gemscool{
			{"Pointblank", 2008, "PC only"},
			{"Lost Saga", 2009, "PC only"},
			{"Dragon Nest", 2010, "PC only"},
		}
		res, err := json.Marshal(arrayObjects)
		if err != nil {
			http.Error(w, "gagal konersi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type","applicaiton/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerHoyo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		arrayObjects := []Hoyoverse{
			{"Honkai Impact 3rd", 2016, "PC &Console"},
			{"Genshin Impact", 2020, "PC & Console"},
			{"Honkai Star Rail", 2022, "PC & Console"},
		}
		res, err := json.Marshal(arrayObjects)
		if err != nil {
			http.Error(w, "gagal konersi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type","applicaiton/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerEA(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		arrayObjects := []EA{
			{"The Sims 4", 2015, "PC & Console"},
			{"Battlefield 1", 2013, "PC & Console"},
			{"Command & Conquer", 2018, "PC Only"},
		}
		res, err := json.Marshal(arrayObjects)
		if err != nil {
			http.Error(w, "gagal konersi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type","applicaiton/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}	
}

func HandlerRG(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		arrayObjects := []Rockstar{
			{"Grand Thief Auto V", 2015, "PC & Console"},
			{"Red Dead Redemption", 2013, "PC &Console"},
			{"Max Payne", 2014, "PC & Console"},
		}
		res, err := json.Marshal(arrayObjects)
		if err != nil {
			http.Error(w, "gagal konersi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type","applicaiton/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func HandlerVal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		arrayObjects := []Valve{
			{"Counter Strike Global Offensive", 2012, "PC Only"},
			{"Half-Life 2", 2019, "PC Only"},
			{"Team Fortress 2", 2018, "PC Only"},
		}
		res, err := json.Marshal(arrayObjects)
		if err != nil {
			http.Error(w, "gagal konersi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type","applicaiton/json")
	} else {
		http.Error(w, "Method Tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/about", HandlerAbout) //anonymouse function, harus pascal case (besar semua di depan)
	http.HandleFunc("/listgemscool", Handlergm)
	http.HandleFunc("/listhoyo", HandlerHoyo)
	http.HandleFunc("/listea", HandlerEA)
	http.HandleFunc("/listrg", HandlerRG)
	http.HandleFunc("/listval", HandlerVal)
	fmt.Println("Server running at http://localhost:8080") // tampil di console.log
	fmt.Println("Pastikan baca di /about dulu ya sebelum mengakases :)")
	http.ListenAndServe(":8080", nil)//port koneksi ke port :8080
}