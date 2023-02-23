package db

import (
	"reflect"
)

type Session struct {
	Modal interface{}
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) Model(m interface{}) *Session {
	s.Modal = m
	return s
}

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				// log.Error(err)
				println(err)
			}
		}
	}
	return
}

// API

func (s *Session) Insert(value interface{}) (session *Session, e error) {
	s.CallMethod("BeforeInsert", value)
	return s, e
}

// Util

// a - b
func difference(a []string, b []string) (diff []string) {
	cacheB := make(map[string]bool)
	for _, v := range b {
		cacheB[v] = true
	}
	for _, v := range a {
		if _, ok := cacheB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}
