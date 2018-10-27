package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "os"
        "strings"
        "io/ioutil"
        "regexp"
        "github.com/tidwall/sjson"
        "strconv"
)

var paths []string
func printPaths(v interface{}, key, path string) {
        switch v := v.(type) {
        case map[string]interface{}:
                for mk, mv := range v {
                        p := path + "." + mk
                        if mk == key {
                                paths=append(paths,p)
                        }
                        printPaths(mv, key, p)
                }
        case []interface{}:
                for i, sv := range v {
                        printPaths(sv, key, fmt.Sprintf("%s.%d", path, i))
                }
        }
}

func iteratePath(data,keyvalue,modifyvalue,typeofvalue string) (string){
        key := keyvalue
        paths = nil
         var re = regexp.MustCompile(`\s`)
        s := re.ReplaceAllString(data, ``)
                var v interface{}
                b := []byte(s)
                if err := json.Unmarshal(b, &v); err != nil {
                        fmt.Println("JSON parsing error: %s\n", err.Error())
                }
                printPaths(v, key, "")
     if len(paths)==0{
             fmt.Println("Key is not present in the json, Enter valid key")
                        os.Exit(0)
     }
     var err error
     json:=data
     for i,value  := range paths{
         x:= strings.TrimPrefix(value, ".")
         paths[i]=x
          if typeofvalue=="NUMBER"{
                  f, err1 := strconv.ParseFloat(modifyvalue, 64)
                  if err1!=nil{
                        fmt.Println("value and it's DataType doesn't match")
                        os.Exit(0)
                  }
                  json,err=sjson.Set(json,x,f)
          }else if typeofvalue=="NULL"{
                  json,err=sjson.Set(json,x,nil)
          }else if typeofvalue=="BOOL"{
                  b, err1 := strconv.ParseBool(modifyvalue)
                  if err1!=nil{
                        fmt.Println("value and it's DataType doesn't match")
                        os.Exit(0)
                  }
                  json,err=sjson.Set(json,x,b)
          }else{
                  json,err=sjson.Set(json,x,modifyvalue)
           }
          if err!=nil{
                    fmt.Println(" error while modifying the value")
                    os.Exit(0)
	    }
     }
 return json
}

func main(){
       flag.Usage = usage
        flag.Parse()
        if flag.NArg() != 4 {
                usage()
                os.Exit(1)
        }
        key := os.Args[1]
        modifiedValue := os.Args[2]
        typeOfValue := os.Args[3]
	typeOfVlaue = strings.ToUpper(typeOfValue)
        filename := os.Args[4]
        dat, err := ioutil.ReadFile(filename)
        check(err)
        out:=iteratePath(string(dat),key,modifiedValue,typeOfValue)
        err=ioutil.WriteFile("output.json",[]byte(out),0644)
        check(err)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func usage() {
        fmt.Println("one of the mandatory arguments are missing <key> <modifiedValue> <Type of Value> <Filename>")
}
