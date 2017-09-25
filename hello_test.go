package hello

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/appengine-guestbook-go/testdata"
	"google.golang.org/appengine/aetest"
	// "gopkg.in/jarcoal/httpmock.v1"
)

var inst aetest.Instance

func TestMain(m *testing.M) {
	inst, _ = aetest.NewInstance(nil)

	m.Run()

	defer tearDown()
}

func TestHandler(t *testing.T) {
	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Cant create request")
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(recorder, req)

	expected := testdata.FRONTEND

	if strings.TrimSpace(recorder.Body.String()) != expected {
		t.Errorf("Handler returned a different frontend than expected. Expected: %v, got %v", recorder.Body.String(), expected)
	}
}

func TestParseData(t *testing.T) {
	m := []string{"Python", "C++", "Javascript"}

	user1 := User{
		Login:         "omeyjey",
		Contributions: 69,
	}

	user2 := User{
		Login:         "henrik",
		Contributions: 22,
	}

	cons := Contributors{
		Users: []User{user1, user2},
	}

	own := Owner{
		Login: "omeyjey",
	}

	repo1 := Repo{
		Project:      "git/git",
		Owner:        own,
		Contributors: "",
		Languages:    "",
	}

	data := parseData(cons, m, repo1)

	if data.Project != "github.com/git/git" {
		t.Fatalf("Wrong project name")
	}

	if data.Owner != "omeyjey" {
		t.Fatalf("Wrong owner")
	}
}

// func TestGetContributors(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", "/testgetcommits",
// 		httpmock.NewStringResponder(200, testdata.COMMITTERS))
//
// 	req, err := inst.NewRequest("GET", "/testgetcommits", nil)
// 	if err != nil {
// 		t.Fatal("Cant create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	comsAPI := getData(recorder, req, "/testgetcommits")
// 	comitters := getContributors(comsAPI)
//
// 	com := Contributors{}
// 	com.Users = []User{
// 		{
// 			Login:         "klyve",
// 			Contributions: 276,
// 		},
// 		{
// 			Login:         "Hanssen97",
// 			Contributions: 104,
// 		},
// 		{
// 			Login:         "omeyjey",
// 			Contributions: 57,
// 		},
// 		{
// 			Login:         "henriktre",
// 			Contributions: 12,
// 		},
// 	}
//
// 	expected := com
//
// 	if !reflect.DeepEqual(expected, comitters) {
// 		t.Fail()
// 	}
// }

// func TestGetLanguages(t *testing.T) {
//
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", "/testgetlangs",
// 		httpmock.NewStringResponder(200, `{"Javascript": 69, "C++": 2}`))
//
// 	req, err := inst.NewRequest("GET", "/testgetlangs", nil)
// 	if err != nil {
// 		t.Fatal("Cant create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	langsAPI := getData(recorder, req, "/testgetlangs")
// 	langs := getLanguages(langsAPI)
//
// 	var m []string
//
// 	for k := range langs {
// 		m = append(m, k)
// 	}
//
// 	expected := []string{"Javascript", "C++"}
//
// 	if !reflect.DeepEqual(m, expected) {
// 		t.Fail()
// 	}
// }
//
// func TestFailGetLanguages(t *testing.T) {
//
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", "/testgetlangs",
// 		httpmock.NewStringResponder(200, `{"PHP"}`))
//
// 	req, err := inst.NewRequest("GET", "/testgetlangs", nil)
// 	if err != nil {
// 		t.Fatal("Cant create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	langsAPI := getData(recorder, req, "/testget")
//
// 	if langsAPI != nil {
// 		t.Fail()
// 	}
// }

// func TestHandlerRepo(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()
//
// 	httpmock.RegisterResponder("GET", "https://api.github.com/repos/omeyjey/database_assigment1",
// 		httpmock.NewStringResponder(200, testdata.INITIALDATA))
//
// 	httpmock.RegisterResponder("GET", "https://api.github.com/repos/omeyjey/database_assigment1/contributors",
// 		httpmock.NewStringResponder(200, testdata.CONTRIBUTORS))
//
// 	httpmock.RegisterResponder("GET", "https://api.github.com/repos/omeyjey/database_assigment1/languages",
// 		httpmock.NewStringResponder(200, testdata.LANGUAGES))
//
//
// 	req, err := inst.NewRequest("GET", "/projectinfo/v1/github.com/omeyjey/database_assigment1", nil)
// 	if err != nil {
// 		t.Fatal("Cant create request")
// 	}
//
// 	recorder := httptest.NewRecorder()
//
// 	handler := http.HandlerFunc(handlerRepo)
// 	handler.ServeHTTP(recorder, req)
//
// 	expected := testdata.REPODATA
//
// 	if strings.TrimSpace(recorder.Body.String()) != expected {
// 		t.Errorf("Handler returned a different frontend than expected. Expected: %v, got %v", recorder.Body.String(), expected)
// 	}
// }
//
func tearDown() {
	if inst != nil {
		inst.Close()
	}
}
