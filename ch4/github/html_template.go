package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

/*
The text/template and html/template packages provide a mechanism
for substituting the values of variables into a text or html template.

A template is a string or file containing one or more portions enclosed in
double braces {{...}}, called actions.
Each action contains an expression in the template langiage, a simple but powerful notation for printing values, selecting struct fields, calling
functions and methods, expressing control flow such as if-else statements
and range loops, and instantiating other templates.
*/

const templ = `{{.TotalCount}} issues:
	{{range .Items}} --------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt |  daysAgo}} days
	{{end}}`

// Within an action, there is a notion of the current value, referred to
// the template's parameter, which will be a github.IssuesSearchResult from // the previous json example.
// The {{.TotalCount}} action expands to the value of the TotalCount filed.
// The {{range .Items}} and {{end}} actions create a loop, so the text
// between them is expanded multiple times, with dot bound to successive
// elements of Items.
// Within an action the | notation makes the result of one operation the
// argument of another analogous to pipe of unix shell.
// The printf is a built-in synonym of fmt.Sprintf in all templates.
// The daysSgo function converts the CreatedAt field into an elapsed
// time, using time.Since:

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Notice that the type of CreatedAt is time.Time, not string. In
// the same way that a type may control its string formatting by defining
// certain methods, a type may also define methods to control its JSON
// marshaling and unmarshaling behavior. The JSON-marshaled value of a
// time.Time is a string in a standard format.

// Producing output with a template is a 2 step process.
// First, we parse the template into a suitable internal representation,
// and then execute in on specific inputs.
// Parsing need be done only once. The code below creates and parses the
// template temp1 defined above.

// report, err := template.New("report").
//  	Funcs(template.FuncMap{"daysAgo": daysAgo}).
//  	Parse(templ)

// Note the chaining of method calls:
// template.New creates and returns a template
// Funcs add daysAgo to the set of functions accessible within this template
// finally Parse is called on the result.

// Because templates are usually fixed t compile time, failure to parse a
// template indicates a fatal bug in the program
// The template.Must helper function makes error handling more convenient.
// It accepts a template and an error, checks that the error is nil
// and then returns the template.
var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

// Coming to html/template package. It uses the same API and expression
// language as text/template but adds features for automatic and
// context-appropriate escaping of string appearing within HTML, JavaScript
// CSS or URLs.
// These feature can help avoid a prennial security problem of HTML
// generation, an injection attack in which an adversary crafts a string
// value like the title of an issue to include malicious code that
// when improperly escaped by a template, gives them control over the page.

var issueList = template.Must(template.New("issuelist").Parse(`
     <h1>{{.TotalCount}} issues</h1>
     <table>
     <tr style='text-align: left'>
       <th>#</th>
       <th>State</th>
       <th>User</th>
       <th>Title</th>
     </tr>
     {{range .Items}}
     <tr>
       <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
       <td>{{.State}}</td>
       <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
       <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
     </tr>
     {{end}}
     </table>
     `))

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	// The html/template package automatically HTML-escaped the titles so that
	// they appear literally.

	fmt.Print("End of Main")
}
