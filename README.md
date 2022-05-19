## CC26 Polyglottal Project - Typeahead Suggestion  
  
**Typeahead Suggestion**  
An auto suggestion feature built in Go. When typing a word in the input form, it auto-suggests all the English words that share the same prefix as you type.  
Implemented using a [trie data structure](https://en.wikipedia.org/wiki/Trie).  
This repository only contains code for backend.  
  
  
## Frontend and Backend  
This project's frontend and backend are built separately.  
Source files for frontend are at: https://github.com/mikako-shirai/Typeahead-Suggestion-frontend  
  
  
## Deployment  
Frontend: deployed at https://typeahead-suggestion.an.r.appspot.com/  
Backend: deployed at https://typeahead-suggestion-be.an.r.appspot.com/  
(Since this is a server side, there will be nothing on the page)  
  
  
## Built with  
**Frontend**  
- [React](https://reactjs.org/)  
- [GCP App Engine](https://cloud.google.com/appengine/)  
  
**Backend**  
- [GCP App Engine](https://cloud.google.com/appengine/)  
  
  
## Getting Started    
1. clone this repository  
```
$ git clone https://github.com/mikako-shirai/Typeahead-Suggestion.git
```  
2. move to the root `Typeahead-Suggestion` directory  
```
$ cd Typeahead-Suggestion
```  
3. run the server locally  
```
$ go run main.go
```  
  