package main

import ("fmt"
		"strings"
		"bytes"
		"bufio"
		 "os"
		"os/exec"
		)
		
var endString string		
var buffer bytes.Buffer	
var instDisplay string = "<inst>"
var listArray[20] string
var tempCount = 1

//Consist of Array with the shapes, the position which it is located, and the shape to be processed
// function returns a int value after it moves through the array it sends back its last location	
func treeSQR( sqrArray[] string, arrayPos int) int{

//Function check to see if the there is any shapes in the in the listArray
//if not it changes the structures as it moves to the end
if listArray[0] != "SQR" { 
			listArray[1]= "<inst>"
			listArray[0] = "SQR"
		arrayPos++
		
		// Called the function so it can be processed with the valid format 
		treeSQR(sqrArray[0:],arrayPos)	
	}else{ if listArray[1] == "" || listArray[1] == "<inst>"{ // after transforming it is place in a format that can be parsed 
	
			if sqrArray[arrayPos] == "SQR"{ // Ensure we are not Validating a Shape
				arrayPos++
			}
			
			// Retrieve the Coordinated from the array
			// Proceeding to the next value
			var curCoord string=sqrArray[arrayPos]
			arrayPos++
			var secCoord string=sqrArray[arrayPos]
			//arrayPos++		
			
			// Using Slices we get each Values 
			var test string 
			x:=curCoord[0:1]
			y:=curCoord[1:2]
			yy:=secCoord[0:1]
			xx:=secCoord[1:2]
			arrayPos=arrayPos-2
			
			//The Printing format for the lower part of the tree
			fmt.Printf("\n |\n")
			fmt.Printf("SQR\n/ \\\n")
			fmt.Printf("<coord>,<coord>\n")
			fmt.Printf(" /  \\   /  \\\n")
			fmt.Printf("<x><y>,<x><y>\n"+x+"   "+y+"  "+yy+"  "+xx)
		
			listArray[0] = "<inst>"
			listArray[1] = "<inst_list>"
						
			tempCount=tempCount-1
				if(tempCount >= 0){				
					listArray[tempCount]=""
					
				}
			return arrayPos	
		}
		return arrayPos
		 }
	
	return arrayPos
}

//Consist of Array with the shapes, the position which it is located, and the shape to be processed
// function returns a int value after it moves through the array it sends back its last location
func treeTRI( triArray[] string, arrayPos int) int{

//Function check to see if the there is any shapes in the in the listArray
//if not it changes the structures as it moves to the end
	if listArray[0] != "TRI" { 
			listArray[1]= "<inst>"
			listArray[0] = "TRI"
		arrayPos++
		
// Called the function so it can be processed with the valid format 	
		treeTRI(triArray[0:],arrayPos)	
	}else{ if listArray[1] == "" || listArray[1] == "<inst>"{ // after transforming it is place in a format that can be parsed   
			if triArray[arrayPos] == "TRI"{ // Ensure we are not Validating a Shape
				arrayPos++
			}
			
			// Retrieve the Coordinated from the array
			// Proceeding to the next value
			var curCoord string=triArray[arrayPos]
			arrayPos++
			var secCoord string=triArray[arrayPos]
			arrayPos++
			var triCoord string=triArray[arrayPos]
			
			// Using Slices we get each Values 
			x:=curCoord[0:1]
			y:=curCoord[1:2]			
			yy:=secCoord[1:2]
			xx:=secCoord[0:1]			
			xxx:=triCoord[0:1]
			yyy:=triCoord[1:2]
			
			//The Printing format for the lower part of the tree
			fmt.Printf("\n |\n")
			fmt.Printf("TRI\n/ \\\n")
			fmt.Printf("<coord>,<coord>,<coord>\n")
			fmt.Printf("<x><y>,<x><y>,<x><y>\n"+x+"   "+y+"   "+xx+"  "+yy+"   "+xxx+"  "+yyy)
			listArray[0] = "<inst>"
			listArray[1] = "<inst_list>"
			
			tempCount=tempCount-1
				if(tempCount >= 0){				
					listArray[tempCount]=""					
				}
			
		}
		 }
	
	return arrayPos
}

