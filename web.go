package main
import (
    "io"
    "log"
    "net/http"
    "os"
    "strings"
)
const (
    HostVar = "VCAP_APP_HOST"
    PortVar = "VCAP_APP_PORT"
)
func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/data/", dataHandler)
    http.HandleFunc("/graph/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "static/index.html")
    })
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "index.html")
    })
    
    port := os.Getenv(PortVar)
    if port == "" {
        port = "3000"
    }
    http.ListenAndServe(":" + port, nil)
}

func dataHandler(res http.ResponseWriter, req *http.Request) {
    globalKey := strings.TrimSpace(req.URL.Path[len("/data/"):])
    globalKey = strings.TrimRight(globalKey, "/")

    res.Header().Set("Content-Type", "application/json;charset=UTF-8")
    res.Header().Set("Access-Control-Allow-Origin", "*")
    res.Header().Set("Access-Control-Allow-Methods", "GET")
    res.Header().Set("Access-Control-Max-Age", "1728000")

    err := fetchActivenessData(globalKey, res)
    if err != nil {
        log.Println(err)
    }
}

func internalError(res http.ResponseWriter, req *http.Request) {
    http.Error(res, "500 internal error", http.StatusInternalServerError)
}

func fetchActivenessData(globalKey string, writer http.ResponseWriter) error {
    res, err := http.Get("https://coding.net/api/user/activeness/data/" + globalKey)
    if err != nil {
        return err
    }
    defer res.Body.Close()
    _, err = io.Copy(writer, res.Body)
    return err
}