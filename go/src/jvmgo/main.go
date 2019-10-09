package main

import "fmt"
import "strings"
import "jvmgo/classpath"

func main()  {
	cmd := ParseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
    className := strings.Replace(cmd.class, ".", "/", -1)
    data, _, _ := cp.ReadClass(className)
    if data != nil {
    	fmt.Printf("class data:%v\n", data)
	} else {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
	}
}

