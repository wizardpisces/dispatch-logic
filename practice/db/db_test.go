package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	println("before inert", account.ID)
	account.ID += 1000

	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	println("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_, _ = s.Insert(&Account{1, "123456"})
}

func TestUtil(t *testing.T) {
	a := []string{"2", "3"}
	b := []string{"1", "2"}
	assert.Equal(t, difference(a, b), []string{"3"})

}
