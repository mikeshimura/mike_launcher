package util

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
)

func Download(bucket string, skey string) {
	skeyx := strings.Replace(skey, "/", string(os.PathSeparator), -1)
	S3Download(bucket, skey, skeyx)
}
func UnzipApp(zip string) {
	err := Unzip(zip, "."+string(os.PathSeparator))
	if err != nil {
		panic(err)
	}
}
func Unzip(src string, dest string) (error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return  err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return  err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return  err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return  err
			}

		}
	}
	return  nil
}
func Exec(scmd string) {
	cols := strings.Split(scmd, " ")
	cmd := exec.Command(cols[0], cols[1:len(cols)]...)
	cmd.Run()
}
func GetHis(his string) map[string]interface{} {
	hismap := make(map[string]interface{})
	_, err := os.Stat(his)
	if err != nil {
		return hismap
	}
	shis, _ := ioutil.ReadFile(his)
	if err := json.Unmarshal(shis, &hismap); err != nil {
		panic(err)
	}
	return hismap
}
func CheckNew(skey string, keys []*s3.Object, hismap map[string]interface{}, bucket string) bool {
	fmt.Printf("skey %v\n", skey)
	for _, key := range keys {
		if *key.Key == skey {
			his := hismap[skey]
			var modtime string = (*key.LastModified).Format(time.RFC3339)
			if his == nil {
				hismap[skey] = modtime
				Download(bucket, skey)
				return true
			}
			if his.(string) != modtime {
				hismap[skey] = modtime
				Download(bucket, skey)
				return true
			}
			return false
		}
	}
	panic(skey + "がBucketに見つかりません")
}
func CheckWatch(watch string, keys []*s3.Object, hismap map[string]interface{}, bucket string) {
	swatch, err := ioutil.ReadFile(watch)
	if err != nil {
		panic(watch + "が読めません")
	}
	lines := GetLines(string(swatch))
	for _, s := range lines {
		if len(s) > 0 {
			CheckNew(s, keys, hismap, bucket)
		}
	}

}
func SaveHis(hismap map[string]interface{}, his string) {
	jsonFile, _ := os.Create(his)
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err := encoder.Encode(&hismap)
	if err != nil {
		panic("json が書き込めません")
	}
}
func GetLines(s string) []string {
	if strings.Index(s, "\r\n") > -1 {
		return strings.Split(string(s), "\r\n")
	}
	if strings.Index(s, "\r") > -1 {
		return strings.Split(string(s), "\r")
	}
	return strings.Split(string(s), "\n")
}
