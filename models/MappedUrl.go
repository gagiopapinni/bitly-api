package models

import(
      // "fmt"
       "github.com/gagiopapinni/bitly-api/helper"

       "context"
       "log"
       "errors"
       "go.mongodb.org/mongo-driver/bson"
       "go.mongodb.org/mongo-driver/mongo"
 )


type MappedUrl struct {
	Url string `json:"url,omitempty"`	
	Key string `json:"key,omitempty"`
}



func (u MappedUrl) doesKeyExist(collection *mongo.Collection) bool {
            
            var result bson.M
            err := collection.FindOne(context.TODO(), bson.D{{"key", u.Key}}).Decode(&result)
            if err != nil {
                if err == mongo.ErrNoDocuments {
                     return false
                } else { 
                     log.Fatal(err) 
                }
            }

            return true          
    
}


func (u MappedUrl) validate(collection *mongo.Collection) error {
        if len(u.Key)==0 || len(u.Key)>60  {
           return errors.New("bad key length")
        } else if u.doesKeyExist(collection) { 
           return errors.New("specified key exists") 
        }

        if !helper.IsValidUrl(u.Url) {
           return errors.New("invalid url")
        }

        return nil
}

func (u *MappedUrl) SetUniqueKey(collection *mongo.Collection) {
           var m_url MappedUrl
           found := false
           for !found {
               m_url.Key = helper.RandSeq(8)
               if !m_url.doesKeyExist(collection) {
                  found = true
               } 
           }
           u.Key = m_url.Key           
}



func (u MappedUrl) Save(collection *mongo.Collection) error {
	err := u.validate(collection)
        if err != nil {
           return err
        }
   
        collection.InsertOne(context.TODO(),u)
       
        return nil
}


func (u *MappedUrl) FetchUrl(collection *mongo.Collection) error {
   
        var result bson.M
        err := collection.FindOne(context.TODO(), bson.D{{"key", u.Key}}).Decode(&result)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                return errors.New("no such key")
            } else { 
                log.Fatal(err) 
            }
        }

        u.Url = result["url"].(string)

        return nil
}
