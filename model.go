package main

import (
	"time"
)

type Kadai struct {
	_type    string
	title    string
	course   string
	start    time.Time
	deadline time.Time
}

func NewKadai(typ, title, course string, start, deadline time.Time) Kadai {
	return Kadai{
		_type:    typ,
		title:    title,
		course:   course,
		start:    start,
		deadline: deadline,
	}
}

type KadaiGetter interface {
	GetAll() ([]Kadai, error)
}

type kadaiGetter struct{}

func (kg *kadaiGetter) GetAll() ([]Kadai, error) {
	// TODO: implement
	return nil, nil
}

func NewKadaiGetter() KadaiGetter {
	return &kadaiGetter{}
}