//Creating a Generic tree Function allows to input the shape for shapes with similar structures 
//Consist of Array with the shapes, the position which it is located, and the shape to be processed
// function returns a int value after it moves through the array it sends back its last location
func treeGene( geneArray[] string, arrayPos int, shapeType string) int{

//Function check to see if the there is any shapes in the in the listArray
//if not it changes the structures as it moves to the end
	if listArray[0] != shapeType {  
			listArray[1]= "<inst>"
			listArray[0] = shapeType
		arrayPos++
		treeGene(geneArray[0:],arrayPos,shapeType)	// Called the function so it can be processed with the valid format 	
	}else{ if listArray[1] == "" || listArray[1] == "<inst>"{ // after transforming it is place in a format that can be parsed   
			if geneArray[arrayPos] == shapeType{ 	// Ensure we are not Validating a Shape
				arrayPos++
			}
			// Retrieve the Coordinated from the array
			var curCoord string=geneArray[arrayPos] 
			arrayPos++
			
			// Using Slices we get each Values 
			x:=curCoord[0:1] 
			y:=curCoord[1:2]
			fmt.Printf("\n |\n")
			
			// The Only Difference with Circle is that it Has a Coordinated
			fmt.Printf(shapeType+"\n / \\\n")
			if shapeType == "CIR"{
				fmt.Printf("<coord>\n")
			}			
			fmt.Printf("<x><y>\n"+x+"   "+y)
			
		}
		 }

	return arrayPos
}

func parseString(vartobeParse string){
	
	//Assigning the upper parse to the array statically 
	//This is done by assigning value to the string with Tabs and new Line
	var topParse string = "\n\t<program> \n /          |     \t\\ "
	var secondParse string = "\n start \t <inst_list> \t finish \n\t    |"
	
	var buffer bytes.Buffer
	
	var displayTree[5] string 	//Variable for displaying the Tree Array 
	displayTree[0] = topParse
	displayTree[1] = secondParse

	var amtShape int=0

	var x int
	//Place the sentence into a string array remove all the space, commas and semi colon
	// want to achieve acquiring pure values  
	stringarray := strings.FieldsFunc(vartobeParse, func(r rune) bool {
	if r==' '|| r==',' || r == ';'{
			return true
		}
		return false
	})

	// Count the total amount of Semi Colons in the sentence
	//with that we can produce how many inst_list 
	for _, r := range vartobeParse {
        if string(r) == ";"{
			amtShape++
    }
	}
	
	//After acquiring the total amount of shape, For Loops the decides how many addition inst_list
	// It is placed in the listArray with is basically hold the string for inst_list
	fmt.Print("%v",amtShape)
	if amtShape > 0{
		amtShape--
			for ; amtShape >= 0; amtShape--{
					buffer.WriteString("<inst_list>")
					newOutput:=buffer.String()					
					listArray[1]=(newOutput)				
				}		
		}
	var count int=0

	
	for _, r := range listArray[1] {
						if string(r) == "_"{
							count++	}
					}
					
				listArray[0]= "<inst>"
				for ; count != 0; count--{	
					listArray[tempCount]= "<inst_list>"
				tempCount++
				}
	// Using the array with pure values we can isolate which values is the Given Shape
	// Running through the array we place a switch to decide which function is called 
	fmt.Println(strings.Trim(fmt.Sprint(displayTree), "[] "))	//Prints the upper tree
	for x=0; stringarray[x] != "finish"; { 
	
	//Running a Loop in the formatted array to Check which values has the make up of the shape
	//We Use a switch to call each function 
	//fmt.Println(stringarray[x])
	///fmt.Println("%v",x)
		switch stringarray[x] {
			case "SQR":
						//Each case has a Print function for the upper part of the tree
						//below each function is call this is done for all case
						fmt.Println(strings.Trim(fmt.Sprint(listArray), "[] "))
						x= treeSQR(stringarray[0:], x )
						x++
						//fmt.Println("%v",x)
			case "TRI":
						fmt.Println(strings.Trim(fmt.Sprint(listArray), "[] "))
						x= treeTRI(stringarray[0:], x )
			case "GRID":
						fmt.Println(strings.Trim(fmt.Sprint(listArray), "[] "))
						var tempString string = "GRID"
						x= treeGene(stringarray[0:],x , tempString)
			case "FILL":
						fmt.Println(strings.Trim(fmt.Sprint(listArray), "[] "))
						var tempString string = "FILL"
						x= treeGene(stringarray[0:],x , tempString)
			case "CIR":
						fmt.Println(strings.Trim(fmt.Sprint(listArray), "[] "))
						var tempString string = "CIR"
						x= treeGene(stringarray[0:],x , tempString)		
			default:
					x++
				
		}
	}
/*----------------------------------------------------------------------------------------------------------*/	
}

