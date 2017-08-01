package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"path"
	"os"
)

func handleFile(p string) {
	//get all videos and subs
	log.Println("Handle " + p)
	files, _ := ioutil.ReadDir(p)
	videoList := make(map[string][]string)
	subList := make(map[string][]string)
	for _, f := range files {
		if !f.IsDir() {
			tp := checkSuffix(f.Name(), videoSuffix)
			if len(tp) <= 0 {
				tp = checkSuffix(f.Name(), subSuffix)
				if len(tp) > 0 {
					subList[tp] = append(subList[tp], f.Name())
				}
			} else {
				videoList[tp] = append(videoList[tp], f.Name())
			}
		}
	}
	//get the key videos and subs
	keyVType := ""
	keySType := ""
	num := 0
	for k := range videoList {
		if len(videoList[k]) > num {
			keyVType = k
			num = len(videoList)
		}
	}
	num = 0
	for k := range subList {
		if len(subList[k]) > num {
			keySType = k
			num = len(subList)
		}
	}
	vr, err := getReg(videoList[keyVType])
	if err != nil {
		log.Println(err)
		log.Println("skip " + p + " because of error")
		return
	}
	sr, err := getReg(subList[keySType])
	if err != nil {
		log.Println(err)
		log.Println("skip " + p + " because of error")
		return
	}
	vreg := regexp.MustCompile(vr)
	sreg := regexp.MustCompile(sr)
	for _, vName := range videoList[keyVType] {
		re1 := vreg.FindAllStringSubmatch(vName, -1)
		for _, sName := range subList[keySType] {
			re2 := sreg.FindAllStringSubmatch(sName, -1)

			if len(re1) <= 0 || len(re2) <= 0 {
				log.Println("skip " + p + " found error with reg")
				return
			}
			if re1[0][1] == re2[0][1] {
				fileSuffix := path.Ext(sName) //获取文件后缀
				var filenameOnly string
				filenameOnly = strings.TrimSuffix(vName, path.Ext(vName)) //获取文件名
				log.Println(vName, sName, filenameOnly+fileSuffix)
				os.Rename(p+"/"+sName, p+"/"+filenameOnly+fileSuffix)
				break
			}
		}
	}
}

func handleFolder(path string) {
	fileHandled := false
	files, _ := ioutil.ReadDir(path)
	log.Println("In " + path)
	for _, f := range files {
		if f.IsDir() {
			handleFolder(path + "/" + f.Name())
		} else if !fileHandled {
			handleFile(path)
			fileHandled = true
		}
	}
}
