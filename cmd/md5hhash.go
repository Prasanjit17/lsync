package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/syndtr/goleveldb/leveldb"
	"io"
	"os"
	"strings"
	//"github.com/syndtr/goleveldb"

)


type (
	KV struct {
		File string
		Hash string
	}
)



func CheckHash(file string) ([]byte, error){

	//db, err := database.New()
	//if err != nil {
	//	color.Red("ERROR_DB_CONN: %s", err)
	//	os.Exit(2)
	//}

	db, err := leveldb.OpenFile("dbms-db", nil)
	if err != nil {
		return nil, err
	}

	//defer db.Close()

	b, e := db.Get([]byte(file), nil)
	color.Red("ERROR_IN_LDB: %s",e)
	if e != nil {
		if strings.Contains(e.Error(), "not found") {
			hash, err := HashMd(file)
			if err != nil {
				fmt.Printf("** Error in generating Hash** \n %s\n", err)
				return nil, err
			}
			err = db.Put([]byte(file), hash,nil)
			if err != nil {
				fmt.Printf("%s in while data is PUT to DB \n", err)
				return nil, err
			}
			return nil, e
		}
	}

	return b, e
}

func HashMd(file string) ([]byte, error) {
	var data []KV

	h := md5.New()
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error in opening the file:%v because %s", file, err)
		return nil, err
	}
	_, err = io.Copy(h, f)
	if err != nil {
		fmt.Printf("ERROR_BYTE_COPY: %s\n", err)
		return nil, err
	}

	hash := h.Sum(nil)
	returnMD5 := hex.EncodeToString(hash)
	data = append(data, KV{File: file, Hash: returnMD5})

	return json.Marshal(data)
}


