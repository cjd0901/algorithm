package _021_05

// Employee Importance
// leetcode: https://leetcode-cn.com/problems/employee-importance/
// Now given the employee information of a company, and an employee id, you need to return the total importance value of this employee and all their subordinates.

type Employee struct {
    Id int
    Importance int
    Subordinates []int
}

func EmployeeImportance2(employees []*Employee, id int) int {
    mp := map[int]*Employee{}
    for _, employee := range employees {
        mp[employee.Id] = employee
    }
    total := 0
    var dfs func(int)
    dfs = func(id int) {
        employee := mp[id]
        total += employee.Importance
        for _, subId := range employee.Subordinates {
            dfs(subId)
        }
    }
    dfs(id)
    return total
}