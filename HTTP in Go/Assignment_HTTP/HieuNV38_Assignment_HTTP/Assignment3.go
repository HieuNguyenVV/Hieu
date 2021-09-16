package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	requestid "github.com/sumit-tembe/gin-requestid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/http2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var store = sessions.NewCookieStore([]byte("mysession"))
var sessionn *mgo.Session
var connectionError error
var tpl *template.Template
var count = 0
var Qstion []Data

const (
	CONST_PORT   = ":443"
	CONST_HOST   = "localhost"
	MONGO_DB_URL = "127.0.0.1"
)

type User struct {
	Pass string `json:"pass"`
	Name string `json:"name"`
}
type QS struct {
	Id         int    `json:"id"`
	Selection1 string `json:"selection1"`
	Selection2 string `json:"selection2"`
	Selection3 string `json:"selection3"`
	Selection4 string `json:"selection4"`
	Selector   string `json:"selector"`
}

type Data struct {
	Id     int    `json:"id"`
	QS     string `json:"qs"`
	As1    string `json:"as1"`
	As2    string `json:"as2"`
	As3    string `json:"as3"`
	As4    string `json:"as4"`
	Option Option
	Type   int `json:"type"`
}
type Option struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

//Encrypt password
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, 10)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//Read the list of questions in the database
func Readdata() {
	collection := sessionn.DB("mydb").C("Qs2")
	err := collection.Find(bson.M{}).All(&Qstion)
	if err != nil {
		log.Print("error occurred while reading documents from database :: ", err)
		return
	}
}

//Database connection
func init() {
	sessionn, connectionError = mgo.Dial(MONGO_DB_URL)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	sessionn.SetMode(mgo.Monotonic, true)
}

//Get the login form
func Index(c *gin.Context) {
	//Xoa database
	collection := sessionn.DB("mydb").C("UserAnser")
	for i := 0; i < len(Qstion); i++ {
		removeErr := collection.Remove(bson.M{"id": i})
		if removeErr != nil {
			log.Print("error removing document from database :: ", removeErr)
		}
	}
	//Get the login form
	tpl.ExecuteTemplate(c.Writer, "index.html", nil)
}

//Get the registration form
func Index2(c *gin.Context) {
	//Get the registration form
	tpl.ExecuteTemplate(c.Writer, "index1.html", nil)
}

//register a new account
func registration(c *gin.Context) {
	//Get data from Form data
	c.Request.ParseForm()
	name := c.Request.FormValue("username")
	pass := c.Request.FormValue("password")
	//password and empty name
	if name == "" && pass == "" {
		data := map[string]interface{}{
			"err": "Invalid",
		}
		tpl.ExecuteTemplate(c.Writer, "index1.html", data)
	} else {
		//Encrypt password
		pass1 := getHash([]byte(pass))

		var user []User
		var user1 User
		user1.Pass = pass1
		user1.Name = name
		//Read data from the DB, check if the registered account is the same as the one stored in the database
		collection := sessionn.DB("mydb").C("Users")
		err := collection.Find(bson.M{"name": name}).All(&user)
		if err != nil {
			log.Print("error occurred while reading documents from database :: ", err)
			return
		}
		var check = 0
		//check if the registered account is the same as the one stored in the database
		for i := 0; i < len(user); i++ {
			if name == user[i].Name {
				data := map[string]interface{}{
					"err": "Username already used",
				}
				check = 1
				tpl.ExecuteTemplate(c.Writer, "index1.html", data)
				break
			}
		}
		//Save the new account in the db, call the signed form successfully
		if check == 0 {
			connectionError = collection.Insert(&user1)
			tpl.ExecuteTemplate(c.Writer, "rgss.html", nil)
		}
	}
}

//Log in to the system
func login(c *gin.Context) {
	//Get data from Form data
	c.Request.ParseForm()
	name := c.Request.FormValue("username")
	pass := c.Request.FormValue("password")
	var user []User
	//Read data from the DB
	collection := sessionn.DB("mydb").C("Users")
	err := collection.Find(bson.M{"name": name}).All(&user)
	if err != nil {
		log.Print("error occurred while reading documents from database :: ", err)
		return
	}
	fmt.Println(user)
	//Check if the account is valid
	for i := 0; i < len(user); i++ {
		if name == user[i].Name {
			//check password once encrypted
			passErr := bcrypt.CompareHashAndPassword([]byte(user[i].Pass), []byte(pass))
			if passErr == nil {
				http.Redirect(c.Writer, c.Request, "/account/quiz1", http.StatusSeeOther)
			}
		}
	}
	data := map[string]interface{}{
		"err": "Incorrect username or password. Please re-entering",
	}
	count = 0
	tpl.ExecuteTemplate(c.Writer, "index.html", data)
}