//Check the coord of the Shape Passed 
//Function parameters are the shapes the stringarray and the outputArray
//We are Using end string to append to the end of every derivation
func checkCoord(outputArray []string, stringarray []string, shape string, crtPosition int)string{
	var i int = 0
	
	//Validates if the Shape is a square
	//if it is we should have two values so increment by two to proceed with the Right most derivation
	if shape == "SQR"{
	crtPosition=crtPosition+2

	for ; i != 3; i++{
		var frtCoord string=stringarray[crtPosition]
		x:=frtCoord[0:1]
		y:=frtCoord[1:2]
	// Validates the Y to the give range
	if y=="0" ||y=="1" || y=="2" || y=="3" || y=="4" || y=="5" || y=="6"|| y=="7"|| y=="8"|| y=="9"{
			if i == 1 {	//If the loop was incremented if so proceed with this new instructions					
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x>"+y+endString)
					endString=(""+y+endString)			
					}else{ if endString != "" { //If the endString is not empty if so proceed with this new instructions
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<coord>"+endString)
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<x><y>"+ endString)		
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<x>"+y+endString)						
							endString=(""+y+endString)
														
					}else{ //The First time that we are entering the loop so access the default format
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<x><y>"+ "finish")		
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<x>"+y+" finish")}
					}

					if x=="A" || x=="B" || x=="C" || x=="D" || x=="E" || x=="F" || x=="G"|| x=="H"|| x=="I"|| x=="J"{
						if i == 1 {	//If the loop was incremented if so proceed with this new instructions	
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" "+x+endString)
							endString=(" SQR "+x+endString)
									
						}else{	if endString !=""{//If the endString is empty if so proceed with this new instructions
								fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,"+x+endString)
								fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x><y>, "+x+endString)
								endString=(","+x+endString)
															
						}else{		//The First time that we are entering the loop for the x so access the default format
									fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,"+x+y+" finish")
									fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x><y>,"+x+y+" finish")
									endString=(","+x+""+y+" finish")
									}

						}							
					}else{
						fmt.Println("Error: "+x+ " is invalid")
					}
	crtPosition--
	if stringarray[crtPosition] == shape{
	break
	}
	
	
	}else{
		fmt.Println("Error: "+y+ " is invalid")
		break
	}
}

}

/*************************************************************************************/
	//Validates if the Shape is a Triangle
	//if it is we should have three values so increment by two to proceed with the Right most derivation
	
	if shape == "TRI"{
	crtPosition=crtPosition+3

	for ; i != 3; i++{
		var frtCoord string=stringarray[crtPosition]
		x:=frtCoord[0:1]
		y:=frtCoord[1:2]
			// Validates the Y to the give range
		if y=="0" ||y=="1" || y=="2" || y=="3" || y=="4" || y=="5" || y=="6"{
			if i == 1 {	//If the loop was incremented if so proceed with this new instructions		
				fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<x>"+y+endString)
				endString=(""+y+endString)			
					}else {
						if i == 2 {//If the loop was incremented the third time if so proceed with this new instructions
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<x><y>,"+endString)
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<x>"+y+","+endString)	
							endString = (y+","+endString)		
							}else{ 
								if endString != "" {//If the endString is not empty if so proceed with this new instructions 
										fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<coord>,<coord>"+endString)
										fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<coord>,<x><y>"+ endString)		
										fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<coord>,<x>"+y+endString)						
										endString=(""+y+endString)	
								}else{ //The First time that we are entering the loop so access the default format
										fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<coord>,<x><y>"+ "finish1")		
										fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<coord>,<x>"+y+" finish1")
									}																	
								}					
					
					}

		if x=="A" || x=="B" || x=="C" || x=="D" || x=="E" || x=="G" || x=="F"{
				if i == 1 { //If the loop was incremented if so proceed with this new instructions	
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,"+x+endString)
							endString=(x+endString)									
						}else{ 						
							if endString !=""{							
								if i == 2{//If the loop was incremented the third time if so proceed with this new instructions
									fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" "+x+endString)
									endString = (" TRI "+x+endString)									
									}else{//If the endString is not empty if so proceed with this new instructions 
								fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<coord>,"+x+endString)
								fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<x><y>, "+x+endString)
								endString=(","+x+endString)
								}															
							}else{		//The First time that we are entering the loop so access the default format
									fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<coord>,"+x+y+" finish2")
									fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+"<coord>,<x><y>,"+x+y+" finish2")
									endString=(","+x+""+y+" finish")

									}

							}							
		}
	crtPosition--
	
	}

}
}

