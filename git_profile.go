package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

/*http://127.0.0.1:4000/api/getUser/tiqdesign*/
var newUrl = ""
var userData gitUser

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/getUser/{username}", getUser).Methods("GET")
	router.HandleFunc("/api/getUserCl/{username}", getUserCl).Methods("GET")
	http.ListenAndServe(":4000", router)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var raw = "https://api.github.com/users/_username"
	newUrl = strings.Replace(raw, "_username", params["username"], -1)

	u, err := url.Parse(newUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(newUrl)

	resp, err2 := http.Get(u.String())
	if err2 != nil {
		log.Fatal(err)
	} else {
		var userData gitUser
		data, _ := ioutil.ReadAll(resp.Body)
		err3 := json.Unmarshal(data, &userData)
		if err3 != nil {
			fmt.Println(err3.Error)
		}
		json.NewEncoder(w).Encode(userData)
		return
	}

}

func getUserCl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var raw = "https://api.github.com/users/_username"
	newUrl = strings.Replace(raw, "_username", params["username"], -1)

	u, err := url.Parse(newUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(newUrl)

	resp, err2 := http.Get(u.String())
	if err2 != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		err3 := json.Unmarshal(data, &userData)
		if err3 != nil {
			fmt.Fprint(w, err3.Error)
		}
	}
	fmt.Fprintln(w, "Username:", userData.Login)
	fmt.Fprintln(w, "Name:", userData.Name)
	fmt.Fprintln(w, "Repos:", userData.PublicRepos)
	fmt.Fprintln(w, "URL:", userData.URL)

}

type gitUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
