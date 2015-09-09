package main

import (
    "os"
    "bufio"
    "github.com/seiflotfy/cuckoofilter"
    "log"
)

var pws = []string{
        "HelloWorld",
        "123456",
        "password",
        "Swag",
    }

func main() {

	cf := cuckoofilter.NewCuckooFilter(16000000)

	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        success := cf.Insert([]byte(line))
        if !success{
        	log.Fatal("FAAAAIL!")
            os.Exit(1)
        }
    }
    log.Printf("Added %d entries.",cf.GetCount())


    for _,pw := range pws {
        if cf.Lookup([]byte(pw)) {
            log.Printf("Contains '%s'\n",pw)
        }else{
            log.Printf("Yay, not in list: '%s'\n",pw)
        }
    }

}  