/**************************************************************************************************/
//Validates if the Shape is a Triangle
//if it is we should have one value so increment by one to proceed with the Right most derivation
//Either of these values can we FILL GRID or CIR

if shape == "FILL" || shape == "GRID" || shape == "CIR"{
	crtPosition=crtPosition+1
	for ; i != 1; i++{
		var frtCoord string=stringarray[crtPosition]
		x:=frtCoord[0:1]
		y:=frtCoord[1:2]
		// Validates the Y to the give range
		if y=="0" ||y=="1" || y=="2" || y=="3" || y=="4" || y=="5" || y=="6"{
				if shape == "CIR"{ // Circle Carry a different format so the printing is as follows
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>"+ "finish1")
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x><y>,<x>"+ "finish1")
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x>"+y+" finish1")
					}else{ // Its just the regular format
					
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x><y>"+ "finish1")		
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <x>"+y+" finish1")
					}
			if x=="A" || x=="B" || x=="C" || x=="D" || x=="E" || x=="G" || x=="F"{
					fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" "+x+y+" finish2")
					endString=(","+x+""+y+" finish")									
					}												
		crtPosition--	
	}
}
}

return endString
}
// Check instruction checks the Shape add see if its has the Proper Shape
//Accepts the string array along with the outputArray and sentence passed		
func checkInstructions(outputArray []string, stringarray []string, curPosition int,sentence string){	
		
		var x int=0
		var tempString string 
		
		//outputArray has the amount of shape to be printed 
		//Controls how many time to look for a give shape
		
		for ; outputArray[x] != "";x++{		
		}				
		for ; curPosition !=0; curPosition--{
		//Validates the Give shapes to proceed to the next step
		if stringarray[curPosition] == "CIR" || stringarray[curPosition] == "GRID" || stringarray[curPosition] == "FILL" || stringarray[curPosition] == "TRI" ||stringarray[curPosition] == "SQR"{
			if outputArray[x] == ""{	
				x--}
			if endString != ""{ //Check to see if we have appended before  if we had we proceed with rearranging the inst_list
				if outputArray[x] == "<inst>"{ //Displays at key point to show the derivation
						fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+endString)
						outputArray[x]=stringarray[curPosition]
						fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+endString)
						tempString = outputArray[x] 
						endString=checkCoord(outputArray[0:],stringarray[0:],tempString,curPosition)
						outputArray[x]=""	
						x--
						if x != 0 {
						outputArray[x] = "<inst>"
						}					
				}
			}else{ // outputArray is arrange properly we can perform the derivation if not we fix it				
			if outputArray[x] == "<inst>"{
			outputArray[x]=stringarray[curPosition]
			// Just to is play the proper formatting a the derivation point of coords
			switch stringarray[curPosition] {
				case "SQR":
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<coord> finish")
				case "TRI":
							fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" <coord>,<coord>,<coord> finish")
					
				case "GRID":
							break
				case "FILL":
							break
				case "CIR":
							break	
				default:
						break
			}			
				tempString = outputArray[x] 
				endString=checkCoord(outputArray[0:],stringarray[0:],tempString,curPosition)
				outputArray[x]=""
				x--
				if x != 0 {
				outputArray[x] = "<inst>"
				}
			}
		
			}
			
			}
			
			
		}
}	

