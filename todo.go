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
	Creadted-At	time.Time
	Completed-At	time.Time
}

type List []item

//Our application should add new tasks,mark them as completed,delete tasks,save the list and get the tasks.
//First we will start with Add Function

func (a *List) Add(task string) {
	t:=item{
		Task:	task,
		Done:	false,
		Created-At	time.Now(),
		Completed-At	time.Time(),
	}

	*a = append(*a, t)
}

//Complete Function

func(a *List) Complete(i int) error {
	ls:=*a
	if i<=0 || i>len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done=true
	ls[i-1].Complete-At= time.Now()

	return nil
}

//Save Function
func(a *List)  Save(filename string) error {
	json, err := json.Marshall(a)
	if err != nil {
		return err
	}
	return ioutill.WriteFile(filename, json, 0644)
}

//Get Function

func (a *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		if errors.Is(err,os.ErrNotExists) {
			return nil
	}
	return err
}

	if len(file)==0 {
		retunr nil
}
	return json.Unmarshal(file, a)
}

	
