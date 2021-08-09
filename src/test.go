package src

import (

	"fmt"
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

type MonthlyRequest struct {
	
	Jan int `json:"jan"`
	Feb int `json:"feb"`
	Mar int `json:"mar"`
	Apr int `json:"apr"`
	May int `json:"may"`
	Jun int `json:"jun"`
	Jul int `json:"jul"`
	Aug int `json:"aug"`
	Sep int `json:"sep"`
	Oct int `json:"oct"`
	Nov int `json:"nov"`
	Dec int `json:"dec"`
	
}

func MonthlyRequestReport(filename string) (string,error) {

	file,err := os.Open(filename)
	
	if err != nil {
	
		panic(err)
		
		return "",err
		
	}
	
	scanner := bufio.NewScanner(file)
	
	var monthRep MonthlyRequest 
	
	for scanner.Scan() {
		
		report := strings.Split(scanner.Text()," ")
		
		date := report[3]
		
		date = date[1:]
		
		dateRay := strings.Split(date,"/")
		
		month := dateRay[1:2]
		
		
		switch month[0] {
			
			case "May" :
				
				monthRep.May = monthRep.May +1
			
			case "Jun":
				
				monthRep.Jun = monthRep.Jun + 1
			
			case "Jan":
				
				monthRep.Jan = monthRep.Jan + 1
				
			case "Feb":
				
				monthRep.Feb = monthRep.Feb + 1
			
			case "Mar":
				
				monthRep.Mar = monthRep.Mar + 1
		
		}

	}
	
	data,err := json.Marshal(monthRep)
	
	if err != nil {
	
		panic(err)
		
		return "",nil
		
	}
	
	result := string(data)
	
	fmt.Println(result)
	
	return result,nil
	
	
}



//MonthlyUserBaseStatusReport

func MonthlyUserBaseStatusReport(filename string) (string,error) {

	file,err := os.Open(filename)
	
	
	if err != nil {
	
		panic(err)
		
		return "",err
	}
	
	defer file.Close()

	currentMonth := "Jan"
	
	iplist := make(map[string]int)
	
	scanner := bufio.NewScanner(file)
	
	var monthRep MonthlyRequest
	
	for scanner.Scan() {
		
		report := strings.Split(scanner.Text()," ")
		
		ip := report[0]
		
		date := report[3]
		
		date = date[1:]
		
		dateRay := strings.Split(date,"/")
		
		month := dateRay[1:2]
		
		
		
		if currentMonth != month[0] {
		
			currentMonth = month[0]
			
			iplist = make(map[string]int)
			
		
		}
		
		switch month[0] {
		
			case "May":
				
				
				if val,ok := iplist[string(ip)];ok {
				
					iplist[ip] = iplist[ip] + 1
					
					if val > 100 {
					
						monthRep.May = monthRep.May + 1 
					}
					
				}else{
				
					iplist[ip] = 0
				}
				
			
			case "Jun":
				
				val,ok := iplist[ip]
				
				if ok {
					
					iplist[ip] = iplist[ip] + 1
					
					if val > 100 {
					
						monthRep.Jun = monthRep.Jun + 1 
					}
				
				}else {
				
					iplist[ip] = 0
				}
		
		
		}
	
	}
	
	fmt.Println(monthRep)
	
	data,err := json.Marshal(monthRep)
	
	if err != nil {
		
		panic(err)
		
		return "",err
	}
	
	return string(data),nil

	

}


//


func ReadFile(filename string) (result string ,err error) {
	
	file,err := os.Open(filename) 
	
	if err != nil {
		
		panic(err)
		
		return "",err
	
	}
	
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	var reportArray []ReportStruct
	
	
	
	for scanner.Scan() {
		
		
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



