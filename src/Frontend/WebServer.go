package Frontend

import (
	"MyProjectForWB/src/JsonStruct"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type InfoModels struct {
	Model  string
	Length string
}

func ModelConvert(bam []JsonStruct.JsonStruct, id int) string {
	buf := bytes.Buffer{}
	convert, err := json.Marshal(bam[id])
	if err != nil {
		return "Marshal error"
	}
	_ = json.Indent(&buf, convert, "", "\t")
	return buf.String()
}

func WebServ(bam *[]JsonStruct.JsonStruct) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseFiles("/home/one/GolandProjects/MyProjectForWB/src/Frontend/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		value := r.FormValue("input_id")
		id, er := strconv.Atoi(value)
		info := InfoModels{"", fmt.Sprint(len(*bam))}
		if value == " " {
			info.Model = ""
		} else if er != nil {
			info.Model = "Invalid value"
		} else if id > len(*bam) || id <= 0 {
			info.Model = "Value out of range"
		} else {
			info.Model = ModelConvert(*bam, id-1)
		}
		_ = tmp.Execute(w, info)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
