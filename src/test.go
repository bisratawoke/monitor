package src

import (

	
	"os"
	"bufio"
	"strings"
	"encoding/json"
	

)

type ReportStruct struct {

	Ip string `json:"ip"`
	
	Method string `json:"method"`
	
	Status string `json:"status"`
	
	Date string `json:"date"`
	
	Month string `json:"month"`
	
	Day string `json:"day"`
	
	Year string `json:"year"`

}

func ReadFile(filename string) (result string ,err error) {
	
	file,err := os.Open("./nginx_logs") 
	
	if err != nil {
		
		panic(err)
		
		return "",err
	
	}
	
	scanner := bufio.NewScanner(file)
	
	var reportArray []ReportStruct
	
	count := 1
	
	for scanner.Scan() {
		
		if count > 10 {
		
			break
			
		}
		
		count = count + 1 
		
		report := strings.Split(scanner.Text()," ")
		
		ip := report[0]
		
		date := report[3]
		
		date = date[1:]
		
		dateRay := strings.Split(date,"/")
		
		day := dateRay[0:1]
		
		dayStr := string(day[0])
		
		month := dateRay[1:2]
		
		monthStr := string(month[0])
		
		year := dateRay[2:3]
		
		year = strings.Split(year[0],":")
		
		yearStr := string(year[0])
		
		method := report[5]
		
		method = method[1:]
		
		status := report[8]
		
		//fmt.Println(report)
		
		rep := ReportStruct{Ip:ip,Date:date,Method:method,Status:status,Day:dayStr,Month:monthStr,Year:yearStr}
		
		reportArray = append(reportArray,rep)
		
	
		//fmt.Println(fmt.Sprintf(" ip address : %s \t date : %s \t  method : %s \t status : %s",ip,date,method,status))
		
		
	
	}
	
	data,err := json.MarshalIndent(reportArray," "," ")
		
	if err != nil {
			
		panic(err)
		
		return "",err
			
	}
		
	result = string(data)
		
		
	return result,nil


}



/*
func main() {

	file,err := os.Open("../nginx_logs")
	
	if err != nil {
	
	
		panic(err)
		
	}
	
	scanner := bufio.NewScanner(file)
	
	count := 1
	
	for scanner.Scan() {
		
		if count > 10 {
		
			break
			
		}
		
		count = count + 1
		
		report := strings.Split(scanner.Text()," ")
		
		/*ip := report[0]
		
		date := report[3]
		
		date = date[1:]
		
		method := report[5]
		
		method = method[1:]
		
		status := report[8]
	
		
		date := report[3]
		
		date = date[1:]
			
		dateRay := strings.Split(date,"/")
		
		day := dateRay[0:1]
		
		dayStr := string(day[0])
		
		month := dateRay[1:2]
		
		monthStr := string(month[0])
		
		year := dateRay[2:3]
		
		year = strings.Split(year[0],":")
		
		yearStr := string(year[0])
		
		fmt.Println(fmt.Sprintf("date : %s \t day : %s \t , month : %s \t , year : %s",date,dayStr,monthStr,yearStr))
	}


}

*/
