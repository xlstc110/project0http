package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type stage struct {
	box [][]string
}

var templates = template.Must(template.ParseFiles("template.html", "decision.html"))

func main() {
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/decision", decision)
	http.ListenAndServe(":8080", nil)
}

func mainpage(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "mainpage.html", nil)
}

func decision(w http.ResponseWriter, r *http.Request) {
	decision := r.FormValue("decision")
	fmt.Println(decision)
	if decision == "1" {
		box := []string{
			"box 1", "box 2", "box 3", "box 4", "box 5", "box 6", "box 7", "box 8", "box 9", "box 10",
		}

		p := stage{box: [][]string{box}}

		templates.ExecuteTemplate(w, "pickbox.html", nil)

		fmt.Println(p)
		// 	//generate a random nonrepeated array to mix the order the box value.
		// 	rand.Seed(time.Now().UnixNano())
		// 	order := rand.Perm(10)

		// 	//Prize list of remaining box value
		// 	Prize := []int{1000, 2500, 5000, 10000, 20000, 30000, 45000, 60000, 80000, 100000}

		// 	//make map dictionary for box
		// 	Value := make(map[string]int)

		// 	//giving value to the boxes
		// 	for i := 0; i < len(box); i++ {
		// 		Value[box[i]] = Prize[order[i]]
		// 	}
		// 	fmt.Println(box)

		// 	lucky1 := r.FormValue("pickedbox")
		// 	luckybox := Value[box[lucky1-1]]

		// 	//drop boxes that have been selected
		// 	box[lucky1-1] = " "

		// 	//create a slice to restore all the box that have been selected
		// 	picked := []int{}
		// 	picked = append(picked, lucky1)
	}
}
