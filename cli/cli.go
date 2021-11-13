package cli

import (
	"bufio"
	"fmt"
	"strings"
)

func setCommandSlice() []string {
	cs := []string{"-h", "-v", "-l", "-c", "-a", "-m", "-d", "-q"}
	return cs
}
func SetHelpCommandSlice() []string {
	hcs := []string{"todo -h #help", "todo -v #version", "todo -l #list un-completed items", "todo -c #list completed items", "todo -a #add new item 'todo -a [TODO]'", "todo -m #mark as completed 'todo -m [ID]'", "todo -d #delete item 'todo -d [ID]'", "todo -q #exit program"}
	return hcs
}

func PrintSkeletonTable() {
	fmt.Println("   ID   |  Item                            | Date     ")
	fmt.Println("--------:----------------------------------:----------")
}

func ReadConsole(r *bufio.Reader) (string, string, error) {
	text, _ := r.ReadString('\n')
	tArr := strings.Split(text, " ")
	var err error

	if len(tArr) < 2 {
		err = fmt.Errorf("Missing argument, for more help 'todo -h'")
		return "", "", err
	}
	todo := strings.TrimSpace(tArr[0])
	arg := strings.TrimSpace(tArr[1])
	in := strings.TrimSpace(strings.Join(tArr[2:], " "))

	if todo != "todo" {
		err = fmt.Errorf("No argument such as '%v',for more help 'todo -h'", todo)
	} else if !contains(arg) {
		err = fmt.Errorf("No argument such as '%v %v',for more help 'todo -h'", todo, arg)
	}
	return arg, in, err
}

func contains(s string) bool {
	cs := setCommandSlice()
	for _, v := range cs {
		if s == v {
			return true
		}
	}
	return false
}