//Get quiz form
func Quiz1(c *gin.Context) {
	data := map[string]interface{}{
		"t": len(Qstion),
	}
	tpl.ExecuteTemplate(c.Writer, "quiz.html", data)
}

//Get the first question
func Welcom(c *gin.Context) {
	//Put the question in the quiz form
	data := map[string]interface{}{
		"check1":   "",
		"check2":   "",
		"check3":   "",
		"check4":   "",
		"t":        len(Qstion),
		"ID":       count + 1,
		"question": Qstion[count].QS,
		"A":        Qstion[count].As1,
		"B":        Qstion[count].As2,
		"C":        Qstion[count].As3,
		"D":        Qstion[count].As4,
	}
	//Questions with one or more answers
	if Qstion[0].Type == 1 {
		//One Answer
		tpl.ExecuteTemplate(c.Writer, "quiz2.html", data)
	} else {
		//more answers
		tpl.ExecuteTemplate(c.Writer, "quiz1.html", data)
	}

}

//Nect and back button handling when doing quiz
func Quiz(c *gin.Context) {
	//Get the check status of the previous question
	c.Request.ParseForm()
	select1 := c.Request.FormValue("A")
	select2 := c.Request.FormValue("B")
	select3 := c.Request.FormValue("C")
	select4 := c.Request.FormValue("D")
	collection := sessionn.DB("mydb").C("UserAnser")
	isNext := c.Request.FormValue("isNext")
	selector := c.Request.FormValue("selector")
	var qs1 QS
	qs1.Id = count
	qs1.Selection1 = select1
	qs1.Selection2 = select2
	qs1.Selection3 = select3
	qs1.Selection4 = select4
	qs1.Selector = selector
	//Check if the question has been done, if done, update the latest status
	_, err := collection.Upsert(bson.M{"id": count}, &qs1)
	if err != nil {
		log.Print("error occurred while updating record in database :: ", err)
		//if not, save the state to db
		connectionError = collection.Insert(&qs1)
	}
	//The isNect variable stores the state of pressing the nec
	if isNext == "1" {
		count++
	}
	if isNext == "0" {
		count--
	}
	//Check the status of the next question, show the previous check of the question
	if count >= 0 && count < len(Qstion) {
		var qs2 QS
		err = collection.Find(bson.M{"id": count}).One(&qs2)
		if err != nil {
			log.Print("error occurred while  read in database :: ", err)
			//return
		}
		//check status of the next question
		checkA, checkB, checkC, checkD := false, false, false, false
		if qs2.Selector == Qstion[count].As1 {
			checkA = true
		}
		if qs2.Selector == Qstion[count].As2 {
			checkB = true
		}
		if qs2.Selector == Qstion[count].As3 {
			checkC = true
		}
		if qs2.Selector == Qstion[count].As4 {
			checkD = true
		}
		if qs2.Selection1 != "" {
			checkA = true
		}
		if qs2.Selection2 != "" {
			checkB = true
		}
		if qs2.Selection3 != "" {
			checkC = true
		}
		if qs2.Selection4 != "" {
			checkD = true
		}
		//Put the question and check status in the next question form
		data := map[string]interface{}{
			"checkA":   checkA,
			"checkB":   checkB,
			"checkC":   checkC,
			"checkD":   checkD,
			"ID":       count + 1,
			"t":        len(Qstion),
			"question": Qstion[count].QS,
			"A":        Qstion[count].As1,
			"B":        Qstion[count].As2,
			"C":        Qstion[count].As3,
			"D":        Qstion[count].As4,
		}
		////Questions with one or more answers
		if Qstion[count].Type == 1 {
			tpl.ExecuteTemplate(c.Writer, "quiz2.html", data)
		} else {
			tpl.ExecuteTemplate(c.Writer, "quiz1.html", data)
		}
	} else {
		//When back to sentence number 0
		if count < 0 {
			var qs2 QS
			err = collection.Find(bson.M{"id": 0}).One(&qs2)
			if err != nil {
				log.Print("error occurred while  read in database :: ", err)
			}
			//check status of the next question
			checkA, checkB, checkC, checkD := false, false, false, false
			if qs2.Selection1 != "" {
				checkA = true
			}
			if qs2.Selection2 != "" {
				checkB = true
			}
			if qs2.Selection3 != "" {
				checkC = true
			}
			if qs2.Selection4 != "" {
				checkD = true
			}
			if qs2.Selector == Qstion[0].As1 {
				checkA = true
			}
			if qs2.Selector == Qstion[0].As2 {
				checkB = true
			}
			if qs2.Selector == Qstion[0].As3 {
				checkC = true
			}
			if qs2.Selector == Qstion[0].As4 {
				checkD = true
			}
			//Put the question and check status in the next question form
			data := map[string]interface{}{
				"checkA":   checkA,
				"checkB":   checkB,
				"checkC":   checkC,
				"checkD":   checkD,
				"ID":       1,
				"t":        len(Qstion),
				"question": Qstion[0].QS,
				"A":        Qstion[0].As1,
				"B":        Qstion[0].As2,
				"C":        Qstion[0].As3,
				"D":        Qstion[0].As4,
				"err":      "Sentence number =0, can't back anymore",
			}
			////Questions with one or more answers
			if Qstion[0].Type == 1 {
				tpl.ExecuteTemplate(c.Writer, "quiz2.html", data)
			} else {
				tpl.ExecuteTemplate(c.Writer, "quiz1.html", data)
			}
			count = 0
		} else {
			//When nect to sentence number max
			var qs2 QS
			err = collection.Find(bson.M{"id": len(Qstion) - 1}).One(&qs2)
			if err != nil {
				log.Print("error occurred while  read in database :: ", err)
			}
			//check status of the next question
			checkA, checkB, checkC, checkD := false, false, false, false
			if qs2.Selection1 != "" {
				checkA = true
			}
			if qs2.Selection2 != "" {
				checkB = true
			}
			if qs2.Selection3 != "" {
				checkC = true
			}
			if qs2.Selection4 != "" {
				checkD = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As1 {
				checkA = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As2 {
				checkB = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As3 {
				checkC = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As4 {
				checkD = true
			}
			//Put the question and check status in the next question form
			data := map[string]interface{}{
				"checkA":   checkA,
				"checkB":   checkB,
				"checkC":   checkC,
				"checkD":   checkD,
				"t":        len(Qstion),
				"ID":       len(Qstion),
				"question": Qstion[len(Qstion)-1].QS,
				"A":        Qstion[len(Qstion)-1].As1,
				"B":        Qstion[len(Qstion)-1].As2,
				"C":        Qstion[len(Qstion)-1].As3,
				"D":        Qstion[len(Qstion)-1].As4,
				"err":      "Sentence number max, can't nect anymore",
			}
			////Questions with one or more answers
			if Qstion[len(Qstion)-1].Type == 1 {
				tpl.ExecuteTemplate(c.Writer, "quiz2.html", data)
			} else {
				tpl.ExecuteTemplate(c.Writer, "quiz1.html", data)
			}
			count = len(Qstion) - 1
		}
	}

}

//Get form review
func Review(c *gin.Context) {
	collection := sessionn.DB("mydb").C("UserAnser")
	var score = 0
	var qs2 []QS
	err := collection.Find(bson.M{}).All(&qs2)
	if err != nil {
		log.Print("error occurred while  read in database :: ", err)
	}
	//Calculate the score of the quiz
	//Score more answer question
	for i := 0; i < len(Qstion); i++ {
		for j := 0; j < len(qs2); j++ {
			if Qstion[i].Type == 0 {
				if Qstion[i].Option.A == qs2[j].Selection1 && Qstion[i].Option.B == qs2[j].Selection2 && Qstion[i].Option.C == qs2[j].Selection3 && Qstion[i].Option.D == qs2[j].Selection4 && Qstion[i].Id == qs2[j].Id {
					score++
				}
			}
		}
	}
	//Score one answer question
	for i := 0; i < len(Qstion); i++ {
		for j := 0; j < len(qs2); j++ {
			if Qstion[i].Type == 1 {
				if Qstion[i].Option.A == qs2[j].Selector && Qstion[i].Id == qs2[j].Id {
					score++
					break
				}
			}

		}
	}
	//Display score on new form
	data := map[string]interface{}{
		"score": score,
		"t":     len(Qstion),
	}
	tpl.ExecuteTemplate(c.Writer, "end.html", data)
}

//Nect and back button handling when review
func Quizreview(c *gin.Context) {
	//Get data from Form data
	c.Request.ParseForm()
	collection := sessionn.DB("mydb").C("UserAnser")
	isNext := c.Request.FormValue("isNext")
	//The isNect variable stores the state of pressing the nec
	if isNext == "1" {
		count++
	}
	if isNext == "0" {
		count--
	}
	//Check the status of the next question
	if count >= 0 && count < len(Qstion) {
		var qs2 QS
		err := collection.Find(bson.M{"id": count}).One(&qs2)
		if err != nil {
			log.Print("error occurred while  read in database :: ", err)
		}
		//Check the right or wrong answer
		var check = false
		for i := 0; i < len(Qstion); i++ {
			if Qstion[i].Type == 0 {
				if Qstion[i].Option.A == qs2.Selection1 && Qstion[i].Option.B == qs2.Selection2 && Qstion[i].Option.C == qs2.Selection3 && Qstion[i].Option.D == qs2.Selection4 && Qstion[i].Id == qs2.Id {
					check = true
				}
			}
		}
		for i := 0; i < len(Qstion); i++ {

			if Qstion[i].Type == 1 {
				if Qstion[i].Option.A == qs2.Selector && Qstion[i].Id == qs2.Id {
					check = true
					break
				}
			}
		}
		//check status of the next question
		checkA, checkB, checkC, checkD := false, false, false, false
		if qs2.Selector == Qstion[count].As1 {
			checkA = true
		}
		if qs2.Selector == Qstion[count].As2 {
			checkB = true
		}
		if qs2.Selector == Qstion[count].As3 {
			checkC = true
		}
		if qs2.Selector == Qstion[count].As4 {
			checkD = true
		}
		if qs2.Selection1 != "" {
			checkA = true
		}
		if qs2.Selection2 != "" {
			checkB = true
		}
		if qs2.Selection3 != "" {
			checkC = true
		}
		if qs2.Selection4 != "" {
			checkD = true
		}
		var option = "Incorect"
		if check == true {
			option = "Corect"
		}
		//Put the question and check status in the next question form
		data := map[string]interface{}{
			"checkA":   checkA,
			"checkB":   checkB,
			"checkC":   checkC,
			"checkD":   checkD,
			"ID":       count + 1,
			"t":        len(Qstion),
			"question": Qstion[count].QS,
			"A":        Qstion[count].As1,
			"B":        Qstion[count].As2,
			"C":        Qstion[count].As3,
			"D":        Qstion[count].As4,
			"err":      option,
		}
		////Questions with one or more answers
		if Qstion[count].Type == 1 {
			tpl.ExecuteTemplate(c.Writer, "review1.html", data)
		} else {
			tpl.ExecuteTemplate(c.Writer, "review2.html", data)
		}
	} else {
		if count < 0 {
			var qs2 QS
			err := collection.Find(bson.M{"id": 0}).One(&qs2)
			if err != nil {
				log.Print("error occurred while  read in database :: ", err)
			}
			//check status of the next question
			checkA, checkB, checkC, checkD := false, false, false, false
			if qs2.Selection1 != "" {
				checkA = true
			}
			if qs2.Selection2 != "" {
				checkB = true
			}
			if qs2.Selection3 != "" {
				checkC = true
			}
			if qs2.Selection4 != "" {
				checkD = true
			}
			if qs2.Selector == Qstion[0].As1 {
				checkA = true
			}
			if qs2.Selector == Qstion[0].As2 {
				checkB = true
			}
			if qs2.Selector == Qstion[0].As3 {
				checkC = true
			}
			if qs2.Selector == Qstion[0].As4 {
				checkD = true
			}
			//Check the right or wrong answer
			var check = false
			for i := 0; i < len(Qstion); i++ {
				if Qstion[i].Type == 0 {
					if Qstion[i].Option.A == qs2.Selection1 && Qstion[i].Option.B == qs2.Selection2 && Qstion[i].Option.C == qs2.Selection3 && Qstion[i].Option.D == qs2.Selection4 && Qstion[i].Id == qs2.Id {
						check = true
					}
				}
			}
			for i := 0; i < len(Qstion); i++ {

				if Qstion[i].Type == 1 {
					if Qstion[i].Option.A == qs2.Selector && Qstion[i].Id == qs2.Id {
						check = true
						break
					}
				}
			}
			var option = "Incorect"
			if check == true {
				option = "Corect"
			}
			//Put the question and check status in the next question form
			data := map[string]interface{}{
				"checkA":   checkA,
				"checkB":   checkB,
				"checkC":   checkC,
				"checkD":   checkD,
				"ID":       1,
				"t":        len(Qstion),
				"question": Qstion[0].QS,
				"A":        Qstion[0].As1,
				"B":        Qstion[0].As2,
				"C":        Qstion[0].As3,
				"D":        Qstion[0].As4,
				"err":      option,
			}
			////Questions with one or more answers
			if Qstion[0].Type == 1 {
				tpl.ExecuteTemplate(c.Writer, "review1.html", data)
			} else {
				tpl.ExecuteTemplate(c.Writer, "review2.html", data)
			}
			count = 0
		} else {
			var qs2 QS
			err := collection.Find(bson.M{"id": len(Qstion) - 1}).One(&qs2)
			if err != nil {
				log.Print("error occurred while  read in database :: ", err)
			}
			//check status of the next question
			checkA, checkB, checkC, checkD := false, false, false, false
			if qs2.Selection1 != "" {
				checkA = true
			}
			if qs2.Selection2 != "" {
				checkB = true
			}
			if qs2.Selection3 != "" {
				checkC = true
			}
			if qs2.Selection4 != "" {
				checkD = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As1 {
				checkA = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As2 {
				checkB = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As3 {
				checkC = true
			}
			if qs2.Selector == Qstion[len(Qstion)-1].As4 {
				checkD = true
			}
			//Check the right or wrong answer
			var check = false
			for i := 0; i < len(Qstion); i++ {
				if Qstion[i].Type == 0 {
					if Qstion[i].Option.A == qs2.Selection1 && Qstion[i].Option.B == qs2.Selection2 && Qstion[i].Option.C == qs2.Selection3 && Qstion[i].Option.D == qs2.Selection4 && Qstion[i].Id == qs2.Id {
						check = true
					}
				}
			}
			for i := 0; i < len(Qstion); i++ {

				if Qstion[i].Type == 1 {
					if Qstion[i].Option.A == qs2.Selector && Qstion[i].Id == qs2.Id {
						check = true
						break
					}
				}
			}
			var option = "Incorect"
			if check == true {
				option = "Corect"
			}
			//Put the question and check status in the next question form
			data := map[string]interface{}{
				"checkA":   checkA,
				"checkB":   checkB,
				"checkC":   checkC,
				"checkD":   checkD,
				"t":        len(Qstion),
				"ID":       len(Qstion),
				"question": Qstion[len(Qstion)-1].QS,
				"A":        Qstion[len(Qstion)-1].As1,
				"B":        Qstion[len(Qstion)-1].As2,
				"C":        Qstion[len(Qstion)-1].As3,
				"D":        Qstion[len(Qstion)-1].As4,
				"err":      option,
			}
			////Questions with one or more answers
			if Qstion[len(Qstion)-1].Type == 1 {
				tpl.ExecuteTemplate(c.Writer, "review1.html", data)
			} else {
				tpl.ExecuteTemplate(c.Writer, "review2.html", data)
			}
			count = len(Qstion) - 1
		}

	}
}

func main() {
	//Call a function that reads data from db
	Readdata()
	//Using Gin-Gonic
	router := gin.Default()
	var HttpServer = http.Server{
		Addr:    CONST_PORT,
		Handler: router,
	}
	//http2
	var Http2Server = http2.Server{}
	_ = http2.ConfigureServer(&HttpServer, &Http2Server)
	//Create log file
	f, _ := os.Create("serverAsm3.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//middleware
	{
		router.Use(gin.Recovery())
		router.Use(requestid.RequestID(nil))
		router.Use(gin.LoggerWithConfig(requestid.GetLoggerConfig(nil, nil, nil)))
	}
	//Template loads all html files in the folder
	tpl, _ = template.ParseGlob("TEMPLATE/*.html")
	//API
	router.GET("/account", Index)
	router.GET("/registration", Index2)
	router.POST("/account/login", login)
	router.POST("/account/registration", registration)
	router.GET("/account/welcom", Welcom)
	router.GET("/account/quiz1", Quiz1)
	router.POST("/account/quiz", Quiz)
	router.POST("/account/review", Review)
	router.POST("/account/reviewNect", Quizreview)
	log.Printf("Go Backend: { HTTPVersion = 2 }; serving on https://localhost:443/account")
	log.Fatal(HttpServer.ListenAndServeTLS("./server.crt", "./server.key"))
}
