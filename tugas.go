package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	if r.Method == "GET" {
		obj:= About {
			Tentang: "ini adalah website list game dari semua developer games oleh Yohanes Hubert,silahkan di cek!",
		}
		res, err := json.Marshal(obj)
		if err != nil {
			http.Error(w, "Gagal konversi ke Json", http.StatusInternalServerError)
		}
		w.Write(res)
		w.Header().Set("Content-Type","application/json")
	} else {
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
	http.HandleFunc("/about", HandlerAbout)
	http.HandleFunc("/listgemscool", Handlergm)
	http.HandleFunc("/listhoyo", HandlerHoyo)
	http.HandleFunc("/listea", HandlerEA)
	http.HandleFunc("/listrg", HandlerRG)
	http.HandleFunc("/listval", HandlerVal)
	fmt.Println("Server running at http://localhost:8080")
	fmt.Println("Pastikan baca di /about dulu ya sebelum mengakases :)")
	http.ListenAndServe(":8080", nil)
}