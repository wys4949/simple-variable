# simple-variable
This is a library for parameter conversion

	//example::
	//params1 :=777     //int
	//params1 :=777.999   //float64
	//params1 :=777.9992222222222   //float64
	//params1 :=323.22  //float64
	//params1 :=time.Now().Unix()  //time is int64
	params1 :=3.988812
	T := new(simple_variable.ToSring)  
	//type1 :=T.TypeOf(params1) //get type
	//values := T.Str(params1)  //to string
	//values := T.Int(params1) //to int
	//values := T.Int64(params1) //to int64
	//values := T.Float64(params1)
	values := T.Float64(params1) //to float64
	values = T.Round(values,2,false) //Rounding or false ,Keep two decimal places                                                       
	//values := T.StrAdd("33",1,"1.02")  //string link
	//values := T.StrSubtract("1a2bsdf","f",2) // string trim
	//values := T.MathAdd("33",1,"1.02")  //num add
	//values := T.MathSubtract("33",1,"1.02") //num subtract
	//values := T.MathMultiply("33",2,"1.02") //num multiply
	//values := T.MathDivide("33",2,"1.02")//num divide
	//values := T.MathPow("2",2,"3.0")  //num pow
	//values := T.MathRemainder("6",4)  //num remainder
	//value := T.Byte("wdwddw")  // to byte
    //values := T.Str(value)    // byte to str
    values := T.IsSame("6",6)   // this is Weak type comparison
    
	fmt.Println(values)
	//fmt.Printf("%t id %v",values)