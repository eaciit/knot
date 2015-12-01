package knot

import (
	"github.com/eaciit/toolkit"
	"time"
)

type Sessions map[string]toolkit.M

var (
	sessionCookieId string
	sessions        Sessions
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
	s[tokenid].Set(key, value)
}

func (s Sessions) Get(tokenid, key string, def interface{}) interface{} {
	s.InitTokenBucket(tokenid)
	return s[tokenid].Get(key, def)
}

func getSessionTokenIdFromCookie(r *WebContext) string {
	tokenId := ""
	c, _ := r.Cookie(SessionCookieId(), "")
	if c == nil {
		tokenId = toolkit.GenerateRandomString("", 32)
		r.SetCookie(SessionCookieId(), tokenId, time.Hour*24*30)
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
