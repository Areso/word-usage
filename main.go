package main

import (
    "os"
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

func load() string {
   b, err := os.ReadFile("file.txt")
   if err != nil {
      panic(err)
   }
   return string(b)
}

func loadGin(c *gin.Context){
    c.Set("big", load())
    c.Next()
}

func main() {
    b:=load()
    fmt.Println(b)
    router := gin.Default()
    router.Use(loadGin)
    router.GET("/get_snils", genSnils)
    router.Run("localhost:8080")
}

func genSnils(c *gin.Context) {
    v := c.MustGet("big")
    var i interface{}
    i = v
    fmt.Println(v)
    snils := "its a "+i.(string)
    c.IndentedJSON(http.StatusOK, snils)
}
