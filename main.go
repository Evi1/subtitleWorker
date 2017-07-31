package main

import (
	"regexp"
	"log"
	"strconv"
	"sort"
	"fmt"
)

func findIntInSlice(sli []int, obj int) (re int) {
	re = -1
	for i, v := range sli {
		if v == obj {
			re = i
			break
		}
	}
	return
}

func findDiff(sli []string) (re []int, err error) {
	re = make([]int, 0)
	for i, v1 := range sli {
		log.Println(strconv.Itoa(i) + ":" + v1)
		for j, v2 := range sli {
			if i == j {
				continue
			}
			if len(v1) != len(v2) {
				err = error("wrong with length: "+v1+" & "+v2)
				return
			}
			for x := 0; x < len(v1); x++ {
				if v1[x] != v2[x] {
					if findIntInSlice(re, x) < 0 {
						re = append(re, x)
					}
				}
			}
		}
	}
	return
}

func getReg(titles []string) (re string, err error) {
	dif, e := findDiff(titles)
	if e != nil {
		err = e
		return
	}
	if len(dif) <= 0 {
		err = error("wrong titles: "+fmt.Sprint(titles))
		return
	}
	sort.Ints(dif)
	//check if the num
	for i := 0; i < len(dif)-1; i++ {
	}
	return
}

func main() {
	str1 := "Onimonogatari - 01.Chs&Jap"
	log.Println(string(str1[0]))
	reg := regexp.MustCompile(`Onimonogatari - ([0-9]+?).Chs&Jap`)
	str := reg.FindAllStringSubmatch(str1, -1)
	log.Println(str)
	var sli = make([]string, 3)
	sli[0] = "Onimonogatari - 01.Chs&Jap"
	sli[1] = "Onimonogatari - 02.Chs&Jap"
	sli[2] = "Onimonogatari - 13.Chs&Jap"
	log.Println(getReg(sli))
}
