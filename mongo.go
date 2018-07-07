package goutil

import "github.com/globalsign/mgo"

func DialMongoDB(mongodbUrl string) *mgo.Database {
	di, err := mgo.ParseURL(mongodbUrl)
	PanicIfErr(err)
	se, err := mgo.DialWithInfo(di)
	PanicIfErr(err)
	return se.DB("")
}