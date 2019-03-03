package parser

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

const profileRe = `<div class="m-btn purple"[^>]+>([^<]+)</div>`

func ParseProfile(contents []byte, name string) engine.ParseResult  {
	profile := model.Profile{Name:name}
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		//fmt.Printf("%s\n", m[1])
		if strings.Index(string(m[1]), "岁") != -1 {
			age, _ := strconv.Atoi(strings.Trim(string(m[1]), "岁"))
			profile.Age = age
		}
		if strings.Index(string(m[1]), "cm") != -1 {
			height, _ := strconv.Atoi(strings.Trim(string(matches[3][1]), "cm"))
			profile.Height = height
		}
		if strings.Index(string(m[1]), "kg") != -1 {
			weight, _ := strconv.Atoi(strings.Trim(string(matches[4][1]), "kg"))
			profile.Weight = weight
		}
		if strings.Index(string(m[1]), "座(") != -1 {
			profile.XinZou = string(m[1])
		}
		if strings.Index(string(m[1]), "月收入:") != -1 {
			profile.Income = string(m[1])
		}
		if strings.Index(string(m[1]), "工作地:") != -1 {
			profile.Occupation = string(m[1])
		}
	}


	fmt.Println(profile)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}
