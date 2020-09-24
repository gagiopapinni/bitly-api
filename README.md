# bitly-api

__demo is available here http://207.154.252.235/__    

__Description__    

There are two endpoints: */create* [POST] and */<key\>* [GET]   

* */create* makes a short link by accepting a json object with fields *url* and *key*    
And returns a json object with either *error* or *key* field   
if *key* is ommited in request, a random one will be generated instead     

* */$key* redirects to the original url or to a *not-found* page, if no such key has been created before   

__Running__    
 
 1. create a *config.json* file in the project folder, with *PORT* (int) and *MONGOHQ_URL* (string) fields   
 2. run $ *go install*   
 3. run $ *go run main.go*     
 the landing page should be available now at http://localhost:<PORT\>


