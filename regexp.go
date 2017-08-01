package main

import (
	"fmt"
	"sort"
	"bytes"
	"log"
	"errors"
	"regexp"
	"strconv"
)

func findDiff(sli []string) (re []int) {
	re = make([]int, 0)
	for i, v1 := range sli {
		for j, v2 := range sli {
			if i == j {
				continue
			}

			for x := 0; x < getMin(len(v1), len(v2)); x++ {
				if v1[x] != v2[x] {
					_, e := strconv.Atoi(string(v1[x]))
					if e != nil {
						break
					}
					_, e = strconv.Atoi(string(v2[x]))
					if e != nil {
						break
					}
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
	dif := findDiff(titles)

	if len(dif) <= 0 {
		err = errors.New("wrong dif length: " + fmt.Sprint(titles))
		return
	}
	sort.Ints(dif)

	var buffer bytes.Buffer
	buffer.WriteString(regexp.QuoteMeta(titles[0][0:dif[0]]))
	buffer.WriteString("([0-9]+)")
	log.Println(buffer.String())
	re = buffer.String()
	return
}
