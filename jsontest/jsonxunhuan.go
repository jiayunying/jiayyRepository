package main 
import ( 
	"encoding/json" 
"fmt" 
) 
type Transport struct 
{ 
	Time string 
	MAC string 
	Id string 
	Rssid string 
	} 
	func main() { 
		var st []Transport 
		t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"} 
		st = append(st, t1) 
		t2 := Transport{Time: "66", MAC: "77", Id: "88", Rssid: "99"} 
		st = append(st, t2) 
		buf, _ := json.Marshal(st) 
		fmt.Println(string(buf)) 
		var str = string(buf) 
		var st1 []Transport 
		err := json.Unmarshal([]byte(str), &st1) 
		if err != nil { 
			fmt.Println("some error") 
			} 
		fmt.Println(st1) 
		fmt.Println(st1[0].Time) 
		fmt.Println(len(st1))
		var Msg []map[string]string 
		json.Unmarshal([]byte(str), &Msg) 
		fmt.Println(Msg) 
	}
