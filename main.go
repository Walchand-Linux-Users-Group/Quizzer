package main
import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Person struct {
	name  string
	score string
}
type api struct{
	Api string
}

func main() {
	
	res ,_:=http.Get("https://golangapi-production.up.railway.app/api")
	a,err:=	ioutil.ReadAll(res.Body)

	var GETAPI api;
	json.Unmarshal(a,&GETAPI)
	
	BASE_URL := GETAPI.Api;
	
	var name string
	url:=BASE_URL+"q"
	//fmt.Print(url);
	fmt.Print("Enter your name : ")
	nameC := make(chan string)
	go func() {
		fmt.Scanf("%s", &name)
		nameC <- name
	}()

	nam := <-nameC
	m := map[string]string{
		"name": nam,
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Print("conversion to json went wrong ")
		return
	}
	postData, err := http.Post(BASE_URL+"user", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Print("Something went wrong in post req")
	}
	
	id, _ := ioutil.ReadAll(postData.Body)
	problems := questionPuller(url)
	plen := len(problems)
	tobj := time.NewTimer(30*time.Duration(plen) * time.Second) // Time for all the questions --> 1 question => 10 seconds
	start := time.Now()

	var correctAns int = 0
ProblemLoop:

	for i, problem := range problems {
		var answer string
		// fmt.Printf(,"\nProblem %d: %s\n", i+1, problem.Question)

		fmt.Println(string("\033[37m"),"Problem ", i+1, " out of ",plen,": \n", problem.Question)
		fmt.Println(string("\033[34m"),"a. ", problem.Options.A)
		fmt.Println(string("\033[36m"),"b. ", problem.Options.B)
		fmt.Println(string("\033[33m"),"c. ", problem.Options.C)
		fmt.Println(string("\033[35m"),"d. ", problem.Options.D)
		fmt.Print(string("\033[32m"), "Enter Your Choice (a, b, ,c ,d ) : ")

		// fmt.Printf("\n a. %s \n b. %s \n c. %s \n d. %s \n Select Option 'a','b','c','d' : ",
		// 	problem.Options.A, problem.Options.B, problem.Options.C, problem.Options.D)
		ansC := make(chan string)
		reader:= bufio.NewReader(os.Stdin)
		
		go func() {
			ans,error := reader.ReadString('\n')
			// fmt.Scanf("%s", &answer)
			if error!=nil{
				fmt.Println("Error in reading file")
			}
			if ans ==" "{
				ans = "f"
			}
			answer = ans[0:1]
			
			ansC <- answer

		}()

		select {
		case <-tobj.C:
			fmt.Println("\nTime Over !!! Your Quiz has been Submitted\n")
			break ProblemLoop
		case iAns := <-ansC:
			if iAns == problem.Answer {
				correctAns = correctAns + 800
				
				
			}
			if i == len(problems)-1 {
				fmt.Print("All Questions Submitted Successfully...:)\n ")
				timeRemaining := float64(plen*30) - time.Since(start).Seconds()
				
				correctAns += int(timeRemaining)
			}
			rtime := float64(plen*30) - time.Since(start).Seconds()
			fmt.Println(string("\033[31m"),"---------:Time Remaining ", int(rtime), " Seconds:---------")


		}

	}

	m2 := map[string]string{
		"_id":   string(id),
		"score": strconv.Itoa(correctAns),
	}
	jsonData2, err := json.Marshal(m2)
	if err != nil {
		fmt.Print("conversion to json went wrong ")
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, BASE_URL+"updateUser", bytes.NewBuffer(jsonData2))
	if err != nil {
		panic(err)
	}
	// Set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	id2, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string("\033[32m"),string(id2))
	fmt.Println(string("\033[32m"),"Score is ", correctAns)
}
