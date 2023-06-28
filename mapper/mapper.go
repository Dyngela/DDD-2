package mapper

import (
	"awesomeProject/pkg/user"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Sender struct {
	Name string
	Pat  *string
}

type Receiver struct {
	Name string
	Pat  *string
}

func GenerateMapping() {
	/*
		get the path of the interface used
		erase current file previously generated
		get through the method declaration and get the input struct and output struct needed
		get the package of each struct
		handle import
		generate method for each declaration preceded my (self *mapper)
	*/
	var str string

	myinterface := user.UserMapper()
	elements := reflect.TypeOf(&myinterface).Elem()

	pkgPath := elements.PkgPath()
	temp := strings.Split(pkgPath, "/")
	pkgName := temp[len(temp)-1]
	log.Println(pkgName)

	str = fmt.Sprintf("package %s\n\n", pkgName)

	var methods []reflect.Method
	for i := 0; i < elements.NumMethod(); i++ {
		methods = append(methods, elements.Method(i))
		//Get the type of param and type of return. Always 0 for our mapper use case
		in := methods[i].Type.In(0)
		out := methods[i].Type.Out(0)
		// we get the name properly formatted of our Type in param and of return
		inputParamStr := strings.Split(in.String(), ".")[1]
		outputParamStr := strings.Split(out.String(), ".")[1]
		methodName := methods[i].Name
		str = str + fmt.Sprintf("func (self *mapper) %s(input %s) %s {\n\tvar out %s\n",
			methodName, inputParamStr, outputParamStr, outputParamStr)

		for inIndex := 0; inIndex < in.NumField(); inIndex++ {
			log.Println(in.Field(inIndex))
			for outIndex := 0; outIndex < out.NumField(); outIndex++ {
				log.Println(out.Field(outIndex))
			}
		}
	}
	log.Println(str)
}
