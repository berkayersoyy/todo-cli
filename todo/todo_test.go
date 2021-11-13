package todo

import (
	"testing"
	"time"
)

func InitializeTodoSlice() TodoSlice {
	ts := TodoSlice{}
	td := Todo{Id: 1, Item: "Buy milk", Date: time.Now(), Completed: NewFalse()}
	td2 := Todo{Id: 2, Item: "Buy water", Date: time.Now(), Completed: NewFalse()}
	td3 := Todo{Id: 3, Item: "Buy apple", Date: time.Now(), Completed: NewFalse()}
	td4 := Todo{Id: 4, Item: "Buy banana", Date: time.Now(), Completed: NewFalse()}
	td5 := Todo{Id: 5, Item: "Buy potato", Date: time.Now(), Completed: NewFalse()}
	ts = append(ts, td, td2, td3, td4, td5)
	return ts
}

func TestGetTodoWithId(t *testing.T) {
	ts := InitializeTodoSlice()
	td, err := GetTodoWithId(&ts, "1")
	if err != nil {
		t.Error(err)
	}
	if td.Completed != ts[0].Completed && td.Date != ts[0].Date && td.Id != ts[0].Id && td.Item != ts[0].Item {
		t.Errorf("TestGetTodoWithId: Some struct properties are not equal!")
	}
}
func TestMarkItemAsCompleted(t *testing.T) {
	ts := InitializeTodoSlice()
	err := MarkItemAsCompleted(&ts, "1")
	if err != nil {
		t.Error(err)
	}
	if ts[0].Completed == NewFalse() {
		t.Errorf("TestMarkItemAsCompleted: got: %v, want: true", *(ts[0].Completed))
	}
}
func TestAddItem(t *testing.T) {
	ts := InitializeTodoSlice()
	l := len(ts) + 1
	id := l
	AddItem(&ts, &id, "Buy chocolate")
	if len(ts) != l {
		t.Errorf("TestAddItem: length of TodoSlice got: %v, want: %v", len(ts), l)
	}
}
func TestDeleteItem(t *testing.T) {
	ts := InitializeTodoSlice()
	l := len(ts) - 1
	id := "1"
	err := DeleteItem(&ts, id)
	if err != nil {
		t.Error(err)
	}
	if len(ts) != l {
		t.Errorf("TestDeleteItem: length of TodoSlice got: %v, want: %v", len(ts), l)
	}
}
func TestCalculateTimeDifference(t *testing.T) {
	td := time.Now().Add(-24 * time.Hour)
	r := CalculateTimeDifference(td)
	w := "1 day ago"
	if r != w {
		t.Errorf("TestCalculateTimeDifference: got: %v, want: %v", r, w)
	}
}
