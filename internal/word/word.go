package word

import (
	"strings"	
	"unicode"
)


func ToUpper(s string)  string{
	return strings.ToUpper(s)
}

func ToLower(s string)  string{
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string)string{
	s = strings.Replace(s,"_"," ",-1)
	s = strings.Title(s)
	return strings.Replace(s," ","",-1)
}

func CamelCase(s string)string{
	var ans []rune
	for k,v := range s{
		if k == 0{
			ans = append(ans,unicode.ToLower(v))
			continue
		}
		if unicode.IsUpper(v){
			ans = append(ans,'_')
		}
		ans = append(ans,unicode.ToLower(v))
	}
	return string(ans)
}

