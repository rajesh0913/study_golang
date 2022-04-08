package main

import (
	"go_code/chapter14_unit_test/demo02_testing_exce/testing_exce"
)

func main() {
	var monster testing_exce.Monster = testing_exce.Monster{
		Name:  "qinzhibao",
		Age:   25,
		Skill: "type",
	}
	monster.Store()
}
