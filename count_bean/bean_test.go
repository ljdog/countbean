package count_bean

import (
	"strings"
	"testing"
	"time"
)

//;收入(Income)”、“支出(Expenses)”、“负债(Liabilities)”、“资产(Assets)”
func Test_Sample(t *testing.T) {
	assert_check := func(got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("want %#v, got %#v", want, got)
		}
	}
	t.Run("创建测试", func(t *testing.T) {
		//got := Wallet()
		got := Crash{Money: 10, Name: "资产"}
		want := Crash{}
		want.Name = "资产"
		want.Money = 10
		assert_check(got, want)
	})
	t.Run("金钱条目", func(t *testing.T) {
		got := MyItems{beans: 12, itime: time.Now(), notes: "买了一个大西瓜", itype: 1}
		want := MyItems{beans: 12, itime: got.itime, notes: "买了一个大西瓜", itype: 1}
		assert_check(got, want)
	})
}

func Test_splitstring(t *testing.T) {
	assert_check := func(got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("\r\n want %#v, \r\n got %#v", want, got)
		}
	}
	// t.Run("check readfile", func(t *testing.T) {
	// 	data, err := ioutil.ReadFile("../bean_test")
	// 	if err != nil {
	// 		fmt.Println("File reading error", err)
	// 		return
	// 	}
	// 	got := string(data)
	// 	want := `2018-07-26 * 小桔科技 滴滴打车
	// 	Assets:Cash                 -15.00 CNY
	// 	Expenses:Traffic:Didi       15.00 CNY`

	// 	assert_check(got, want)
	// })

	t.Run("check tr2Item", func(t *testing.T) {
		got_str := `2018-07-26  *  小桔科技    滴滴打车
		Assets:Cash                       -15.00 CNY
		Expenses:Traffic:Didi       15.00 CNY`
		got_arr := strings.Split(got_str, "\n")

		change_var := MyItems{}
		got := change_var.tr2Item(got_arr)
		want := MyItems{
			itime:  time.Date(2018, 07, 26, 0, 0, 0, 0, time.Local),
			status: 0,
			beans:  15,
			toer:   "小桔科技",
			notes:  "滴滴打车",
			from:   "Assets:Cash",
			to:     "Expenses:Traffic:Didi",
		}
		assert_check(got, want)
	})

	t.Run("check tr2Item1", func(t *testing.T) {
		got_str := `2018-07-26  *    滴滴打车
		Assets:Cash                       -15.00 CNY
		Expenses:Traffic:Didi       15.00 CNY`
		got_arr := strings.Split(got_str, "\n")
		change_var := MyItems{}
		got := change_var.tr2Item(got_arr)
		want := MyItems{
			itime:  time.Date(2018, 07, 26, 0, 0, 0, 0, time.Local),
			status: 0,
			beans:  15,
			toer:   "",
			notes:  "滴滴打车",
			from:   "Assets:Cash",
			to:     "Expenses:Traffic:Didi",
		}
		assert_check(got, want)
	})

	t.Run("check tr2Item1", func(t *testing.T) {
		got_str := `2018-07-26  *    滴滴打车
		   Assets:Cash                       -15.00 CNY
		    Expenses:Traffic:Didi           15.00 CNY`
		got_arr := strings.Split(got_str, "\n")
		change_var := MyItems{}
		got := change_var.tr2Item(got_arr)
		want := MyItems{
			itime:  time.Date(2018, 07, 26, 0, 0, 0, 0, time.Local),
			status: 0,
			beans:  15,
			toer:   "",
			notes:  "滴滴打车",
			from:   "Assets:Cash",
			to:     "Expenses:Traffic:Didi",
		}
		assert_check(got, want)
	})
}

func map_check(a, b Crash) bool {
	if len(a.crashCount) != len(b.crashCount) {
		return false
	}
	for k, v := range a.crashCount {
		if !map_check(v, b.crashCount[k]) {
			return false
		}
	}
	return true
}

func Test_crash(t *testing.T) {
	assert_check := func(got, want Crash) {
		t.Helper()
		if len(got.crashCount) != len(want.crashCount) {
			t.Errorf("\r\n want %#v, \r\n got %#v", want, got)
		}
		if !map_check(got, want) {
			t.Errorf("\r\n want %#v, \r\n got %#v", want, got)
		}
	}
	stat_crash := Crash{}
	got_arr := stat_crash.createByName("a:b:c")
	got := stat_crash.buildCrash(got_arr)
	want := Crash{}
	want.crashCount = make(map[string]Crash)
	want.crashCount["a"] = Crash{crashCount: make(map[string]Crash)}
	want.crashCount["a"].crashCount["bb"] = Crash{crashCount: make(map[string]Crash)}
	want.crashCount["a"].crashCount["bb"].crashCount["c"] = Crash{crashCount: make(map[string]Crash)}

	assert_check(got, want)
}
