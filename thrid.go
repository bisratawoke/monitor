package main

import (

	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"github.com/bisratawoke/third/src"
	
	
)

//authenticate user token

type CustomError struct{}

func (m *CustomError) Error() string {

    return "unauthenticated"
    
}



//authenticate user
func auth(r * http.Request) (err error) {
	
	val,ok := r.Header["Authorization"]
	
	if ok {
		
		token := val[0][len("token "):]
		
		client := &http.Client{}
		
		req,err := http.NewRequest(http.MethodGet,"http://account/api/account/service/check",nil)
		
		req.Header.Add("authorization",fmt.Sprintf("token %s",token))
		
		if err != nil {
			
			panic(err)
			
			return err
			
		}
		
		res,err := client.Do(req)
		
		if err != nil {
		
			panic(err)
			
			return err
		
		}
		
		defer res.Body.Close()
		
		if res.StatusCode == 200 {
		
			return nil
		
		}else {
			
			fmt.Println("something went wrong")
			
			return &CustomError{}
		}
		
	}else {
	
		return &CustomError{}
	
	}

}


//access log handler
func monthlyRequestStatusHandler(w http.ResponseWriter, r * http.Request) {
	
	
		/*err := auth(r)	
	
		if err != nil {
		
			fmt.Fprintf(w,"not authenticated",401)
			
			return
		}
		*/
		
		data,err := src.MonthlyRequestReport("./nginx_logs")
		
		if err != nil {
		
			http.Error(w,err.Error(),500)
			
			return
		}
		
		w.Header().Add("Content-Type","application/json")
		
		w.Header().Add("Access-Control-Allow-Origin","*")
		
		w.Header().Set("Access-Control-Allow-Headers","Authorization")
		
		fmt.Fprintf(w,data)
		
	
	

}


//read files


//error log handler

func errorLogHandler(w http.ResponseWriter, r * http.Request) {
	
	
		
		err := auth(r)
		
		if err != nil {
		
			http.Error(w,err.Error(),400)
			
			return
		}
		
		data,err := ioutil.ReadFile(os.Args[2])
		
		if err != nil {
		
			http.Error(w,err.Error(),500)
			
			return
		}
		
		
		
		
		fmt.Fprintf(w,string(data))
		
	
	

	
	


}

//monthlyUserStatusHandler

func monthlyUserStatusHandler(w http.ResponseWriter, r * http.Request) {

	/*err := auth(r)
		
		if err != nil {
		
			http.Error(w,err.Error(),400)
			
			return
	}
	*/
	
	data,err := src.MonthlyUserBaseStatusReport("./nginx_logs")
	
	if err != nil {
	
		panic(err)
		
		http.Error(w,err.Error(),500)
		
		return 
	
	}
		w.Header().Add("Content-Type","application/json")
		
		w.Header().Add("Access-Control-Allow-Origin","*")
		
		w.Header().Set("Access-Control-Allow-Headers","Authorization")
		
	
	fmt.Fprintf(w,string(data))
		

}


// main function
func main () { 


	http.HandleFunc("/status/monthlyReq",monthlyRequestStatusHandler)
	
	http.HandleFunc("/status/monthlyUser",monthlyUserStatusHandler)
	
	http.HandleFunc("/logs/error",errorLogHandler)
	
	log.Fatal(http.ListenAndServe(":8000",nil))	
	
}











