package main

import "fmt"

// slice
type taskList struct{
	tasks []*task
}

func (t *taskList) add_task(ts *task){
	t.tasks = append(t.tasks, ts)
}

func (t *taskList) delete_task(index int){
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) print_list(){
	for _, task := range t.tasks{
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.description)
	}
}

func (t *taskList) print_list_completed(){
	for _, task := range t.tasks{
		if task.complete{
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.description)
		}
	}
}

type task struct{
	name string
	description string
	complete bool
}

func (t *task) mark_complete() {
	t.complete = true
}

func (t *task) update_description(new_description string)  {
	t.description = new_description
}

func (t *task) update_name(new_name string) {
	t.name = new_name
}

func main() {
	t1 := &task{
		name: "Complete go course",
		description: "Complete go course this week",
	}
	
	t2 := &task{
		name: "Complete python course",
		description: "Complete python course this week",
	}

	t3 := &task{
		name: "Complete react course",
		description: "Complete react course this week",
	}

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	list.add_task(t3)
	list.print_list_completed()

	// for i := 0; i < len(list.tasks); i++{
	// 	fmt.Println("Index", i, "Task", list.tasks[i].name)
	// }

	// for index, task := range list.tasks{
	// 	fmt.Println("Index", index, "Task", task.name)
	// }

	task_map := make(map[string]*taskList)

	task_map["Kevin"] = list
	fmt.Println("Kevin Task")
	task_map["Kevin"].print_list()

	t4 := &task{
		name: "Complete java course",
		description: "Complete java course this week",
	}

	t5 := &task{
		name: "Complete c# course",
		description: "Complete c# course this week",
	}

	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	task_map["Maria"] = list2
	fmt.Println("Maria Task")
	task_map["Maria"].print_list()

}