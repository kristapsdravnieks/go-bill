package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//User is used for all customers and users
type User struct { //some
	Number         int
	FirstName      string
	LastName       string
	Mail           string
	Address        string
	PostNrAndPlace string
	PhoneNr        string
	OrgNr          string
	CountryID      string
}

var pDB *sql.DB                        //The pointer to use with the Database
var tmpl map[string]*template.Template //map to hold all templates
var indexNR int                        //to store the index nr. in slice where chosen person is stored

func init() {
	//initate the templates
	tmpl = make(map[string]*template.Template)
	tmpl["init.html"] = template.Must(template.ParseFiles("templates.html"))
}

func main() {
	//create DB and store pointer in pDB
	pDB = createDB()
	defer pDB.Close()

	http.HandleFunc("/sp", showUsersWeb)
	http.HandleFunc("/ap", addUsersWeb)
	http.HandleFunc("/mp", modifyUsersWeb)
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/du", deleteUserWeb)
	http.HandleFunc("/createBill", billCreateWeb)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":7000", nil)

}
