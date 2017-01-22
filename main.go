package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Conf struct {
	URL         string `json:"url"`
	AccessToken string `json:"access_token"`
}

var (
	configFile = flag.String("config", "./manage-members-csv.json", "full path to the JSON config file")
	quiet      = flag.Bool("q", false, "if supplied, will silence prompts")
)

func action(message string) {
	color.Set(color.FgRed)
	fmt.Println(message)
	color.Unset()
}

func success(message string) {
	color.Set(color.FgGreen)
	fmt.Println(message)
	color.Unset()
}

func usage(exitCode int) {
	if !(*quiet) {
		flag.Usage()
	}
	exit(exitCode)
}

func exit(exitCode int) {
	if !(*quiet) {
		success("Press 'Enter' to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	os.Exit(exitCode)
}

func main() {
	flag.Parse()

	// Load config file
	if configFile == nil || strings.TrimSpace(*configFile) == "" {
		action("Expected path to config file in -config")
		usage(1)
	}

	fmt.Printf("Reading config file at %s\n", *configFile)
	f, err := os.Open(*configFile)
	if err != nil {
		action(err.Error())
		usage(1)
	}
	defer f.Close()

	fmt.Println("Decoding config file")
	d := json.NewDecoder(f)
	var conf Conf
	err = d.Decode(&conf)
	if err != nil {
		action(fmt.Sprintf("Config parsing error: %v\n", err))
		usage(1)
	}

	fmt.Println("Validating config file")
	if conf.URL == "" {
		action("Please set the 'url' in the config file")
		usage(1)
	}
	if !strings.HasSuffix(conf.URL, `/`) {
		conf.URL = conf.URL + `/`
	}
	url := fmt.Sprintf("%sapi/v1/users/batch", conf.URL)

	if conf.AccessToken == "" {
		action("Please set the 'access_token' in the config file")
		usage(1)
	}

	csvFile := flag.Arg(0)
	if csvFile == "" {
		action("The CSV file path must be the last argument")
		usage(1)
	}

	fmt.Printf("Reading CSV file at %s\n", csvFile)
	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
		action(err.Error())
		usage(1)
	}

	b, err := ioutil.ReadFile(csvFile)
	if err != nil {
		action(err.Error())
		exit(1)
	}
	br := bytes.NewReader(b)

	fmt.Printf("Sending CSV to %s\n", url)
	req, err := http.NewRequest("POST", url, br)
	if err != nil {
		action(err.Error())
		exit(1)
	}
	req.Header.Set("Content-Type", "text/csv")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.AccessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		action(err.Error())
		exit(1)
	}
	defer resp.Body.Close()
	nb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		action(err.Error())
		exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		action(string(nb))
		exit(1)
	}

	success("Finished processing all users in CSV without error")
	exit(0)
}
