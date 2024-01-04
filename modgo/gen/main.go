package main

import "C"
import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "add" {
		fmt.Println("Usage: go run main.go add <function_name> --client=<client_function>")
		os.Exit(1)
	}

	functionName := os.Args[2]
	clientFunction := os.Args[3]

	if err := findFiles(functionName, clientFunction); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("generate code success 🎉🎉🎉")

}

// findGoFiles finds all Go files in the specified directory and its subdirectories.
func findGoFiles(rootDir string) ([]string, error) {
	var goFiles []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return goFiles, nil
}
func ttt(field interface{}) (pkg string, name string) {
	switch fieldType := field.(type) {
	case *ast.SelectorExpr:
		// fmt.Println("Found return type:", fieldType.X.(*ast.Ident).Name+"."+fieldType.Sel.Name)
		return fieldType.X.(*ast.Ident).Name, fieldType.Sel.Name

	case *ast.StarExpr:
		return ttt(fieldType.X)
		// if ident, ok := fieldType.X.(*ast.Ident); ok {
		// 	fmt.Println("Found return type:", ident.Name)
		// }
	case *ast.Field:
		return ttt(fieldType.Type)
		// return "", fieldType.Type.(*ast.Ident).Name
	case *ast.ReturnStmt:
		if len(fieldType.Results) != 0 {
			return ttt(fieldType.Results[0])
		}
	case *ast.CompositeLit:
		return ttt(fieldType.Type)
	case *ast.UnaryExpr:
		return ttt(fieldType.X)
	case *ast.Ident:
		return "", fieldType.Name
	}
	return "", ""
}

