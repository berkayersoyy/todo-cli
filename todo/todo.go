package todo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	cli "todo/cli"
)

type Todo struct {
	Id        int
	Item      string
	Date      time.Time
	Completed *bool
}

type TodoSlice []Todo

func (t TodoSlice) PrintItems() {
	cli.PrintSkeletonTable()
	for _, v := range t {
		if !(*v.Completed) {
			b := "                             "
			r := len(b) - len(v.Item)
			if !(r < 0) {
				b = b[:r+1]
			}
			fmt.Println("   ", v.Id, "  | ", v.Item, b, "|", CalculateTimeDifference(v.Date))
		}
	}
}

func (t TodoSlice) PrintCompletedItems() {
	cli.PrintSkeletonTable()
	for _, v := range t {
		if *v.Completed {
			b := "                             "
			r := len(b) - len(v.Item)
			if !(r < 0) {
				b = b[:r+1]
			}
			fmt.Println("   ", v.Id, "  | ", v.Item, b, "|", CalculateTimeDifference(v.Date))
		}
	}
}
func GetTodoWithId(t *TodoSlice, id string) (*Todo, error) {
	cid, err := strconv.Atoi(id)
	var todo *Todo
	if err != nil {
		err = fmt.Errorf("Invalid input '%v'", id)
	}
	for _, v := range *t {
		if v.Id == cid {
			todo = &v
			return todo, err
		}
	}
	err = fmt.Errorf("No such as todo found with id '%v'", id)
	return todo, err
}
func MarkItemAsCompleted(t *TodoSlice, id string) error {
	todo, err := GetTodoWithId(t, id)
	if err != nil {
		return err
	}
	*todo.Completed = true
	return nil
}
func AddItem(t *TodoSlice, id *int, in string) error {
	var err error
	if len(strings.TrimSpace(in)) == 0 {
		err = fmt.Errorf("Cannot add empty todo")
		return err
	}
	td := Todo{Id: *id, Item: in, Date: time.Now(), Completed: NewFalse()}
	*t = append(*t, td)
	*id = *id + 1
	return err
}
func DeleteItem(t *TodoSlice, id string) error {
	cid, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("Invalid input '%v'", id)
	}
	for i, v := range *t {
		if v.Id == cid {
			Delete(t, i)
			return nil
		}
	}
	return fmt.Errorf("No such as todo found with id '%v'", id)
}
func Delete(t *TodoSlice, item int) {
	s := *t
	s = append(s[:item], s[item+1:]...)
	*t = s
}
func NewFalse() *bool {
	b := false
	return &b
}
func CalculateTimeDifference(t time.Time) string {
	d := time.Now()
	a := daysBetween(d, t)
	if a == 0 {
		return "Today"
	} else if a > 0 && a < 7 {
		v := fmt.Sprint(a)
		var vr string
		if a < 2 {
			vr = v + " day ago"
			return vr
		}
		vr = v + " days ago"
		return vr
	} else if a >= 7 && a <= 30 {
		w := a / 7
		v := fmt.Sprint(w)
		var vr string
		if w < 2 {
			vr = v + " week ago"
			return vr
		}
		vr = v + " weeks ago"
		return vr
	} else if a > 30 && a < 365 {
		m := a / 30
		v := fmt.Sprint(m)
		var vr string
		if m < 2 {
			vr = v + " month ago"
			return vr
		}
		vr = v + " months ago"
		return vr
	} else {
		y := a / 365
		v := fmt.Sprint(y)
		var vr string
		if y < 2 {
			vr = v + " year ago"
			return vr
		}
		vr = v + " years ago"
		return vr
	}

}
func daysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}
