package seo

import (
	"reflect"
	"fmt"
)

func InvokeObjectMethod( object interface{},  methodName string ,args ...interface{}  ){
	t := reflect.TypeOf( object )
	v := reflect.ValueOf( object )

	//fmt.Println(t)
	//fmt.Println(v)


	// 获取i所指向的对象的类型
	structType := t.Elem()
	// 获取对象的名字
	structName := structType.Name()
	fmt.Println( "Invoke:: structname ::" +  structName)

	//通过t获取对象方法的信息
	//method, _ := t.MethodByName( methodName )
	//fmt.Println(  "Invoke:: methodname ::" +   method )


	/* 通过v获取对象方法的信息 */
	method := v.MethodByName( methodName )
	//fmt.Println( method.Pointer() )
	//fmt.Println(method.Type()  )
	//fmt.Println(method.CanSet())

	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	// 通过v进行调用
	method.Call(inputs)




	//reflect.ValueOf(object).MethodByName(methodName).Call(   )

}