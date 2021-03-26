package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "strings"
  "strconv"
)

var bDebug bool = true
var iIndention = 0
var iVerbosity = 3

//Verbosity 0 is End User
//Verbosity 1 is Troubleshooting User
//Verbosity 2 is Troubleshooting Dev
//Verbosity 3 is "I need to know everything you're doing"




type interfaceStructFuncs interface {
//TODO: Make Cross OS
//Struct function
	//sPrint() string
	getFullDir() string
	getFullPath() string
	oneName(sName string)
	printAll()
	printConsole()
}

type structLocalizedString struct{
	//Localized for three release regions
	sUSA string	//Super Nintendo
	sEUR string	//Super Famicom
	sJPN string	//Super Famicom
}

type structConsole struct{
	sAbbreviation string
	sConsoleName structLocalizedString
	sCoreName string
	sExtension string
	sPath_Games string
}


type structFilePath struct{
	sDir []string	//{"Program Files", "RetroArch"}
	sExecutable string //retroarch.exe
	sInstallation string //
}

//structLocalizedString funcs
func (sLS *structLocalizedString)oneName(sName string) {
	sLS.sUSA = sName
	sLS.sEUR = sName
	sLS.sJPN = sName
}

func (sLS structLocalizedString)printAll() {
	fmt.Println("(USA): " + sLS.sUSA)
	fmt.Println("(EUR): " + sLS.sEUR)
	fmt.Println("(JPN): " + sLS.sJPN)
}
	


//structConsoleFunctions
func (sC structConsole)printConsole() {
	fmt.Println( "***About this Console***")
	fmt.Println("-Abbreviation: " + sC.sAbbreviation )
	fmt.Println("-ConsoleName:")
	sC.sConsoleName.printAll()
	fmt.Println("-CoreName: " + sC.sCoreName )
	fmt.Println("-Extension: " + sC.sExtension )
}


//structFilePath funcs
//Merge these...
func (sFP structFilePath)getFullDir() string {
		sDir := sFP.sDir
		sSliceFlat := ""
		for _, dir := range sDir {
			if runtime.GOOS == "windows"{
				sSliceFlat += dir + "\\"	
			}
		}	
		return sSliceFlat
}

func (sFP structFilePath)getFullPath() string {
		sDir := sFP.sDir
		sSliceFlat := ""
		for _, dir := range sDir {
			if runtime.GOOS == "windows"{
				sSliceFlat += dir + "\\"	
			}
		}	
		return sSliceFlat+sFP.sExecutable
}


//structLocalizedString funcs

var iMessageCounter = 1
//func indent() {
//	strings.Repeat("\t", iIndention)
//}

func verbosePrint(iTextLVL int, sText string) {
	if iTextLVL >= iVerbosity {
		fmt.Println("* " + strconv.Itoa(iMessageCounter) + ")" + sText)
	}

}

//func EOD_FirstRun()
//func EOD_Launcher()



func retroArch_TestInstall(sPath string, bPrintOutput bool){
//run "retroarch --help" from installation to test installation, paths.
	sFlags := "--help"

	if runtime.GOOS == "windows" {
		if bDebug==true {
			fmt.Println("DEBUG: Running Retroarch command:")
			fmt.Println("\t> "+sPath + " " + sFlags)
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
	
	//Eventual Parent directory.
	sPath_EOD_Install := ""
	
	
	verbosePrint(3, "Initializing Console Structs.")
	//Initialize Structs for Consoles
	var gba *structConsole
	gba = new(structConsole)
	gba.sAbbreviation = "GBA"
	gba.sConsoleName.oneName("Gameboy Advance")
	gba.sCoreName = "EmuDonk"
	gba.sExtension = ".gba"
	
	//snes = new(Console)
	//snes.sConsoleNameUSA_Full = "Super Nintendo"
	if iVerbosity >=3 {
		iIndention++
		gba.printConsole()
		iIndention--
	}
	
	var retroarch *structFilePath
	retroarch = new(structFilePath)
	
	
	if sOS == "windows" {
	
		verbosePrint(3, "Configuring WINDOWS system directories ")
		sPath_EOD_Install =""
		
		//Manual, for now...
		retroarch.sDir = make([]string,3)
		retroarch.sDir[0] = "\\"
		retroarch.sDir[1] = "Program Files\\"
		retroarch.sDir[2] = "RetroArch\\"
		retroarch.sExecutable = "retroarch.exe"
		
		verbosePrint(3, "Configuring WINDOWS Emulator Paths.")
		gba.sPath_Games = "\\GBA\\Games"
		
		
		//Roll up installed consoles into list
		//var lConsoleList = []*structConsole { gba }
		
		
		
		//Test RetroArch install
		
		
		//"C:\Program Files\RetroArch\retroarch.exe" --help
		retroArch_TestInstall("\"C:" + sPath_EOD_Install + retroarch.getFullPath() + "\"", iVerbosity>=3)
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