func findFiles(functionName, clientFunction string) error {
	funcArr := strings.Split(functionName, ".")
	ClientArr := strings.Split(clientFunction, ".")
	if len(funcArr) != 2 || len(ClientArr) != 2 {
		return fmt.Errorf("functionName or clientFunction error")
	}
	rootDir := "./"
	originPath := ""
	targetPath := ""
	// targetFuncRespPath := ""
	goFiles, err := findGoFiles(rootDir)
	if err != nil {
		fmt.Println("Error finding Go files:", err)
		return err
	}
	gomod := GetGoMod()
	fset := token.NewFileSet()
	var fileTar *ast.File
	var targetFunc *ast.FuncDecl
	var targetStruct *ast.GenDecl
	var targetStructNum int
	var originClient *ast.FuncDecl
	// 目标函数的返回结构体名称
	identName := ""
	identPkgName := funcArr[0]

	// ss, vv, dds := findStructResp(goFiles, funcArr[1], funcArr[0])
	// fmt.Println(ss, vv, dds)
	// 将当前mod下所有文件搜集到，一一检查
	over := 0
	i := 0
	for over < 2 {
		goFile := goFiles[i]
		i++
		if len(goFiles) == i {
			i = 0
			over++
		}

		if targetFunc != nil && originClient != nil && targetStruct != nil {
			break
		}
		// Parse the source code of the current file
		file, err := parser.ParseFile(fset, goFile, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Error parsing file %s: %v\n", goFile, err)
			continue
		}

		if file.Name.Name == funcArr[0] {
			for _, decl := range file.Decls {

				// 找到目标函数func
				if f, ok := decl.(*ast.FuncDecl); ok && f.Name.Name == funcArr[1] && targetFunc == nil {
					// 获取func真实返回的 结构信息，目的是给这个结构增加成员变量
					pkg, name := ttt(f.Body.List[0])
					// ss := f.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.UnaryExpr).X.(*ast.CompositeLit).Type.(*ast.Ident).Name
					// 获取func 形参返回的结构
					// pkg, name := ttt(f.Type.Results.List[0])
					identName = name

					if pkg != "" {
						identPkgName = pkg
						// identName = pkg + "." + name
					}
					targetFunc = f
					fileTar = file
					targetPath = goFile
				}

			}

		}
		// 找到目标函数func 的返回结构体，可能当前包，也可能在其他包内
		if file.Name.Name == ClientArr[0] {
			for _, decl := range file.Decls {
				if f, ok := decl.(*ast.FuncDecl); ok && f.Name.Name == ClientArr[1] && originClient == nil {
					originClient = f
					originPath = getPathFile(goFile)
				}
			}
		}

		// 找到目标函数func 的返回结构体，可能当前包，也可能在其他包内

		// 重复第二遍目的是为了找到这个文件中结构体结构
		if file.Name.Name == identPkgName && len(identName) != 0 {
			for _, decl := range file.Decls {

				if f, ok := decl.(*ast.FuncDecl); ok {
					if f.Name.Name == identName {
						targetFunc = f
						if f.Type.Results != nil && len(targetFunc.Type.Results.List) > 0 {
							returnType := f.Type.Results.List[0].Type
							// targetStruct = returnType
							_ = returnType
						}

						break
					}
				}
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for i, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.Name == identName {
							// Check if the type is a struct

							if _, ok := typeSpec.Type.(*ast.StructType); ok {
								// 怼到这一层 目的是校验确实有这个struct
								targetStruct = genDecl
								targetStructNum = i
							}
						}
					}
				}
			}
		}
	}

	if targetFunc == nil {
		return fmt.Errorf("Function %s not found", functionName)
	}
	if originClient == nil {
		return fmt.Errorf("Client %s not found", clientFunction)
	}

	// 向目标函数添加参数
	for _, CC := range originClient.Type.Results.List[:1] {
		// 校验CC是否返回接口类型
		name := ""
		star := ""
		pkg, name := ttt(CC.Type)
		if pkg == "" && ClientArr[0] != funcArr[0] {
			pkg = ClientArr[0]
		}
		if _, ok := CC.Type.(*ast.StarExpr); ok {
			// 	name = fmt.Sprintf("%s%s", ClientArr[0], CC.Type.(*ast.StarExpr).X.(*ast.Ident).Name)
			star = "*"
		}

		tTypeName := ""
		if pkg == "" {
			tTypeName = star + name
		} else {
			tTypeName = star + pkg + "." + name
		}
		tType := &ast.Ident{Name: tTypeName}
		targetFunc.Type.Params.List = append(targetFunc.Type.Params.List, &ast.Field{
			Names: []*ast.Ident{{Name: pkg + name}},
			Type:  tType,
		})
		// Add a new field to the struct
		if targetStruct != nil {
			targetStruct.Specs[targetStructNum].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List = append(targetStruct.Specs[targetStructNum].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List, &ast.Field{
				Names: []*ast.Ident{{Name: name}},
				Type:  tType,
			})
		}
	}
	targetFunc.Type.Params.Closing = targetFunc.Type.Params.List[len(targetFunc.Type.Params.List)-1].End()

	if !hasImport(fileTar, fmt.Sprintf("%s/%s", gomod, originPath)) && ClientArr[0] != funcArr[0] {
		// 添加新的依赖导入语句
		newImport := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s/%s\"", gomod, originPath),
			},
		}

		// 创建一个新的导入声明
		importDecl := &ast.GenDecl{
			Tok:    token.IMPORT,
			Specs:  []ast.Spec{newImport},
			Lparen: 1, // 设置为1，表示有括号
		}
		// 将导入声明添加到文件的第一个位置
		fileTar.Decls = append([]ast.Decl{importDecl}, fileTar.Decls...)
	}

	// 打开目标文件，使用os.O_TRUNC标志来覆盖写入
	outputFile, err := os.Create(targetPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}
	defer outputFile.Close()
	// 将AST节点格式化为Go代码并写入文件
	if err := format.Node(outputFile, fset, fileTar); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return err
}
func findStructResp(goFiles []string, name, pkg string) (structPkg, structName, structPath string) {
	fset := token.NewFileSet()

	for _, goFile := range goFiles {
		file, err := parser.ParseFile(fset, goFile, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Error parsing file %s: %v\n", goFile, err)
			continue
		}
		if file.Name.Name == pkg {
			for _, decl := range file.Decls {
				// 找到目标函数func
				if f, ok := decl.(*ast.FuncDecl); ok && f.Name.Name == name {

					// fileTar = file
					if fieldType, ok := f.Type.Results.List[0].Type.(*ast.StarExpr); ok {
						// ident, ok := fieldType.X.(*ast.Ident)
						// fmt.Println(ident, ok)

						if ident, ok := fieldType.X.(*ast.Ident); ok {
							// 给这个返回结构体 赋值 变量，添加整体结构
							// identName = ident.Name
							// targetPath = goFile
							return "", ident.Name, goFile
						} else if ident1, ok1 := fieldType.X.(*ast.SelectorExpr); ok1 {
							// fmt.Println(ident1, ok1)
							return ident1.X.(*ast.Ident).Name, ident1.Sel.Name, goFile
							// 给这个返回结构体 赋值 变量，添加整体结构
							// identName = ident1.X.(*ast.Ident).Name + "." + ident1.Sel.Name
							// targetFuncRespPath = ""
						}
					}

				}

			}

		}
	}
	return

}