//Check inst_list main function is to validate how many shapes is being passed
//by that we can how evaluate how many times we should look for a given shape
//function parameters is the sentence array and outputArray the lenght of the array and the curPosition that we are in the array		
func checkInstList(stringarray []string, outputArray []string, curPosition int, length int,sentence string){
	//declaration of variables for the counter control that follows 
		var buffer bytes.Buffer
		var count int=0
		var amtShape int=0
		var tempCount int = 2
		//Check to see the amount of semi colon in the sentence and by that we can evaluate how many shapes
		for _, r := range sentence {
        if string(r) == ";"{
			amtShape++}		
		}
		//static assignment for the 1st part of the inst
		outputArray[1]="<inst>"
	
		//if there is no one inst then we can processed in making a format for one inst
		if amtShape == 0 {
			outputArray[1]="<inst>"
			fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" finish")
			checkInstructions(outputArray[0:], stringarray[0:],curPosition,sentence)
		//If there is more that two can we can add <inst_list>
		// buffer doest that but as a single string.
		}else{ //
		for ; amtShape >= 0; amtShape--{
					buffer.WriteString("<inst_list>")
					newOutput:=buffer.String()					
					outputArray[tempCount]=(newOutput)
				}		
		//Get the value of how man y inst_list was place by using a token of "_"
		for _, r := range outputArray[2] {
						if string(r) == "_"{
							count++	}
					}
		count--	
		//Proceed with placing it in its respective Array location 		
		for ; count >= 1; {	
			if outputArray[tempCount] == "<inst>"{
				count--
					}else{outputArray[tempCount]= "<inst_list>"
					count--
				tempCount++
				}
				
		}
		//Print the formatted derivation
		fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" finish")
		tempCount--
		//Change the format of the derivation
		if outputArray[tempCount] == "<inst_list>"{
			outputArray[tempCount] = "<inst>"
		}
		//Print the formatted derivation and Proceed with calling the function to check each inst_list
		fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" finish")	
		checkInstructions(outputArray[0:], stringarray[0:],curPosition,sentence)
		}
		
		}
			
	
	


//derivation Function Accepts a string and the string array that was created

func derivation(stringarray []string, sentence string){
	//Variable declaration 
	//getting the total size of the array
	//along with the creating a new array outputArray for each inst
	var buffer bytes.Buffer
	var length int=(len(stringarray)-1)
	var curPosition int=length
	var outputArray [20]string
	stringarray[length]= strings.Trim(stringarray[length], "\r\n ")
	fmt.Printf("\n")

	//The First validation to check whether the sentence beginnings with Start or Finish
	//Displays Output at the end
	if stringarray[0]=="start"{		
		if stringarray[length]=="finish"{
				// Write the 1st message to the array
				//if true then processsed in formatting the outputArray with the given inst_list
				//this is done by calling the function checkInstList
				buffer.WriteString("program->start") 
				outputArray[0]=buffer.String()
				outputArray[1]="<inst_list>"			
				fmt.Println(strings.Trim(fmt.Sprint(outputArray), "[] ")+" finish")	
				outputArray[1]="<inst>"
				checkInstList(stringarray[0:],outputArray[0:],curPosition,length,sentence )
			}else{
				fmt.Println("Error: Invalid ending..."+stringarray[length])
				main()
			}
		
	}else{
		fmt.Println("Error: Invalid beginning..."+stringarray[0])
		main()
	}
	
	
}
		
		
		
func main(){

	//List the Language for the Program
	fmt.Printf("\n ")
	fmt.Printf("<program>   -> start <inst_list> finish\n")
	fmt.Printf(" <inst_list> -> <inst>\n")
	fmt.Printf("             |<inst>;<inst_list>\n")
	fmt.Printf("      <inst> -> SQR <coord>,<coord>\n")
	fmt.Printf("             -> TRI <coord>,<coord>,<coord>\n")
	fmt.Printf("             -> CIR <coord>\n")
	fmt.Printf("             -> GRID <x><y>\n")
	fmt.Printf("             -> FILL <x><y>\n")
    fmt.Printf("     <coord> -> <x> <y>\n")
	fmt.Printf ("         <x> -> A|B|C|D|E|F|G|H|I|J\n")
	fmt.Printf ("         <y> -> 1|2|3|4|5|6|7|8|9 \n")
	fmt.Printf(" \n")
	fmt.Printf(" \n")
	
	// The Get Statement for Capturing the Sentence that is inputted 
	//Validates to see if any string is empty or not
	r := bufio.NewReader(os.Stdin)
    fmt.Printf("Please enter a string base on the grammar above: \n")
    sentence, err := r.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }	
	if sentence=="quit\r\n" || sentence=="QUIT\r\n"{ // Await response for Quit
		return 
	}else{

	// Check to see if the Sentence input has spaces commas and semi Colons and removes
	//It also place it in an array
	stringarray := strings.FieldsFunc(sentence, func(r rune) bool {
	if r==' '|| r==',' || r == ';'{
			return true
		}
		return false
	})
		// Calling the function derivation to beginning derivation
		derivation(stringarray[0:], sentence)
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

		var b []byte = make([]byte, 1)

		fmt.Printf("any key to continue...")
		os.Stdin.Read(b)
		parseString(sentence)
		main()
		
		
	}

	
			
}