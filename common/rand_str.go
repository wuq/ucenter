package common

import (
    "math/rand"
    "time"
    "strings"
    "bytes"
)


func RandStr(randLength int, randType string) (result string) {
    var num string = "0123456789"
    var lower string = "abcdefghijklmnopqrstuvwxyz"
    var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    b := bytes.Buffer{}
    if strings.Contains(randType, "0") {
        b.WriteString(num)
    }
    if strings.Contains(randType, "a") {
        b.WriteString(lower)
    }
    if strings.Contains(randType, "A") {
        b.WriteString(upper)
    }
    var str = b.String()
    var strLen = len(str)
    if strLen == 0 {
        result = ""
        return
    }

    rand.Seed(time.Now().UnixNano())
    b = bytes.Buffer{}
    for i := 0; i < randLength; i++ {
        b.WriteByte(str[rand.Intn(strLen)])
    }
    result = b.String()
    return
}


/*

fmt.Println(RandomString(8, "A"))
fmt.Println(RandomString(8, "a0"))
fmt.Println(RandomString(20, "Aa0"))

*/