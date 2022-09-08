package main

func main() {
	post := newPost("Hello, world!")
	post.publish()
	post.publish()
	post.publish()
}
