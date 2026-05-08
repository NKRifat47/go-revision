package handlers

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ami handler middle a print hbo.")
}
