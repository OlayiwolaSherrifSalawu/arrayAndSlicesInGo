package main

import "fmt"

type UserPresence map[string]map[string]string

func (m UserPresence) AddUser(userId string) error {
	// no two user must have the same id
	_, ok := m[userId]
	if ok {
		return fmt.Errorf("cannot add existing  user.")
	}
	m[userId] = map[string]string{
		"theme": "ligth",
	}
	return nil
}

func (m UserPresence) AddPrensence(userId string, prefrence string, value string) {
	m[userId] = map[string]string{
		prefrence: value,
	}
}
