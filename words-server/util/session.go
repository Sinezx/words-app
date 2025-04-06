package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type funcMap map[string]func()
type pointMap map[string]any

var sessionFuncMap = make(map[string]funcMap)
var sessionPointMap = make(map[string]pointMap)

type session struct {
	session sessions.Session
	fm      funcMap
	pm      pointMap
}

type Session interface {
	ID() string
	Get(key any) any
	Set(key any, val any)
	GetFunc(key string) func()
	SetFunc(key string, val func())
	GetPoint(key string) any
	SetPoint(key string, val any)
	Save() error
	Clear()
	DB() *gorm.DB
}

func (s *session) ID() string {
	return s.session.ID()
}

func (s *session) Get(key any) any {
	return s.session.Get(key)
}

func (s *session) Set(key any, val any) {
	s.session.Set(key, val)
}

func (s *session) GetFunc(key string) func() {
	return s.fm[key]
}

func (s *session) SetFunc(key string, val func()) {
	s.fm[key] = val
}

func (s *session) GetPoint(key string) any {
	return s.pm[key]
}

func (s *session) SetPoint(key string, val any) {
	s.pm[key] = val
}

func (s *session) Save() error {
	return s.session.Save()
}

func (s *session) Clear() {
	s.session.Clear()
	for k := range s.fm {
		delete(s.fm, k)
	}
	for k := range s.pm {
		delete(s.pm, k)
	}
}

func (s *session) DB() *gorm.DB {
	return s.pm[DataBase].(*gorm.DB)
}

func GetSession(c *gin.Context) Session {
	s := sessions.Default(c)
	fm := initFuncMap(s.ID())
	pm := initPointMap(s.ID())
	return &session{session: s, fm: fm, pm: pm}
}

func initPointMap(id string) pointMap {
	pm := sessionPointMap[id]
	if pm == nil {
		pm = make(pointMap)
		sessionPointMap[id] = pm
	}
	return pm
}

func initFuncMap(id string) funcMap {
	fm := sessionFuncMap[id]
	if fm == nil {
		fm = make(funcMap)
		sessionFuncMap[id] = fm
	}
	return fm
}
