package main

var config = Config{}

func init() {
	config.Read()
}

func main() {
	subject := "Get latest Tech News directly to your inbox"
	receiver := "example@gmail.com"
	r := NewRequest([]string{receiver}, subject)
	r.Send("templates/template.html", map[string]string{"username": "9bany"})
}