// hasImport 检查是否已导入指定的包
func hasImport(file *ast.File, packageName string) bool {
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			for _, spec := range genDecl.Specs {
				if importSpec, ok := spec.(*ast.ImportSpec); ok {
					importedPath := strings.Trim(importSpec.Path.Value, `"`)
					if importedPath == packageName {
						return true
					}
				}
			}
		}
	}
	return false
}
func getPathFile(path string) string {
	// 使用 strings.LastIndex 找到最后一个斜杠的索引
	lastIndex := strings.LastIndex(path, "/")

	// 如果找到了斜杠，则截取字符串
	if lastIndex != -1 {
		result := path[:lastIndex]
		return result
	} else {
		// 如果没有找到斜杠，则整个字符串都是结果
		return path
	}
}

func GetGoMod() string {

	// 获取当前文件所在目录
	dir := filepath.Dir("./")

	// 寻找包含 go.mod 文件的父目录
	goModPath, err := findGoMod(dir)
	if err != nil {
		fmt.Printf("Error finding go.mod file: %v\n", err)
		return ""
	}
	// 读取 go.mod 文件内容
	content, err := os.ReadFile(goModPath)
	if err != nil {
		fmt.Printf("Error reading go.mod file: %v\n", err)
		return ""
	}
	// 解析 go.mod 文件
	modFile, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		fmt.Printf("Error parsing go.mod file: %v\n", err)
		return ""
	}

	return modFile.Module.Mod.Path

}

func findGoMod(dir string) (string, error) {
	// 从当前目录开始，逐级向上查找 go.mod 文件
	for {
		goModPath := filepath.Join(dir, "go.mod")
		_, err := os.Stat(goModPath)
		if err == nil {
			return goModPath, nil // 找到 go.mod 文件
		}

		// 如果到达根目录仍未找到，返回错误
		if dir == filepath.Dir(dir) {
			return "", fmt.Errorf("go.mod file not found")
		}

		dir = filepath.Dir(dir)
	}
}

// 找到目标结构体用于添加对应形参
func findStruct(goFiles []string, name, pkg string) (*ast.GenDecl, int) {
	fset := token.NewFileSet()
	for _, goFile := range goFiles {
		// Parse the source code of the current file
		file, err := parser.ParseFile(fset, goFile, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Error parsing file %s: %v\n", goFile, err)
			continue
		}
		// 找到目标函数func 的返回结构体，可能当前包，也可能在其他包内

		// 重复第二遍目的是为了找到这个文件中结构体结构
		if file.Name.Name == pkg && len(name) != 0 {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for i, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.Name == name {
							// Check if the type is a struct
							if _, ok := typeSpec.Type.(*ast.StructType); ok {
								// 怼到这一层 目的是校验确实有这个struct
								// targetStruct = genDecl
								// targetStructNum = i
								return genDecl, i
							}
						}
					}
				}
			}
		}
	}
	return nil, 0
}
