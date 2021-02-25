package main

import (
	"fmt"
	"jvmgo/classpath"
	"log"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v ,class:%v,args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	data, _, err := cp.ReadClass(className)
	if err != nil {
		log.Fatal("无法找到或加载主类", err)
	}
	fmt.Printf("class byte code:%v \n", data)

}
