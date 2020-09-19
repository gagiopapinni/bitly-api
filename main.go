package main

import(
   "fmt"
   "github.com/gagiopapinni/bitly-api/models"
   "github.com/gagiopapinni/bitly-api/helper"
   "github.com/gin-gonic/gin"
   "go.mongodb.org/mongo-driver/mongo"
   "net/http"
   "os"
   "encoding/json"
 )

var collection *mongo.Collection
type Config struct {
    PORT int
    MONGOHQ_URL string
}

func main() { 
	var config Config
	conf, _ := os.Open("config.json")
	defer conf.Close()
	err := json.NewDecoder(conf).Decode(&config)
	if err != nil {
		fmt.Println("config reading error")
	}

	collection = helper.ConnectDb(config.MONGOHQ_URL)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.POST("/create", func(c *gin.Context) {

		var obj models.MappedUrl

		err_bind := c.ShouldBindJSON(&obj)
		if err_bind != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err_bind.Error()})
		}

		if obj.Key == "" {
			obj.SetUniqueKey(collection)
		}

		err := obj.Save(collection)
       
		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"key": obj.Key})
		}
              

	})

	r.GET("/:key", func(c *gin.Context) {
		key := c.Param("key")
		obj := models.MappedUrl{Key: key}
		err := obj.FetchUrl(collection)
		if err != nil {     
			c.HTML(404,"not-found.html", gin.H{})
		} else {
			c.Redirect(303, obj.Url)
		}

	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"landing.html", gin.H{ "action":"/create" })
	})

	r.Run(fmt.Sprintf(":%d",config.PORT)) 

}




