# isorender
isomorphic rendering api

proof of concept of small api containing a unique endpoint `/render`  

### Start node api
```
node index.js
```
api will listen on port 3000  
### start go app
```
go run main.go
```
go app listens on port 5000  

### Todo:  
* Define the json parameters for /render  
```json
{
"file": "/index.jsx",
"vars" : { "name": "Mike" }
}
```
* Pass vars to react vars when rendering
* Make template file dynamic
* document / generates clients with swagger 
