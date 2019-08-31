package count_bean

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type money float64
type Crash struct {
	Money      money
	Name       string
	crashCount map[string]Crash
}

var st_crash Crash

type MyItems struct {
	beans money
	itime time.Time

	toer  string //给谁的
	notes string //注释

	from string //从那个账号来的
	to   string //到哪个账号去

	itype  int
	status int //对订单是否表示确认 ! *
}

//2018-07-26 * "小桔科技" "滴滴打车"
//把所有全部切成一块一块符合要求的 记账表
func (t *Crash) split2Items(str []string) [][]string {
	rst := [][]string{}

	assert_check := func(a string) bool {
		return true
	}
	for index, _ := range str {
		if strings.TrimSpace(str[index]) == "" {
			if assert_check(str[0]) {
				tmp := str[:index]
				rst = append(rst, tmp)
			}
		}
	}
	return rst
}

func (t *Crash) splitSpace(str, split_str string) (rst []string) {
	if len(str) < 0 {
		return
	}
	for _, value := range strings.Split(str, split_str) {
		if len(strings.TrimSpace(value)) > 0 {
			rst = append(rst, value)
		}
	}
	return
}

func (t *Crash) buildCrash(str_arr []string) (rst Crash) {
	rst.crashCount = make(map[string]Crash)
	ct := rst.crashCount
	for i := 0; i < len(str_arr); i++ {
		crashName := str_arr[i]
		ct[crashName] = Crash{crashCount: make(map[string]Crash)}
		//ct[crashName].crashCount =
		ct = ct[crashName].crashCount
		//ct = make(map[string]Crash)
	}

	return
}

func (t *Crash) createByName(name string) (rst []string) {
	rst_arr := t.splitSpace(name, ":")
	rst = rst_arr
	return
}

func (t *MyItems) checkIsSame(a MyItems) bool {
	return t.itype == a.itype
}

func (t *MyItems) splitSpace(str, split_str string) (rst []string) {
	if len(str) < 0 {
		return
	}
	for _, value := range strings.Split(str, split_str) {
		if len(strings.TrimSpace(value)) > 0 {
			rst = append(rst, value)
		}
	}
	return
}

func (t *MyItems) changeTime(str string) time.Time {
	var timeLayoutStr = "2006-01-02"
	st, err := time.ParseInLocation(timeLayoutStr, str, time.Local)
	if err != nil {
		// TODO: 还需要处理一下
	}
	return st
}
func (t *MyItems) changeStatus(str string) (rst int) {
	switch str {
	case "*":
		rst = 0
	case "!":
		rst = 1
	}
	return
}

// 2018-07-26 * 小桔科技 滴滴打车
//     Assets:Cash                 -15.00 CNY
//     Expenses:Traffic:Didi       15.00 CNY
func (t *MyItems) tr2Item(its []string) (rst MyItems) {
	if len(its) != 3 {
		return
	}
	its0 := t.splitSpace(its[0], " ")
	if len(its0) >= 3 {
		rst.itime = t.changeTime(its0[0])
		rst.status = t.changeStatus(its0[1])
		rst.notes = its0[2]
	} else {
		return
	}
	if len(its0) >= 4 {
		rst.toer = its0[2]
		rst.notes = its0[3]
	}

	//处理转账信息
	its1 := t.splitSpace(its[1], " ")
	if len(its1) >= 2 {
		rst.from = strings.TrimSpace(its1[0])
		imonery, err := strconv.ParseFloat(its1[1], 64)
		if err != nil {
		}
		rst.beans = money(math.Abs(imonery))
	} else {
		return
	}

	//到哪里去可以省略不写 金额
	its2 := t.splitSpace(its[2], " ")
	if len(its2) > 0 {
		rst.to = strings.TrimSpace(its2[0])
		//TODO: 可以给个警告 如果2个数字不一样说明有问题
	}

	return
}
