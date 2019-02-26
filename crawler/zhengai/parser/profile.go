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

	if len(matches) < 9 {
		fmt.Println("Profile Parse Error:", len(matches), matches)
	}
	for _, m := range matches {
		fmt.Printf("%s\n", m[1])
	}
	profile.Marriage = string(matches[0][1])
	age, _ := strconv.Atoi(strings.Trim(string(matches[1][1]), "岁"))
	profile.Age = age
	profile.XinZou = string(matches[2][1])
	height, _ := strconv.Atoi(strings.Trim(string(matches[3][1]), "cm"))
	profile.Height = height
	weight, _ := strconv.Atoi(strings.Trim(string(matches[4][1]), "kg"))
	profile.Weight = weight
	profile.Occupation = string(matches[5][1])
	profile.Income = string(matches[6][1])
	if len(matches) == 8 {
		profile.Education = string(matches[7][1])
	} else {
		profile.Occupation = string(matches[7][1])
		profile.Education = string(matches[8][1])
	}

	fmt.Println(profile)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}
