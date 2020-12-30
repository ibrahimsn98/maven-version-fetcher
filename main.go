package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Conf struct {
	Token string `json:"token"`
	Urls []string `json:"urls"`
}

type Repo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageType string `json:"package_type"`
	CreatedAt   string `json:"created_at"`
}

func main() {
	jsonFile, err := os.Open("conf.json")

	if err != nil {
		log.Fatal(err)
	}

	bytes, _ := ioutil.ReadAll(jsonFile)

	var conf Conf

	err = json.Unmarshal(bytes, &conf)

	if err != nil {
		log.Fatal(err)
	}

	urlParams := "?order_by=created_at&sort=desc"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	for i := 0; i < len(conf.Urls); i++ {
		req, err := http.NewRequest(http.MethodGet, conf.Urls[i] + urlParams, nil)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("PRIVATE-TOKEN", conf.Token)

		res, err := spaceClient.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		var repos []Repo
		err = json.Unmarshal(body, &repos)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s %s\n", color.GreenString("->"), repos[0].Name, color.GreenString(repos[0].Version))

		if strings.Contains(repos[0].Version, "SINGLE") {
			for j := 0; j < len(repos); j++ {
				if strings.Contains(repos[j].Version, "MULTI") {
					fmt.Printf("%s %s %s\n", color.GreenString("->"), repos[j].Name, color.GreenString(repos[j].Version))
					break
				}
			}
		} else if strings.Contains(repos[0].Version, "MULTI") {
			for j := 0; j < len(repos); j++ {
				if strings.Contains(repos[j].Version, "SINGLE") {
					fmt.Printf("%s %s %s\n", color.GreenString("->"), repos[j].Name, color.GreenString(repos[j].Version))
					break
				}
			}
		}
	}
}
