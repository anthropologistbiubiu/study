package cmd

improt(
	"fmt"
)
func join(strs string)(string,error){
	strArr := strings.Split(strs)
	var result string
	for id,value := range strArr {
		if value != "-"{
			continue
		}else{
			result += value	
		}
	} 
	return result,nil
}