package apitp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	isBs    = 1
	isSen   = 2
	isSense = 3
	isPseq  = 4
)

var propositions []string

func regexReplace(def string) string {
	def = strings.ReplaceAll(def, "[[[text ", "--")
	regex, errRegex := regexp.Compile("{a_link\\|")
	if errRegex != nil {
		log.Fatal("regex compilation error!\n", errRegex)
	}
	def = string(regex.ReplaceAll([]byte(def), []byte{}))
	regex, errRegex = regexp.Compile("{dx_def.+dx_def}")
	if errRegex != nil {
		log.Fatal("regex compilation error!\n", errRegex)
	}
	def = string(regex.ReplaceAll([]byte(def), []byte{}))
	regex, errRegex = regexp.Compile("({d_link\\|)|(\\|[a-zA-Z -]+:?[0-9]?})")
	if errRegex != nil {
		log.Fatal("regex compilation error!\n", errRegex)
	}
	def = string(regex.ReplaceAll([]byte(def), []byte{}))
	regex, errRegex = regexp.Compile("({sx\\|)|(\\|+[0-9a-zA-Z]?.?})")
	if errRegex != nil {
		log.Fatal("regex compilation error!\n", errRegex)
	}
	def = string(regex.ReplaceAll([]byte(def), []byte{}))
	regex, errRegex = regexp.Compile("(] \\[vis.+]+)|(})")
	if errRegex != nil {
		log.Fatal("regex compilation error!\n", errRegex)
	}
	def = string(regex.ReplaceAll([]byte(def), []byte{}))
	return def
}

func apiFetch(word string) []APIdata {
	urlAPI := "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"
	apiParams := "?key=3f40c380-59d2-4171-a2c2-749b02ea6a6c"
	urlReq := urlAPI + word + apiParams

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, urlReq, nil)
	if errReq != nil {
		log.Fatal("Request failed!\n", errReq)
	}
	req.Header.Set("User-Agent", "Thorgan")

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Fatal("Sending request error !\n", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal("Reading body request error !\n", errBody)
	}

	var reqData []APIdata
	errJSON := json.Unmarshal(body, &reqData)
	if errJSON != nil {
		fmt.Println(word, ": not found")
		errJSON = json.Unmarshal(body, &propositions)
		if errJSON != nil {
			log.Fatal("Data format error !\n", errJSON)
		}
		return nil
	}

	return reqData
}

func sseqDisplay(sseq interface{}, dataType int) int {
	switch dataType {
	case isBs:
		byteData, _ := json.Marshal(sseq)
		var bs Bs
		errJSON := json.Unmarshal(byteData, &bs)
		if errJSON != nil {
			log.Fatal("bs unmarshal error!\n", errJSON)
		}
		fmt.Print("\t", bs.Sense.Sn, " ")
		def := strings.ReplaceAll(bs.Sense.Dt[0][1], "{bc}", ": ")
		def = regexReplace(def)
		fmt.Println(def)
		fmt.Println()
		dataType = 0
	case isSen:
		byteData, _ := json.Marshal(sseq)
		var sen Sen
		errJSON := json.Unmarshal(byteData, &sen)
		if errJSON != nil {
			log.Fatal("sen unmarshal error!\n", errJSON)
		}
		fmt.Println("\t", sen.Sn, " ", strings.Join(sen.Sls, " "))
		fmt.Println()
		dataType = 0
	case isSense:
		byteData, _ := json.Marshal(sseq)
		var sense Sense
		errJSON := json.Unmarshal(byteData, &sense)
		if errJSON != nil {
			log.Fatal("sense unmarshal error!\n", errJSON)
		}
		fmt.Print("\t\t", sense.Sn, " ")
		def := strings.ReplaceAll(fmt.Sprint(sense.Dt[0][1]), "{bc}", ": ")
		def = regexReplace(def)
		fmt.Println(def)
		fmt.Println()
		dataType = 0
	}
	return dataType
}

func apiDisplay(data []APIdata) {
	for _, entry := range data {
		fmt.Println(strings.ToUpper(entry.Meta.Id))
		if entry.Meta.Offensive {
			fmt.Println("offensive")
		}
		fmt.Println(entry.Fl)
		fmt.Print(entry.Hwi.Hw)
		if len(entry.Hwi.Prs) > 0 {
			fmt.Print(" | ")
		}
		for i, prs := range entry.Hwi.Prs {
			if prs.Mw != "" {
				fmt.Print("\\", prs.Mw, "\\ ")
				if i != len(entry.Hwi.Prs)-1 {
					fmt.Print("OR ")
				}
			}
		}
		fmt.Println()
		fmt.Println()
		for i, ins := range entry.Ins {
			if ins.If != "" {
				fmt.Print(ins.If, " ")
				if i == len(entry.Ins)-1 {
					fmt.Println()
				}
			}
		}
		fmt.Println()
		for _, def := range entry.Def {
			if def.Vd != "" {
				fmt.Println("\t", def.Vd)
			}
			var dataType int
			for _, sseqRoot := range def.Sseq[0] {
				for _, sseq := range sseqRoot {
					switch sseq {
					case nil, "":
						dataType = 0
						continue
					case "bs":
						dataType = isBs
						continue
					case "sen":
						dataType = isSen
						continue
					case "sense":
						dataType = isSense
						continue
					case "pseq":
						dataType = isPseq
						continue
					}

					if dataType == isPseq {
						byteData, _ := json.Marshal(sseq)
						var pseqArr2d [][]interface{}
						errJSON := json.Unmarshal(byteData, &pseqArr2d)
						if errJSON != nil {
							log.Fatal("pseq unmarshal error!\n", errJSON)
						}
						for _, pseqArr := range pseqArr2d {
							for _, pseq := range pseqArr {
								// fmt.Printf("pseq: %#v\n", pseq)
								switch pseq {
								case nil, "":
									dataType = 0
									continue
								case "bs":
									dataType = isBs
									continue
								case "sen":
									dataType = isSen
									continue
								case "sense":
									dataType = isSense
									continue
								}
								dataType = sseqDisplay(pseq, dataType)
							}
						}
					} else {
						dataType = sseqDisplay(sseq, dataType)
					}
				}
			}
		}
		fmt.Println()
	}
}
