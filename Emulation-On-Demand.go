package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "strings"
)

var bDebug bool = true
var iVerbosity = 3

//Verbosity 0 is End User
//Verbosity 1 is Troubleshooting User
//Verbosity 2 is Troubleshooting Dev
//Verbosity 3 is "I need to know everything you're doing"




type structFilePath interface {
//TODO: Make Cross OS

	sGetFullDir() (sFullDir string){
		sSliceBuild := ""
		for _, dir := range sDir {
			if runtime.GOOS == "windows"{
				sSliceBuild += dir + "\\"	
			}
		}
	}




	msInstallation string
	msDir string[]	//{"Program Files", "RetroArch"}
	msExecutable string //retroarch.exe
	msFullPath string //

}


type structLocalizedString struct{
	//Localized for regions
	nameUSA string	//Super Nintendo
	nameEUR string	//Super Famicom
	nameJPN string	//Super Famicom
}

type structConsole struct{
	sAbbreviation string
	sConsoleName LocalizedString
	sCoreName string
	sExtension string
	sPath_Games string
}

var iMessageCounter = 1
func verbosePrint(iTextLVL int, sText string) {
	if iTextLVL >= iVerbosity {
		fmt.Println("* " + strconv.Itoa(iMessageCounter) + ")" + sText)
	}

}

//func EOD_FirstRun()
//func EOD_Launcher()



func retroArch_TestInstall(sPath string, bPrintOutput bool){

	if runtime.GOOS == "windows" {
		if bDebug==true {
			fmt.Println("DEBUG: Running Retroarch command:")
			fmt.Println("\t> "+sPath+" "+sFlags)
		}
		c := exec.Command(sPath,sFlags)
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


func retroarch_LaunchGame(sPath string, sFlags string, bPrintOutput bool){
	fmt.Print("LaunchGame Function")
}

func windows_WaitForNewline(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press ENTER to Continue.")
	reader.ReadString('\n')
}

func windows_ReplaceSpaces(sPath string)(sOut string){
	strings.Replace(sPath, sOut, "ky", 2)
	return 
}


func main() {

	bRunSetup := false
	sOS := runtime.GOOS
	

	verbosePrint(3, "Initializing Directory and Executable variables.")
	//Anything blank will be determined by OS.
	sDir_EOD_Install := ""
	sDir_RetroArch = ""
	sExec_RetroArch = ""
	
	verbosePrint(3, "Initializing Console Structs.")
	//Initialize Structs for Consoles
	gba = new(Console)
	gba.sConsoleUSAName_Full = "Gameboy Advance"
	gba.sConsoleUSAName_Short = "GBA"
	gba.sCoreName = "EmuDonk"
	gba.sRetroArchFlags = ""
	gba.sExtension = ".gba"
	gba.sPath = ""
	gba.sPathGames = ""
	gba.sPathSaves = ""
	
	
	//snes = new(Console)
	//snes.sConsoleNameUSA_Full = "Super Nintendo"
	
	if sOS == "windows" {

		//Jacobu Laptop (for testing)

		
		//Jacobu Laptop
		//sDir_Default_GBA_Games = "C:\\Users\\jacob\\Desktop\\Emulation\\GBA\\Games"
		//sDir_Default_GBA_Games = "C:\\Emulation\\GBA\\Games"
		
		//sDir_EOD_Install = "\"C\\:Program Files\\RetroArch\\"
		verbosePrint(3, "Configuring WINDOWS system directories ")
		sDir_EOD_Install := "C:\\Users\\jacob\\Desktop\\Emulation\\"
		sDir_RetroArch = "Retroarch\\"
		sExec_RetroArch = "retroarch.exe"
		
		verbosePrint(3, "Configuring WINDOWS Emulator Paths.")
		gba.sPath_Games = "\\GBA\\Games"
		gba.sPath_Games = "\\GBA\\Games"
		
		var lConsoleList = []Console { gba }
		
		
		
		//RetroArch
		retroArch_TestInstall()
		
		

		

		
		
		
		
		
		
		
		//Validate retroarch installation exists
		
		
		//If command fails, search system for Retroarch
		
		
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
			sPath_RetroArch := sDir_Default_Install + "retroarch.exe"
			sArgs_RetroArch := "--help"
			
			
			//Search for Game
			//If Search returns nil, search for supported emulator extensions.
			
			
			

			//sGame := "Wario Land 4 (USA, Europe)"
			
			//retroarch_LaunchGame(sPath_RetroArch, sArgs_RetroArch, true)

			windows_WaitForNewline()

		}else{
			fmt.Println("ERROR: Unsupported Operating System. Only Windows is Supported.")
			os.Exit(1)
		}
	}
	
		
}