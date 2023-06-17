package main

import (
	"GoStudy/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

// 模板
const temol = ` {{.TotalCount}} issues:          
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

const hello = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
	Hello, {{.Name}}!
	Age: {{.Age}}
	Email: {{.Email}}
</body>
</html>
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// 创建模板
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temol))
var helloTempl = template.Must(template.New("hello").Parse(hello))

func main() {
	//result, err := github.SearchIssues(os.Args[1:])
	//if err != nil {
	//	log.Fatal(err)
	//}
	lhs := github.Person{
		Name:  "LHS",
		Age:   24,
		Email: "12345678@qq.com",
	}
	if err := helloTempl.Execute(os.Stdout, lhs); err != nil {
		log.Fatal(err)
	}

}
