package main

import "fmt"

func main()  {
	prerequisites := [][]int{[]int{2 , 0}, []int{1 , 0} , []int{3,1} , []int{3,2} , []int{1,3}}
	res := canFinish(4 , prerequisites)
	fmt.Println(res)
}


func canFinish(numCourses int, prerequisites [][]int) bool {
	coursePrerequisites := make([][]int , numCourses)
	visited := make([]int , numCourses)
	for _ , tmp := range prerequisites {
		coursePrerequisites[tmp[0]] = append(coursePrerequisites[tmp[0]], tmp[1])
	}
	for i := 0 ; i < numCourses; i++ {
		if !dfs(coursePrerequisites , i , visited) {
			return false
		}
	}
	return true
}

func dfs(coursePrerequisites [][]int , course int , visited []int) bool {
	if visited[course] == 1 {
		return false
	}
	if visited[course] == 2 {
		return true
	}
	visited[course] = 1
	for _ , preCourse:= range coursePrerequisites[course] {
		if visited[preCourse] == 1 {
			return false
		}
		if visited[preCourse] == 0 {
			if !dfs(coursePrerequisites , preCourse , visited) {
				return false
			}
		}
	}
	visited[course] = 2
	return true
}
