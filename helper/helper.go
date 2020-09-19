package helper 


import (
       "fmt"

       "net/url"
       "context"
       "log"
       "time" 
       "math/rand"
 
       "go.mongodb.org/mongo-driver/mongo"
       "go.mongodb.org/mongo-driver/mongo/options"
 )


func ConnectDb(MONGOHQ_URL string) *mongo.Collection {
   // MONGOHQ_URL := "mongodb+srv://santa:EgIbOadoED9FVuOu@cluster0-zpliy.mongodb.net/bitly?retryWrites=true&w=majority"
   
    client, err := mongo.NewClient(options.Client().ApplyURI(MONGOHQ_URL))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
            log.Fatal(err)
    }

    fmt.Println("Connected to db")

    collection := client.Database("bitly").Collection("urls")

    return collection
}


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
func RandSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

