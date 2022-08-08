package main

import "log"

func (c *config) CountPending() int {
	total := 0

	tasks, err := c.currentTasks()
	if err != nil {
		log.Println(err)
	}
	for _, item := range tasks {
		if !item.Done {
			total++
		}
	}

	return total
}
