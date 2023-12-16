// You can edit this code!
// Click here and start typing.
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const = `[UV2]
ORGANIZATION="{{.Organization}}"
NAME="{{.LastName}}", "{{.FirstName}}"
EMAIL="{{.Email}}"
[C51]
PATH="{{.ExecutableDir}}\C51"
VERSION=V9.61
BOOK0=HLP\Release_Notes.htm("Release Notes")
BOOK1=HLP\C51TOOLS.chm("Complete User's Guide Selection", C)
TDRV0=BIN\MON51.DLL ("Keil Monitor-51 Driver")
TDRV1=BIN\ISD51.DLL ("Keil ISD51 In-System Debugger")
TDRV2=BIN\MON390.DLL ("MON390: Dallas Contiguous Mode")
TDRV3=BIN\LPC2EMP.DLL ("LPC900 EPM Emulator/Programmer")
TDRV4=BIN\UL2UPSD.DLL ("ST-uPSD ULINK Driver")
TDRV5=BIN\UL2XC800.DLL ("Infineon XC800 ULINK Driver")
TDRV6=BIN\MONADI.DLL ("ADI Monitor Driver")
TDRV7=BIN\DAS2XC800.DLL ("Infineon DAS Client for XC800")
TDRV8=BIN\UL2LPC9.DLL ("NXP LPC95x ULINK Driver")
TDRV9=BIN\JLinkEFM8.dll ("J-Link / J-Trace EFM8 Driver")
TDRV10=BIN\JLinkIS2083.dll ("J-Link / J-Trace IS2083 Driver")
TDRV11=BIN\Nuvoton_8051_Keil_uVision_Driver.dll ("Nuvoton 8051 Keil Driver")
RTOS0=Dummy.DLL("Dummy")
RTOS1=RTXTINY.DLL ("RTX-51 Tiny")
RTOS2=RTX51.DLL ("RTX-51 Full")
`

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Data struct {
	ExecutableDir string //%USERPROFILE%\Downloads\Keil_v5\
	Organization  string
	FirstName     string
	LastName      string
	Email         string
}

var (
	project string
)

func Write_TOOLS_INI(data *Data) error {
	file, err := os.Create(filepath.Join(data.ExecutableDir, "TOOLS.INI"))
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.New("tools_ini").Parse(TOOLS_INI)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, *data)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.StringVar(&project, "project", "", "set project (eg. xxx.uvprojx) for CICD in working directory")
	flag.Parse()

	ex, err := os.Executable()
	checkErr(err)

	org := os.Getenv("ORG")
	firstname := os.Getenv("FIRSTNAME")
	lastname := os.Getenv("LASTNAME")
	email := os.Getenv("EMAIL")

	exePath := filepath.Dir(ex)

	err = Write_TOOLS_INI(&Data{Organization: org, FirstName: firstname, LastName: lastname, Email: email, ExecutableDir: exePath})
	checkErr(err)

	os.Setenv("PATH", "%systemroot%\\System32")

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], ":", pair[1])
	}

	datapath := filepath.Join(exePath, "UV4", "UV4.exe")

	if len(project) > 0 {
		cmd := exec.Command(datapath, "-j0", "-b", project)
		fmt.Println("[CICD] project:", project)
		cmd.Run()
	} else {
		fmt.Println("[Exe]")
		cmd := exec.Command(datapath)
		cmd.Start()
	}
}
