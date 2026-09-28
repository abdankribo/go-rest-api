package main
import("encoding/json";"log";"net/http";"strconv";"sync")
type Item struct{ID int `json:"id"`;Name string `json:"name"`;Done bool `json:"done"`}
var mu sync.Mutex
var items=[]Item{{1,"Design API",true},{2,"Write tests",false}}
func api(w http.ResponseWriter,r *http.Request){w.Header().Set("Content-Type","application/json");mu.Lock();defer mu.Unlock();if r.Method=="GET"{json.NewEncoder(w).Encode(items);return};if r.Method=="POST"{var x Item;if json.NewDecoder(r.Body).Decode(&x)==nil{x.ID=len(items)+1;items=append(items,x);json.NewEncoder(w).Encode(x);return}};if r.Method=="DELETE"&&len(r.URL.Path)>7{id,_:=strconv.Atoi(r.URL.Path[7:]);for i,x:=range items{if x.ID==id{items=append(items[:i],items[i+1:]...);break}};w.WriteHeader(http.StatusNoContent);return};http.Error(w,"bad request",400)}
func main(){http.HandleFunc("/api/items",api);log.Println("Go API :8080");log.Fatal(http.ListenAndServe(":8080",nil))}
