package ini2go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Han-Ya-Jun/json2go"
	"github.com/go-ini/ini"
	"os"
	"path"
	"strings"
	"unicode"
)

/*
* @Author:hanyajun
* @Date:2019/7/27 21:18
* @Name:ini2struct
* @Function: ini文件自动生成struct
 */

func Ini2Go(iniFileName string, pkgName string, goFileName string, outputPath string, writeTag bool, tagKeys []string) error {
	ext := path.Ext(iniFileName)
	if ext != ".ini" {
		return errors.New("file format err(must be end with .ini)")
	}
	config, err := ini.Load(iniFileName)
	if err != nil {
		fmt.Printf("load ini file err:%v", err)
		return err
	}
	sectionsList := config.Sections()
	_, err = os.Stat(outputPath + goFileName)
	if err == nil {
		_ = os.Remove(outputPath + goFileName)
	}
	// 以追加模式打开文件，当文件不存在时生成文件
	outputFile, err := os.OpenFile(outputPath+goFileName, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	for index, section := range sectionsList {
		if index == 0 {
			continue
		}
		jsonMap := make(map[string]interface{})
		for _, key := range section.Keys() {
			keyName := key.Name()
			if strings.ContainsAny(keyName, "[") && strings.Contains(keyName, "]") {
				keyName = strings.Replace(keyName, "[", ",", 1)
				keyName = strings.Replace(keyName, "]", "", 1)
				keyFormat := strings.Split(keyName, ",")
				switch keyFormat[1] {
				case "int":
					jsonMap[changeStringToJsonFormat(keyFormat[0])], err = key.Int()
					if err != nil {
						fmt.Printf("key:%v value format err:%v\n", keyFormat[0], err)
						panic(err)
					}
				case "int64":
					jsonMap[changeStringToJsonFormat(keyFormat[0])], err = key.Int()
					if err != nil {
						fmt.Printf("key:%v value format err:%v\n", keyFormat[0], err)
						panic(err)
					}
				}
			} else {
				jsonMap[changeStringToJsonFormat(keyName)] = key.String()
			}
		}
		b, _ := json.Marshal(jsonMap)
		r := bytes.NewReader(b)
		var buff bytes.Buffer
		calvin := json2go.NewTransmogrifier(section.Name(), r, &buff)
		_ = calvin.SetTagKeys(tagKeys)
		calvin.WriteTag = writeTag
		if index > 1 {
			calvin.WritePkg = false
		} else {
			calvin.SetPkg(pkgName)
			calvin.WritePkg = true
		}
		//calvin.WriteJSON=true
		calvin.SetStructName(section.Name())
		err = calvin.Gen()
		if err != nil {
			fmt.Printf("unexpected error: %s", err)
			return err
		}
		_, err = outputFile.Write(append(buff.Bytes(), []byte("\n")...))
		if err != nil {
			return err
		}
	}
	return nil
}

func changeStringToJsonFormat(str string) string {
	runeList := []rune(str)
	var result []rune
	for i, r := range runeList {
		if unicode.IsUpper(r) {
			if i != 0 {
				result = append(result, []rune("_")...)
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
