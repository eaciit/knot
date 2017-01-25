package knot

import (
	"github.com/eaciit/toolkit"
	"net/http"
	"sync"
	"time"
)

type Sessions map[string]toolkit.M

var (
	sessionCookieId string
	sessions        Sessions
	sessionLocker   *sync.RWMutex = new(sync.RWMutex)
)

func SetSessionCookieId(id string) {
	sessionCookieId = id
}

func SessionCookieId() string {
	if sessionCookieId == "" {
		sessionCookieId = "KnotSessionId"
	}
	return sessionCookieId
}

func InitSessions() *Sessions {
	if sessions == nil {
		sessions = map[string]toolkit.M{}
	}
	return &sessions
}

func (s Sessions) InitTokenBucket(tokenid string) {
	if _, b := s[tokenid]; !b {
		s[tokenid] = toolkit.M{}
	}
}

func (s Sessions) Set(tokenid, key string, value interface{}) {
	s.InitTokenBucket(tokenid)

	sessionLocker.Lock()
	s[tokenid].Set(key, value)
	sessionLocker.Unlock()
}

func (s Sessions) Get(tokenid, key string, def interface{}) interface{} {
	s.InitTokenBucket(tokenid)

	sessionLocker.RLock()
	value := s[tokenid].Get(key, def)
	sessionLocker.RUnlock()

	return value
}

/** use own cookie setter.
using `knot.Cookie` will make cookie path follow the actual request path,
causing returned cookie value will always different (or nil), if accessed from different page.
because of that, fetching session value from one page to another becoming impossible */
func setCookieForSession(r *WebContext, cookieId string, tokenId string, expire time.Duration) {
	c := &http.Cookie{}
	c.Name = cookieId
	c.Value = tokenId
	c.Expires = time.Now().Add(expire)
	c.Path = "/"

	if r.cookies == nil {
		r.cookies = map[string]*http.Cookie{}
	}

	r.cookies[cookieId] = c
	http.SetCookie(r.Writer, c)
}

func getSessionTokenIdFromCookie(r *WebContext) string {
	tokenId := ""
	c, _ := r.Cookie(SessionCookieId(), "")
	if c == nil {
		tokenId = toolkit.GenerateRandomString("", 32)
		setCookieForSession(r, SessionCookieId(), tokenId, time.Hour*24*30)
	} else {
		tokenId = c.Value
	}
	return tokenId
}

func (r *WebContext) Session(key string, defs ...interface{}) interface{} {
	InitSessions()
	tokenId := getSessionTokenIdFromCookie(r)
	var def interface{}
	if len(defs) > 0 {
		def = defs[0]
	}
	return sessions.Get(tokenId, key, def)
}

func (r *WebContext) SetSession(key string, value interface{}) {
	InitSessions()
	tokenId := getSessionTokenIdFromCookie(r)
	sessions.Set(tokenId, key, value)
}
