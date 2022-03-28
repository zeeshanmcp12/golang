package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// todo app logic
/*
func main() {
	fmt.Println("New File for practice!")
	fmt.Println()


		for i := 1; i < 10+1; i++ {
			fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
		}

		fmt.Printf("Enter any number to make table: ")
		tableNum := userInput()

		for i := 1; i < 10+1; i++ {

			// anotherInt := i * 2
			fmt.Printf("%v x %v = ", tableNum, i)
			result := userInput()

			// Working logic
			if result == tableNum*i {
				fmt.Println("result executed")

			} else {
				fmt.Printf("Wrong Answer!\n%v x %v = %v\n", tableNum, i, userInput())
			}

		}

	todos := []string{}
	fmt.Printf("Add task: ")

	for i := 0; true; i++ {

		f_text := strings.TrimSpace(stringInput())

		// fmt.Printf("Add task or Enter done: ")
		if f_text != "done" {
			if f_text != "" {

				fmt.Printf("Add task or enter done: ")
				todos = append(todos, f_text)
				continue

			} else {
				fmt.Printf("Invalid input. Add task or enter done: ")

			}

		} else if f_text == "done" {
			fmt.Println()
			fmt.Println("Here is your tasks:")
			for j, val := range todos {
				fmt.Printf("%v %v\n", j+1, val)
			}
			break
		}
	}
}*/

// working with files
/*
func main() {

	fmt.Println("Working with Files in golang")
	content := "Magic! Magic! Magic!"

	file, err := os.Create("./practiceFile.txt")
	CheckNilErr(err)

	io.WriteString(file, content)

	defer file.Close()
	readFile("./practiceFile.txt")

}*/

// const url string = "https://acloudtechie.com/azure"

// Handling web request
/*
func main() {
	fmt.Println("Handling web request in golang!")

	response, err := http.Get(url)
	CheckNilErr(err)
	fmt.Printf("Type of response is %T", response)
	fmt.Println()

	// fmt.Printf("Site status: %v", string(response.Status))
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	fmt.Println(response.Header)

	content := res

	file, err := os.Create("./disclaimer.html")
	CheckNilErr(err)
	fmt.Println("File has been created and written with response.")

	io.WriteString(file, string(content))

}*/

// URL handling
/*
const customUrl string = "https://acloudtechie.com/azure?article_name=geo-restricting-web-app&publish_date=march"

func main() {
	fmt.Println("Handling URL in golang")
	fmt.Println(customUrl)

	parsedURL, err := url.Parse(customUrl)
	CheckNilErr(err)

	fmt.Printf("Type of parsedURL is %T", parsedURL)
	fmt.Println()

	// fmt.Println("URL scheme is: ", parsedURL.Scheme)
	// fmt.Println("Host is: ", parsedURL.Host)
	// fmt.Println("Hostname is: ", parsedURL.Hostname())
	// fmt.Println("Port is: ", parsedURL.Port())
	// fmt.Println("Query param is: ", parsedURL.Path)
	// fmt.Println("URL scheme is: ", parsedURL.RawQuery)

	queryParams := parsedURL.Query()
	for i, val := range queryParams {
		fmt.Println(i, val)
	}

	constructUrl := &url.URL{
		Scheme: "https",
		Host:   "acloudtechie.com",
		Path:   "/disclaimer",
	}

	fmt.Println(constructUrl.String())

}*/

func main() {
	fmt.Println("Web request verb!")

	// const siteUrl = "https://acloudtechie.com"
	// GetRequest([]byte(siteUrl))
	// CallAzureFunc()
	// PostRequest()
	PostFormRequest()

}

func GetRequest(customUrl []byte) {

	response, err := http.Get(string(customUrl))
	CheckNilErr(err)

	fmt.Println("Site status is ", response.Status)

	// to retrieve headers
	for j, val := range response.Header {
		fmt.Printf("%v %v\n", j, val)
	}
}

func PostRequest() {
	const myUrl = "http://localhost:8000/post"

	fakeJson := strings.NewReader(`
	{
		"name":"zeeshan",
		"blog":"acloudtechie.com"
	}
	`)

	response, err := http.Post(myUrl, "application/json", fakeJson)
	CheckNilErr(err)

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("name", "Zeeshan")
	data.Add("age", "32")
	data.Add("blog", "acloudtechie.com")

	response, err := http.PostForm(myUrl, data)
	CheckNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func CallAzureFunc() {
	const myUrl = "http://localhost:7071/api/getname?name=zeeshan&demo=test"

	response, _ := http.Get(myUrl)

	fmt.Println("getprofile func returning ", response.Status)

	// for j, val := range response.Header {
	// 	fmt.Printf("%v %v\n", j, val)
	// }

	funcContent, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(funcContent))

	parsedUrl, err := url.Parse(myUrl)
	CheckNilErr(err)

	fmt.Println("API path is URL ", parsedUrl.Path)

	qparams := parsedUrl.Query()
	for j, val := range qparams {
		fmt.Printf("Query parameter: %v %v\n", j, val)
	}

	defer response.Body.Close()

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	CheckNilErr(err)

	fmt.Printf("This data is retrieved from file:\n%v", string(data))

}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)

}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)

	return input
}

func stringInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
