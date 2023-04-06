//Test for functions
package todo_test

import (
    "io/ioutil"
    "os"
    "testing"

   todo "github.com/khadijahuseynova/ToDoinGo/parent"
)

//Test for add function

func TestAdd(t *testing.T) {
    l := todo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
    }
}
//Test for complete function
func TestComplete(t *testing.T) {
    l := todo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
    }

    if l[0].Done {
        t.Errorf("New task should not be completed.")
    }
    l.Complete(1)

    if !l[0].Done {
        t.Errorf("New task should be completed.")
    }
}

//Test for Save and Get Functions
func TestSaveGet(t *testing.T) {
    // two list
    l1 := todo.List{}
    l2 := todo.List{}
    taskName := "New Task"
    // saving task into l1 and loading it into l2 -- error if fails
    l1.Add(taskName)
    if l1[0].Task != taskName {
        t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
    }
    tf, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatalf("Error creating temp file: %s", err)
    }
    defer os.Remove(tf.Name())
    if err := l1.Save(tf.Name()); err != nil {
        t.Fatalf("Error saving list to file: %s", err)
    }
    if err := l2.Get(tf.Name()); err != nil {
        t.Fatalf("Error getting list from file: %s", err)
    }
    if l1[0].Task != l2[0].Task {
        t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
    }
}

