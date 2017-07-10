package sesseion

import (
	"html/template"
	"net/http"
	"testing"
	"time"
)

func TestSessionManager(t *testing.T) {
}

func login(w http.ResponseWriter, r *http.Request) {
	session := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, session.Get("username"))
	} else {
		session.Set("username", r.FormValue("usename"))
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	session := globalSessions.SessionStart(w, r)
	createTime := session.Get("createtime")
	if createTime == nil {
		session.Set("createtime", time.Now().Unix())
	} else if (createTime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		session = globalSessions.SessionStart(w, r)
	}
	count := session.Get("countnum")
	if count == nil {
		session.Set("countnum", 1)
	} else {
		session.Set("countnum", (count.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-type", "text/html")
	t.Execute(w, session.Get("countnum"))
}
