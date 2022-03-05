package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const ServerURL = "http://api.socialify.cf/api/v0.1"
const ContentType = "application/json"
const UserAgent = "Socialify-iOS"
const OS = "iOS_14.6"
const AppVersion = "0.1.0"

func getHeaders(timestamp int64, authToken []byte) map[string]string {
    return map[string]string {
        "Content-Type": ContentType,
        "User-Agent":  UserAgent,
        "OS": OS,
        "Timestamp": fmt.Sprint(timestamp),
        "AppVersion": AppVersion,
        "Accept": ContentType,
        "AuthToken": string(authToken),
        
    }
}

func generateAuthToken(endpoint string, timestamp int64) []byte {
    password := []byte(fmt.Sprintf("$begin-%s$.%s+%s+%s#%d#.$end-%s$", endpoint, AppVersion, OS, UserAgent, timestamp, endpoint))
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    return hash
}

type GetKeyResponse struct {
    Data struct {
        Key string `json:"pubKey"`
    } `json:"data"`
}

func getKey(timestamp int64, authToken []byte) (string, error) {
    headers := getHeaders(timestamp, authToken)
    url := fmt.Sprintf("%s/getKey", ServerURL)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        panic(err)
    }
    for key, value := range headers {
        req.Header.Add(key, value)
    }
    client := http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    // fmt.Println(string(body))
    var response GetKeyResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        return "", err
    }

    return response.Data.Key, nil
}

func main() {
    timestamp := time.Now().Unix()
    authToken := generateAuthToken("getKey", timestamp)

    var wg sync.WaitGroup
    for i := 1; i < 1000; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            fmt.Printf("[%d] sending request\n", i)
            key, err := getKey(timestamp, authToken)
            if err != nil {
                fmt.Printf("[%d] failed: %s\n", i, err)
            } else {
                fmt.Printf("[%d] received key: ...%s... \n", i, key[len(key)-10:])
            }
        }(i)
    }
    wg.Wait()
}

