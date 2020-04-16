package myapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type fooHandler struct{}

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name`
	Email     string    `json:"email`
	CreatedAt time.Time `json:"created_at`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request : %v", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {

	// ServeMux struct 안에 func (*ServeMux) ServeHTTP 라는 메소드가 구현되있기
	// 떄문에. http.Handler 인터페이스를 implement(구현하다) 했다고 할 수있다.
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World\n")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	return mux
}
