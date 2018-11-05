package goutil

import (
	"github.com/globalsign/mgo"
	"strings"
)

func DialMongoDB(mongodbUrl string) *mgo.Database {
	if strings.Contains(mongodbUrl, "@") {
		di, err := mgo.ParseURL(mongodbUrl)
		PanicIfErr(err)
		se, err := mgo.DialWithInfo(di)
		PanicIfErr(err)
		return se.DB("")
	}
	se, err := mgo.Dial(mongodbUrl)
	PanicIfErr(err)
	return se.DB("")
}
