package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "strings"
)

func main() {

	bRunSetup := false
	sOS := runtime.GOOS
	
	
	//Main Script
	
	if bRunSetup == true {
		//First Time Setup------------------------------------------------------

	
	} else {
		//Launcher--------------------------------------------------------------
		if sOS == "windows" {
			
			reader := bufio.NewReader(os.Stdin)
			
			c := exec.Command("cmd", "/C", "retroarch --help",)
			stdout, err := c.Output()
			
			if err != nil {
				fmt.Println("Error")
			}

			fmt.Print(string(stdout))
			
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("hi", text) == 0 {
			  fmt.Println("hello, Yourself")
			}
	

		}else{
			fmt.Println("ERROR: Unsupported Operating System. Only Windows is Supported.")
			os.Exit(1)
		}
	}
	
		
}