package testing_exce

import "testing"

func TestStore(t *testing.T) {
	//创建moster实例
	monster := Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	res := (&monster).Store()
	if !res {
		t.Fatalf("monster.Store()错误，期望值为=%v,实际值为%v", true, res)
	}
	t.Logf("monster.Store()测试成功！")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore()错误，期望值为=%v,实际值为%v", true, res)
	}
	//进一步判断
	if (*monster).Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误，期望值为=%v,实际值为%v", "qinzhibao", (*monster).Name)
	}

	t.Logf("monster.Store()测试成功！")
}
