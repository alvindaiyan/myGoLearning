package main

import ("fmt"
		"math"
		"net/http"
		"strconv"
		"log"
		)

// todo: need error checking
func Future_Value_Formula(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("GET Param: ", r.URL.Query())
	if pv, err0 := strconv.ParseFloat(r.URL.Query()["present_value"][0], 64); err0 == nil {		
		if i, err1 := strconv.ParseFloat(r.URL.Query()["rate"][0], 64); err1 == nil {
			if n, err2 := strconv.ParseInt(r.URL.Query()["time"][0], 10, 64); err2 == nil {				
				fmt.Fprintf(w, "future value = %f", future_value_formula(pv, i, n) )
			} else {
				log.Fatal("ParseFloat n: ", err2)
			}
		} else {
			log.Fatal("ParseFloat i: ", err1)
		}
	} else {
		log.Fatal("ParseFloat pv: ", err0)
	}
}


// todo: need error checking
func Present_Value_Formula(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("GET Param: ", r.URL.Query())
	if fv, err0 := strconv.ParseFloat(r.URL.Query()["future_value"][0], 64); err0 == nil {		
		if i, err1 := strconv.ParseFloat(r.URL.Query()["rate"][0], 64); err1 == nil {
			if n, err2 := strconv.ParseInt(r.URL.Query()["time"][0], 10, 64); err2 == nil {				
				fmt.Fprintf(w, "present value = %f", present_value_formula(fv, i, n) )
			} else {
				log.Fatal("ParseFloat n: ", err2)
			}
		} else {
			log.Fatal("ParseFloat i: ", err1)
		}
	} else {
		log.Fatal("ParseFloat pv: ", err0)
	}
}
	


// get the fv (future value), pv = present value, i = rate, n = time
func future_value_formula(pv float64, i float64, n int64) float64 {
	if n < 0 {
		// should throw an error
	}

	var fv float64 = math.Pow( float64(1) + i,  float64(n)) * pv
	return fv
}


func present_value_formula(fv float64, i float64, n int64) float64 {
	if n < 0 {
		// should throw an error
	}

	x := float64(1)
	if 1 + i == 0 {
		return 0
	}

	return fv / math.Pow(x + i, float64(n))

}


func NPV(pv float64, inv float64) float64 {
	return pv - inv
}


// todo implement pv = e(t = 1, n) [1/ (1 + i) ^ t * Ct]
// ct is 每一期现金值的总和


// todo fv = e(t = 1, n) [ 1/ (1 + i) ^ (t - 1)* Cn-t+1]


// todo perpetuitie 
// v0 = C / t




func main() {
	// setup the url path
	http.HandleFunc("/future_value_formula", Future_Value_Formula)
	http.HandleFunc("/present_value_formula", Present_Value_Formula)



	//set up the listsen port
    err := http.ListenAndServe(":9090", nil) 

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}