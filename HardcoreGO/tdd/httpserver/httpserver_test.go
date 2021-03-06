package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string] int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string){
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T){

	store := StubPlayerStore{
		map[string] int{
			"Pepper" : 20,
			"Floyd"  : 10,
		},
		nil,
	}
	server := &PlayerServer{ &store}

	t.Run("return Pepper's score", func(t *testing.T){
		request  := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		gotStatusCode  := response.Code
		wantStatusCode := http.StatusOK

		assertStatusCode(t, gotStatusCode, wantStatusCode)
		assertResponseBody(t, response.Body.String(), "20")
	})
	
	t.Run("return Floyd's score", func(t *testing.T) {
		request  := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		gotStatusCode  := response.Code
		wantStatusCode := http.StatusOK

		assertStatusCode(t, gotStatusCode, wantStatusCode)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request  := newGetScoreRequest("lakshay")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		gotStatusCode  := response.Code
		wantStatusCode := http.StatusNotFound

		assertStatusCode(t, gotStatusCode, wantStatusCode)
	})
}

func TestScoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := &PlayerServer{&store}

	t.Run("it record wins when post", func(t *testing.T){
		player := "Pepper"

		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsAndRetrieveThem(t *testing.T){
	store := InMemoryPlayerStore{ map[string] int {}}
	server := PlayerServer{&store}

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")

}
func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string ){
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatusCode(t * testing.T, got, want int){
	t.Helper()

	if got != want {
		t.Errorf("did not gget correct statusm got %d, want %d", got, want)
	}
}