package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var URL string = "https://jsonplaceholder.typicode.com/posts/"

type Post struct {
	id     int
	userId int
	title  string
	body   string
}

func getPosts() interface{} {
	// get data API from url
	response, err := http.Get(URL)

	// check the error of getting API in above
	if err != nil {
		errMsg := fmt.Sprintf("Gagal mengambil data dari API: %s", err)
		return errMsg
	}

	// close the response body if the process is finished
	defer response.Body.Close()

	// check the status code from data API
	if response.StatusCode == 200 {
		// success and read the API from response body
		data, _ := ioutil.ReadAll(response.Body)

		// print the data
		return string(data)
	} else {
		// error API response
		errResponse := fmt.Sprintf("Gagal mengambil data dari API. Kode status: %d", response.StatusCode)
		return errResponse
	}
}

func getPost(idPost int) interface{} {
	// get data API from url based on id post
	response, err := http.Get(URL + strconv.Itoa(idPost))

	// check the error of getting API in above
	if err != nil {
		errMsg := fmt.Sprintf("Gagal mengambil data dari API: %s", err)
		return errMsg
	}

	// close the response body if the process is finished
	defer response.Body.Close()

	// check the status code from data API
	if response.StatusCode == 200 {
		// success and read the API from response body
		data, _ := ioutil.ReadAll(response.Body)

		// print the data
		return string(data)
	} else {
		// error API response
		errResponse := fmt.Sprintf("Gagal mengambil data dari API. Kode status: %d", response.StatusCode)
		return errResponse
	}
}

func insertNewPost(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)

	if err != nil {
		errConvertToJSON := fmt.Sprintf("Gagal mengonversi data ke JSON: %s", err)

		return errConvertToJSON
	}

	response, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		errSendToServer := fmt.Sprintf("Gagal mengirim data ke server: %s", err)

		return errSendToServer
	}

	defer response.Body.Close()

	if response.StatusCode == 201 {
		successMsg := fmt.Sprintf("Data berhasil di simpan di server")

		return successMsg
	} else {
		errResponse := fmt.Sprintf("Gagal menyimpan data ke server. Kode status: %s", err)

		return errResponse
	}
}

func deletePost(idPost int) string {
	// combine the url with id post
	newUrl := fmt.Sprintf("%s/%d", URL, idPost)

	client := &http.Client{}
	request, err := http.NewRequest("DELETE", newUrl, nil)

	if err != nil {
		failRequestDelete := fmt.Sprintf("Gagal membuat permintaan DELETE: %s", err)

		return failRequestDelete
	}

	response, err := client.Do(request)

	if err != nil {
		failDeleteData := fmt.Sprintf("Gagal menghapus data: %s", err)

		return failDeleteData
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		successDelete := fmt.Sprintf("Data dengan ID %d berhasil dihapus. \n", idPost)

		return successDelete
	} else {
		errResponse := fmt.Sprintf("Gagal menyimpan data ke server. Kode status: %s", err)

		return errResponse
	}
}

func main() {
	// 1.
	fmt.Println(getPosts())

	// 2.
	fmt.Println(getPost(3))

	// 3.
	dataToPost := map[string]interface{}{
		"userId": 11,
		"title":  "Example Posting from user id 11",
		"body":   "Hi this is user 11, thank you for making this public API",
	}

	fmt.Println(insertNewPost(dataToPost))

	// 4.
	fmt.Println(deletePost(3))

	// 5.
	// Link Jawaban: https://github.com/alif-syhm01/task-intro-echo-alif.git
}
