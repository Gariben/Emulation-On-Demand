package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "runtime"
  //"strings"
)

var bDebug bool = true

func retroarch_LaunchGame(sPath string, sFlags string, bPrintOutput bool){

	if runtime.GOOS == "windows" {
		if bDebug==true {
			fmt.Println("DEBUG: Running Retroarch command:")
			fmt.Println("\t> "+sPath+" "+sFlags)
		}
		c := exec.Command(sPath+" "+sFlags,)
		stdout, err := c.Output()
	

	
		if err != nil {
			fmt.Println("ERROR: Failed to run RetroArch command.")
			if string(stdout) != "" {
			
				fmt.Println("***Begin ERROR Log***")	
				fmt.Print(string(stdout))
				fmt.Println("***End ERROR Log***")
			}else {
				fmt.Println("ERROR: No Error was returned...")
			}
		}
		
		if bPrintOutput == true {
			fmt.Print(string(stdout))
		}	
	}
}

func cmd_WaitForNewline(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press ENTER to Continue.")
	reader.ReadString('\n')
}



func main() {

	bRunSetup := false
	sOS := runtime.GOOS
	
	//Initialize Dir variables
	sDir_Default_Install := ""
	
	if sOS == "windows" {
		sDir_Default_Install = "\"C\\:Program Files\\RetroArch\\"
	
	}else{
		fmt.Println("ERROR: Unsupported Operating System. Only Windows is Supported.")
		os.Exit(1)
	}	
	
	
	

	
	
	//Main Script
	
	if bRunSetup == true {
		//First Time Setup------------------------------------------------------

	
	} else {
		//Launcher--------------------------------------------------------------
		if sOS == "windows" {
			sPath_RetroArch := sDir_Default_Install + "retroarch.exe" + "\""
			sArgs_RetroArch := "--help"
			

			retroarch_LaunchGame(sPath_RetroArch, sArgs_RetroArch, true)

			cmd_WaitForNewline()

		}else{
			fmt.Println("ERROR: Unsupported Operating System. Only Windows is Supported.")
			os.Exit(1)
		}
	}
	
		
}