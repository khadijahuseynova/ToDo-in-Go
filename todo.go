package todo
import(
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)
//We create 2 data structures struct item and list []item

type item struct{
	Task	string
	Done    bool
	CreatedAt	time.Time
	CompletedAt	time.Time
}

type List []item

//Our application should add new tasks,mark them as completed,delete tasks,save the list and get the tasks.
//First we will start with Add Function

func (l *List) Add(task string) {
	t:=item{
		Task:	task,
		Done:	false,
		CreatedAt:	time.Now(),
		CompletedAt:	time.Time{},
	}

	*l = append(*l, t)
}

//Complete Function

func(l *List) Complete(i int) error {
	ls:=*l
	if i<=0 || i>len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done=true
	ls[i-1].CompletedAt= time.Now()

	return nil
}

//Save Function
func(l *List)  Save(filename string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, json, 0644)
}

//Get Function

func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		if errors.Is(err, os.ErrNotExist) {
			return nil
	}
	return err
}

	if len(file)==0 {
		return nil
}
	return json.Unmarshal(file, l)
}

	
