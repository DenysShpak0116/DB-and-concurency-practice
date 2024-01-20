package main

import (
	"db_working/client"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	client.InitConnection()

	posts := GetAllUsersPosts(7)

	var wg sync.WaitGroup

	for _, post := range posts {

		wg.Add(1)
		go func(post client.Post) {
			defer wg.Done()
			WritePostToDB(post)

			comments := GetAllPostsComments(post.Id)

			for _, comment := range comments {
				go WriteCommentToDB(comment)
			}
		}(post)
	}

	wg.Wait()
	log.Println("Done.")
}

func GetAllUsersPosts(userId int) []client.Post {
	url := "https://jsonplaceholder.typicode.com/posts?userId=" +
		strconv.Itoa(userId)
	rows, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Body.Close()

	var posts []client.Post

	err = json.NewDecoder(rows.Body).Decode(&posts)
	if err != nil {
		panic(err.Error())
	}

	return posts
}

func WritePostToDB(post client.Post) {
	_, err := client.DBClient.Exec(
		"INSERT into posts (userId, id, title, body) VALUES (?,?,?,?)",
		post.UserId, post.Id,
		post.Title, post.Body,
	)
	if err != nil {
		panic(err.Error())
	}
}

func GetAllPostsComments(postId int) []client.Comment {
	url := "https://jsonplaceholder.typicode.com/comments?postId=" +
		strconv.Itoa(postId)
	rows, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Body.Close()

	var comments []client.Comment

	err = json.NewDecoder(rows.Body).Decode(&comments)
	if err != nil {
		panic(err.Error())
	}

	return comments
}

func WriteCommentToDB(comment client.Comment) {
	_, err := client.DBClient.Exec(
		"INSERT into comments (postId, id, name, email, body) VALUES (?,?,?,?, ?)",
		comment.PostId, comment.Id,
		comment.Name, comment.Email,
		comment.Body,
	)
	if err != nil {
		panic(err.Error())
	}
}
