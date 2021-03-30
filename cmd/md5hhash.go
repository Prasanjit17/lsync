package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	//"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/fatih/color"
	//"github.com/golang/leveldb"
	"io"
	"lsync/database"
	"os"
	"strings"

)


type (
	KV struct {
		File string
		Hash string
	}
)



func CheckHash(file string) ([]byte, error){

	//db, err := leveldb.Open("dbms-db", nil)
	//if err != nil {
	//	fmt.Println("ld")
	//	return nil, err
	//}
	db, _ := database.New()
	//defer db.Close()

	b, e := db.Get(file)
	if e != nil {
		if strings.Contains(e.Error(), "not found") {
			hash, err := HashMd(file)
			if err != nil {
				fmt.Printf("** Error in generating Hash** \n %s\n", err)
				return nil, err
			}
			err = db.Set(file, hash)
			if err != nil {
				fmt.Printf("%s in while data is PUT to DB \n", err)
				return nil, err
			}
			return nil, err
		}
	} else {
		color.Red("File hash found so same file can not be uploded")
		os.Exit(3)
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


