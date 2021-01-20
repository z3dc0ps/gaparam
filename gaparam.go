package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "crypto/tls"
)

func main() {


    const DebugColor   = "\033[0;36m%s\033[0m"

    if len(os.Args) == 1 {
        fmt.Printf(DebugColor,"Usage : gaparam URL")
        return
    }else{

        url := os.Args[1]
        http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //For cancelling certificate checking
        response, err := http.Get("http://web.archive.org/cdx/search/cdx?url=*."+url+"/*&output=txt&fl=original&collapse=urlkey&page=/")
        if err != nil {
            log.Fatal(err)
        }
        defer response.Body.Close()

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }

        responseString := string(responseData)
        fmt.Println(responseString)
    }
    
}