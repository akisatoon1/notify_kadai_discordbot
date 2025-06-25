package main

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/akisatoon1/manaba"
)

type KadaiGetter interface {
	GetAll() ([]Kadai, error)
}

type kadaiGetter struct {
	studentID string
	password  string
	cookie    *cookiejar.Jar
}

func NewKadaiGetter(studentID, password string) KadaiGetter {
	return &kadaiGetter{
		studentID: studentID,
		password:  password,
	}
}

func (kg *kadaiGetter) GetAll() ([]Kadai, error) {
	err := kg.login()
	if err != nil {
		return nil, err
	}

	kadais, err := kg.getKadaiList()
	if err != nil {
		return nil, err
	}

	return kadais, nil
}

func (kg *kadaiGetter) login() error {
	if kg.cookie == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			return err
		}
		kg.cookie = jar
	}
	return manaba.Login(kg.cookie, kg.studentID, kg.password)
}

func (kg *kadaiGetter) getKadaiList() ([]Kadai, error) {
	// httpリクエストのため
	client := &http.Client{
		Jar: kg.cookie,
	}

	listURL := "https://room.chuo-u.ac.jp/ct/home_library_query"
	resp, err := client.Get(listURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch kadai list: " + resp.Status)
	}

	// スクレイピングのため
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var kadais []Kadai
	err = nil // 中の関数でエラーが出るときのため

	selector := "tr.row0,.row1"
	doc.Find(selector).Each(
		func(_ int, s *goquery.Selection) {
			tds := s.Find("td")
			_type := cutSpace(tds.Eq(0).Text())
			title := cutSpace(tds.Eq(1).Text())
			course := cutSpace(tds.Eq(2).Text())

			var start, deadline *time.Time

			if startStr := cutSpace(tds.Eq(3).Text()); startStr != "" {
				start = new(time.Time)
				*start, err = time.Parse("2006-01-02 15:04", startStr)
			}

			if deadlineStr := cutSpace(tds.Eq(4).Text()); deadlineStr != "" {
				deadline = new(time.Time)
				*deadline, err = time.Parse("2006-01-02 15:04", deadlineStr)
			}

			k := NewKadai(_type, title, course, start, deadline)
			kadais = append(kadais, k)
		},
	)
	if err != nil {
		return nil, err
	}

	return kadais, nil
}

// 前後の空白を削除する関数
func cutSpace(s string) string {
	return strings.TrimSpace(s)
}
