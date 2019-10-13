package main

import (
	"fmt"
	"jvmgo/classfile"
)
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
    	cf, err := classfile.Parse(data)
    	if err == nil {
           printClassInfo(cf)
		}
	} else {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
	}
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion, cf.MinorVersion)
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.FieldInfos))
	for _, f := range cf.FieldInfos {
		fmt.Printf("  %s\n", f.Name(&cf.ConstantPool))
	}
	fmt.Printf("methods count: %v\n", len(cf.MethodInfos))
	for _, m := range cf.MethodInfos {
		fmt.Printf("  %s\n", m.Name(&cf.ConstantPool))
	}
}