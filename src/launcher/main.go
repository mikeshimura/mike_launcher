package main

import (
	"fmt"
	"launcher/util"
	"os"
	"os/exec"
	"strings"

	"github.com/go-ini/ini"
	filepath "path/filepath"
)

func main() {
	curDir, _ := filepath.Abs(".")
	fmt.Printf("cur %v\n", curDir)

	if len(os.Getenv("TEST"))==0{
		exe, _ := os.Executable()
		dir:=filepath.Dir(exe)
		os.Chdir(dir)
	}
	argx := strings.Split(os.Args[0], string(os.PathSeparator))
	sexe := strings.Split(argx[len(argx)-1], ".")

	fDir, _ := filepath.Abs(".")
	fmt.Printf("finalpath %v\n", fDir)
	sini := "$" + sexe[0] + ".ini"
	if len(os.Args) >1 && os.Args[1]>""{
		sini = "$" + os.Args[1] + ".ini"
	}
	cfg, err := ini.Load(sini)
	if err != nil {
		panic(sini + "が読めません")
	}
	//http://username:password@127.0.0.1:9999
	//os.Setenv("HTTP_PROXY", "http://proxyIp:proxyPort")
	section, err := cfg.GetSection("default")
	aws := section.Key("AWS_ACCESS_KEY_ID").String()
	sec := section.Key("AWS_SECRET_ACCESS_KEY").String()
	bucket := section.Key("BUCKET").String()
	fmt.Printf("bucket %v\n", bucket)
	cmd := section.Key("CMD").String()
	zip := section.Key("ZIP").String()
	unzip := section.Key("UNZIP").String()
	watch := section.Key("WATCH").String()
	appos := section.Key("OS").String()
	his := section.Key("HIS").String()
	proxy := strings.TrimSpace(section.Key("PROXY").String())
	proxyport := strings.TrimSpace(section.Key("PROXYPORT").String())
	proxyid := strings.TrimSpace(section.Key("PROXYID").String())
	proxypass := strings.TrimSpace(section.Key("PROXYPASS").String())
	if len(proxy) > 0 && len(proxyid) > 0 {
		os.Setenv("HTTP_PROXY", "http://"+proxyid+":"+proxypass+"@"+proxy+":"+proxyport)
		fmt.Printf("Proxy :%v\n", "http://"+proxyid+":"+proxypass+"@"+proxy+":"+proxyport)
	}
	if len(proxy) > 0 && len(proxyid) == 0 {
		os.Setenv("HTTP_PROXY", "http://"+proxy+":"+proxyport)
		fmt.Printf("Proxy :%v\n", "http://"+proxy+":"+proxyport)
	}
	hismap := util.GetHis(his)
	watch = watch
	hismap = hismap
	bucket = bucket
	os.Setenv("AWS_ACCESS_KEY_ID", aws)
	os.Setenv("AWS_SECRET_ACCESS_KEY", sec)
	keys := util.ListKeys(bucket)
	for _, key := range keys {
		fmt.Printf("key %v\n", key)
	}
	newhis := util.CheckNew(zip, keys, hismap, bucket)
	fmt.Printf("newhis %v\n", newhis)
	if newhis {
		if appos != "MAC" {
			util.UnzipApp(zip)
		} else {
			cmd := exec.Command("unzip", "-o", zip)
			cmd.Run()
		}
	}
	util.CheckNew(watch, keys, hismap, bucket)
	util.CheckWatch(watch, keys, hismap, bucket)
	util.SaveHis(hismap, his)
	if len(unzip)>0 {
		if appos != "MAC" {
			util.UnzipApp(unzip)
		} else {
			cmd := exec.Command("unzip", "-o", unzip)
			cmd.Run()
		}
	}
	util.Exec(cmd)
	os.Exit(0)
}
