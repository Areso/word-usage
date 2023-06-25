package main

import (
    "os" //used in load()
    "path/filepath"
    "strings" //used in splitMultilineStringToMap
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors" // used for setup CORS
    "fmt"
)

func loadFile() string {
    executablePath, err := os.Executable()
    if err != nil {
        fmt.Println("Failed to get executable path:", err)
        return "error"
    }

    // Get the directory containing the executable
    executableDir := filepath.Dir(executablePath)
    b, err := os.ReadFile(executableDir+"/wordlist10k_en.txt")
    if err != nil {
        panic(err)
    }
    return string(b)
}

func splitMultilineStringToMap(multilineString string) map[int]string {
    lines := strings.Split(multilineString, "\n")
    result := make(map[int]string)
    thecounter := 1
    for _, line := range lines {
        if line == "" {
            continue // Skip empty lines
        }
        key := thecounter
        result[key] = line
        thecounter = thecounter+1
    }
    return result
}

func loadGin(c *gin.Context){
    var multilineFileContent string
    multilineFileContent = loadFile()
    var themap=splitMultilineStringToMap(multilineFileContent)
    c.Set("thedictionary", themap)
    c.Next()
}

func getKeyByValue(m map[int]string, value string) int {
    for key, val := range m {
        if val == value {
            return key
        }
    }
    return -1
}

func getFreq(c *gin.Context) {
    v := c.MustGet("thedictionary") //v has anytype type
    theword  := c.Query("word")
    freq := -1
    if mapa, ok := v.(map[int]string); ok {
        number:= getKeyByValue(mapa, theword)
        fmt.Println(number)
        freq = number
    } else {
        // Type assertion failed
        fmt.Println("Type assertion failed!")
        // Handle the failure or error case
        // ...
        freq = -1
    }
    c.IndentedJSON(http.StatusOK, freq)
}

func healthcheck(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "OK")
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Use(cors.Default())
    router.Use(loadGin)
    router.GET("/get_freq", getFreq)
    router.GET("/healthcheck", healthcheck)
    router.Run("localhost:8090")
}
