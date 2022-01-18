package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var directory = ""
var logName = "RCS_SENDER"

func getCurrentDate() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
	return fecha

}

func getCurrentYear() string {
	t := time.Now()
	year := fmt.Sprintf("%d",
		t.Year())
	return year
}

func getCurrentMonth() string {
	t := time.Now()
	month := fmt.Sprintf("%02d",
		t.Month())
	return month
}

func checkDirectory() {
	year := getCurrentYear()
	month := getCurrentMonth()
	// Asigna directorio segun fecha actual, si no existe lo crea.
	directory = year + "/" + month
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			panic(err)
		}

	}
}

func getCurrentLog() string {
	currentLog := ""
	fecha := getCurrentDate()
	// Revisa si log diario existe, si existe toma el nombre y lo devuelve, sino crea uno nuevo para la fecha actual
	files, err := ioutil.ReadDir("./" + directory + "/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), fecha) {
			currentLog = file.Name()

		}
	}

	if currentLog == "" {
		currentLog = logName + "." + getCurrentDate() + ".log"
	}

	return currentLog
}

func WriteLog(content string) {
	currentLog := getCurrentLog()
	// abre archivo de log actual, si no existe lo crea segun lo obtenido en la funcion getCurrentLog()
	f, err := os.OpenFile("./"+directory+"/"+currentLog,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(content)

}

type Cursos struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
	Desc   string `json:"desc"`
}

func main() {
	checkDirectory()
	hilo := "012343434"
	curso1 := Cursos{
		Id:     "1",
		Nombre: "sebas",
		Desc:   "algo",
	}

	println(curso1.Nombre)

	curso1.Nombre = "caols"
	mensaje, _ := json.Marshal(curso1)
	salida := string(mensaje)
	WriteLog(fmt.Sprint("Thread:", hilo, " Meesage start validations: ", 12343435, salida))

	num1 := 45
	num2 := 10

	if num1 > num2 {
		WriteLog(fmt.Sprint("El numero ", num1, " es mayor que ", num2))
	}

}
