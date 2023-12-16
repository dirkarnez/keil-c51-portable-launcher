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

const TOOLS_INI = `[UV2]
ORGANIZATION="{{.Organization}}"
NAME="{{.LastName}}", "{{.FirstName}}"
EMAIL="{{.Email}}"
ARMSEL=1
USERTE=1
TOOL_VARIANT=mdk_lite
RTEPATH="{{.ExecutableDir}}\ARM\PACK"
[ARM]
PATH="{{.ExecutableDir}}\ARM"
VERSION=5.38
PATH1="C:\Program Files (x86)\Arm GNU Toolchain arm-none-eabi\11.2 2022.02\"
TOOLPREFIX=arm-none-eabi-
CPUDLL0=SARM.DLL(TDRV16,TDRV17,TDRV18)                                                                       # Drivers for ARM7/9 devices
CPUDLL1=SARMCM3.DLL(TDRV0,TDRV1,TDRV2,TDRV3,TDRV4,TDRV5,TDRV6,TDRV7,TDRV8,TDRV9,TDRV10,TDRV11)               # Drivers for Cortex-M devices
CPUDLL2=SARMCR4.DLL(TDRV4)                                                                                   # Drivers for Cortex-R4 devices
CPUDLL3=SARMV8M.DLL(TDRV2,TDRV4,TDRV6,TDRV7,TDRV8,TDRV12,TDRV13,TDRV14,TDRV15)                               # Drivers for ARMv8-M devices
BOOK0=HLP\RELEASE_NOTES.HTM("Release Notes for MDK 5.38a")
BOOK1=HLP\ARMTOOLS.chm("Complete User's Guide Selection", C)
TDRV0=BIN\UL2CM3.DLL("ULINK2/ME Cortex Debugger")
TDRV1=BIN\ULP2CM3.DLL("ULINK Pro Cortex Debugger")
TDRV2=BIN\ULPL2CM3.dll("ULINKplus Debugger")
TDRV3=BIN\CMSIS_AGDI.dll("CMSIS-DAP Debugger")
TDRV4=Segger\JL2CM3.dll("J-LINK / J-TRACE Cortex")
TDRV5=BIN\DbgFM.DLL("Models Cortex-M Debugger")
TDRV6=STLink\ST-LINKIII-KEIL_SWO.dll ("ST-Link Debugger")
TDRV7=NULink\Nu_Link.dll("NULink Debugger")
TDRV8=PEMicro\Pemicro_ArmCortexInterface.dll("Pemicro Debugger")
TDRV9=SiLabs\SLAB_CM_Keil.dll("SiLabs UDA Debugger")
TDRV10=BIN\ABLSTCM.dll("Altera Blaster Cortex Debugger")
TDRV11=TI_XDS\XDS2CM3.dll("TI XDS Debugger")
TDRV12=BIN\ULP2V8M.DLL("ULINK Pro ARMv8-M Debugger")
TDRV13=BIN\UL2V8M.DLL("ULINK2/ME ARMv8-M Debugger")
TDRV14=BIN\CMSIS_AGDI_V8M.DLL("CMSIS-DAP ARMv8-M Debugger")
TDRV15=BIN\DbgFMv8M.DLL("Models ARMv8-M Debugger")
TDRV16=BIN\UL2ARM.DLL("ULINK2/ME ARM Debugger")
TDRV17=BIN\ULP2ARM.DLL("ULINK Pro ARM Debugger")
TDRV18=Segger\JLTAgdi.dll("J-LINK / J-TRACE ARM")
DELDRVPKG0=ULINK\UninstallULINK.exe("ULINK Pro Driver V1.0")
[ARMADS]
PATH="{{.ExecutableDir}}\ARM"
PATH1=".\ARMCLANG\bin\"
CPUDLL0=SARM.DLL(TDRV16,TDRV17,TDRV18)                                                                # Drivers for ARM7/9 devices
CPUDLL1=SARMCM3.DLL(TDRV0,TDRV1,TDRV2,TDRV3,TDRV4,TDRV5,TDRV6,TDRV7,TDRV8,TDRV9,TDRV10,TDRV11)        # Drivers for Cortex-M devices
CPUDLL2=SARMCR4.DLL(TDRV4)                                                                            # Drivers for Cortex-R4 devices
CPUDLL3=SARMV8M.DLL(TDRV2,TDRV4,TDRV6,TDRV7,TDRV8,TDRV12,TDRV13,TDRV14,TDRV15)                        # Drivers for ARMv8-M devices
BOOK0=HLP\mdk5-getting-started.pdf ("MDK-ARM Getting Started (PDF)")
BOOK1=HLP\mdk5-getting-started_jp.pdf ("MDK-ARM Getting Started (Japanese/PDF)")
BOOK2=HLP\RELEASE_NOTES.HTM ("Release Notes for MDK 5.38")
BOOK3=HLP\ARMTOOLS.chm("Complete User's Guide Selection", C)
BOOK4=ARMCLANG\sw\info\releasenotes.html("Release Notes for Arm Compiler 6.19", GEN)
BOOK5=ARMCLANG\sw\hlp\compiler_user_guide.pdf("Arm Compiler User Guide Version 6.19 (PDF)",GEN)
BOOK6=ARMCLANG\sw\hlp\compiler_reference_guide.pdf("Arm Compiler Reference Guide Version 6.19 (PDF)",GEN)
BOOK7=ARMCLANG\sw\hlp\migration_and_compatibility_guide.pdf("Arm Compiler Migration and Compatibility Version 6.19 (PDF)",GEN)
BOOK8=ARMCLANG\sw\hlp\errors_and_warnings_reference_guide.pdf("Arm Compiler Errors and Warnings Reference Guide Version 6.19 (PDF)",GEN)
BOOK9=ARMCLANG\sw\hlp\libraries_user_guide.pdf("Arm Compiler Arm C and C++ Libraries and Floating-Point Support User Guide Version 6.19 (PDF)",GEN)
BOOK10=ARMCLANG\sw\hlp\arm_instruction_set_reference_guide.pdf("Arm Instruction Set Reference Guide Version 1.0 (PDF)", GEN)
BOOK11=ARMCLANG\sw\hlp\instruction_set_assembly_guide_for_armv7_and_earlier_arm_architectures.pdf("Instruction Set Assembly Guide Armv7 and earlier Version 2.0 (PDF)",GEN)
TDRV0=BIN\UL2CM3.DLL("ULINK2/ME Cortex Debugger")
TDRV1=BIN\ULP2CM3.DLL("ULINK Pro Cortex Debugger")
TDRV2=BIN\ULPL2CM3.dll("ULINKplus Debugger")
TDRV3=BIN\CMSIS_AGDI.dll("CMSIS-DAP Debugger")
TDRV4=Segger\JL2CM3.dll("J-LINK / J-TRACE Cortex")
TDRV5=BIN\DbgFM.DLL("Models Cortex-M Debugger")
TDRV6=STLink\ST-LINKIII-KEIL_SWO.dll ("ST-Link Debugger")
TDRV7=NULink\Nu_Link.dll("NULink Debugger")
TDRV8=PEMicro\Pemicro_ArmCortexInterface.dll("Pemicro Debugger")
TDRV9=SiLabs\SLAB_CM_Keil.dll("SiLabs UDA Debugger")
TDRV10=BIN\ABLSTCM.dll("Altera Blaster Cortex Debugger")
TDRV11=TI_XDS\XDS2CM3.dll("TI XDS Debugger")
TDRV12=BIN\ULP2V8M.DLL("ULINK Pro ARMv8-M Debugger")
TDRV13=BIN\UL2V8M.DLL("ULINK2/ME ARMv8-M Debugger")
TDRV14=BIN\CMSIS_AGDI_V8M.DLL("CMSIS-DAP ARMv8-M Debugger")
TDRV15=BIN\DbgFMv8M.DLL("Models ARMv8-M Debugger")
TDRV16=BIN\UL2ARM.DLL("ULINK2/ME ARM Debugger")
TDRV17=BIN\ULP2ARM.DLL("ULINK Pro ARM Debugger")
TDRV18=Segger\JLTAgdi.dll("J-LINK / J-TRACE ARM")
RTOS0=Dummy.DLL("Dummy")
RTOS1=VARTXARM.DLL ("RTX Kernel")
ARMCCPATH0=".\ARMCLANG" ("V6.19.0")
DELDRVPKG0=ULINK\UninstallULINK.exe("ULINK Pro Driver V1.0")
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
