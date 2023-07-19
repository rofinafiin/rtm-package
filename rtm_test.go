package rtm_package

import (
	"fmt"
	"github.com/aiteung/atdb"
	"os"
	"testing"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "HelpdeskData",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	id := "xx2"
	Nama := "Nafiis"
	Email := "contoh@gmail.com"
	Handphone := "0000000"
	hasil := InsertDataUser(MongoConn, id, Nama, Email, Handphone)
	fmt.Println(hasil)

}

func TestGetDatauser(t *testing.T) {
	id := "xx2"
	hasil := GetDatauser(id, MongoConn, "data_user")
	fmt.Println(hasil)

}

func TestGetDatabyphone(t *testing.T) {
	Handphone := "0000000"
	hasil := GetDataUserFromPhone(Handphone, MongoConn, "data_user")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Handphone := "0000000"
	hasil := DeleteData(Handphone, MongoConn, "data_user")
	fmt.Println(hasil)

}
