package hello

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Joker/jade"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/gorilla/mux"
)

// User containing data about the user
type User struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

// Contributors data about contributions
type Contributors struct {
	Users []User
}

// Owner containing data about the owner
type Owner struct {
	Login string `json:"login"`
}

// Repo struct containing data about repo
type Repo struct {
	Project      string `json:"full_name"`
	Owner        Owner  `json:"owner"`
	Contributors string `json:"contributors_url"`
	Languages    string `json:"languages_url"`
}

// Data struct containing all the data
type Data struct {
	Project   string   `json:"project"`
	Owner     string   `json:"owner"`
	Committer string   `json:"comitter"`
	Commits   int      `json:"commits"`
	Language  []string `json:"language"`
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/projectinfo/v1/github.com/{username}/{repo}", handlerRepo)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	parseJade(w, r)

	parseErr := r.ParseForm()
	if parseErr != nil {
		fmt.Printf("\nExecute error: %v", parseErr)
		return
	}

	userName := strings.Join(r.Form["userName"], "")
	repoName := strings.Join(r.Form["repoName"], "")
	str := "/projectinfo/v1/github.com/" + userName + "/" + repoName

	if userName != "" && repoName != "" {
		http.Redirect(w, r, str, http.StatusMovedPermanently)
	}
}

func handlerRepo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "json; charset=utf-8")

	link := "https://api.github.com/repos/" + vars["username"] + "/" + vars["repo"]

	if vars["username"] == "" && vars["repo"] == "" {
		link = "https://api.github.com/repos/omeyjey/database_assigment1"
	}

	body := getData(w, r, link)

	repo1 := Repo{}

	jsonError := json.Unmarshal(body, &repo1)
	if jsonError != nil {
		log.Fatal(jsonError)
	}

	consAPI := getData(w, r, repo1.Contributors)
	langsAPI := getData(w, r, repo1.Languages)

	cons := getContributors(consAPI)
	langs := getLanguages(langsAPI)

	var m []string

	for k := range langs {
		m = append(m, k)
	}

	data := parseData(cons, m, repo1)

	output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "", string(output))
}

func getContributors(body []byte) Contributors {
	cons := Contributors{}
	jsonError := json.Unmarshal(body, &cons.Users)

	if jsonError != nil {
		log.Fatalf("Unmarshal error: %v", jsonError)
	}
	return cons
}

func getLanguages(body []byte) map[string]interface{} {
	var langs map[string]interface{}
	jsonError := json.Unmarshal(body, &langs)

	if jsonError != nil {
		fmt.Println(jsonError)
	}
	return langs
}

func parseJade(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("view/index.jade")
	if err != nil {
		fmt.Printf("\nReadFile error: %v", err)
		return
	}

	jadeTpl, err := jade.Parse("jade_tp", string(buf))
	if err != nil {
		fmt.Printf("\nParse error: %v", err)
		return
	}

	goTpl, err := template.New("html").Parse(jadeTpl)
	if err != nil {
		fmt.Printf("\nTemplate parse error: %v", err)
		return
	}

	err = goTpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("\nExecute error: %v", err)
		return
	}
}

func getData(w http.ResponseWriter, r *http.Request, link string) []byte {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(link)
	if err != nil {
		panic(err.Error())
		return nil
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil
	}
	resp.Body.Close()

	return body
}

func parseData(cons Contributors, m []string, repo1 Repo) Data {
	data := Data{}

	data.Project = "github.com/" + string(repo1.Project)
	data.Owner = repo1.Owner.Login
	data.Committer = cons.Users[0].Login
	data.Commits = cons.Users[0].Contributions
	data.Language = m

	return data
}
