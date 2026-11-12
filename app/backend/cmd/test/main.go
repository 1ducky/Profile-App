package main

import (
	"log"
	"profile-template/internal/article"
)

func main() {
	content := "asdasdwadwa[test](assets:uploads/7fa90e1e-7fbf-491c-a38f-6dcc9f27774f_b1f39aab-495a-4ca2-8007-61b7198a4379.png)" + "asdasdwadwa[test)(assets:uploads/7fa90e1e-7fbf-491c-a38f-6dcc9f27774f_b1f39aab-495a-4ca2-8007-61b7198a4379.png)" + "![sadsa](assets:uploads/7fa90e1e-7fbf-491c-a38f-6dcc9f27774f_14262ec9-9808-410a-ad6c-7abe38050732.png )"
	filekeys := article.ExtractFileKeys(content)
	for _, key := range filekeys {
		log.Println(key)
	}
	log.Print(filekeys)
}
