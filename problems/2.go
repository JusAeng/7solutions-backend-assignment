package problems

import (
	"fmt"
	"strconv"
)

func LeftRightEqual() string {
	var xcode string
	fmt.Scanln(&xcode)	
	var result string
	allPoint := []int{0}
	vertex := []int{0} 
	latest := ""
	point := 0
	shiftValue := point
	for i := 0; i < len(xcode); i++ {
		x := string(xcode[i])
		if x != "=" {
			// toggle
			if latest != x {
				if (latest == "" && x == "R"){
					vertex = append(vertex, 0)
				}
				vertex = append(vertex, i)
			}
			latest = x
		}
		if x == "L"{
			point -= 1
		}else if(x == "R"){
			point += 1
		}
		allPoint = append(allPoint, point)

		// min point
		if (point < shiftValue) {
			shiftValue = point
		}
		// end point
		if(i == len(xcode)-1){
			if (latest == "L"){
				vertex = append(vertex, len(xcode))
			}else if (latest == "R"){
				vertex = append(vertex, len(xcode),len(xcode))
			}
		}
	}	
	shiftGraph(allPoint,shiftValue*-1)		

	for i := 2; i < len(vertex); i+=2 {
		prevVdown := vertex[i-2]
		prevVup := vertex[i-1]
		v := vertex[i]
		var temp []int
		temp = append(temp,allPoint[prevVdown:v+1]...)		
		tempShift := 0
		if allPoint[prevVdown] >= allPoint[v]{
			// ขวายาวกว่าหรือเท่ากับ 
			tempShift = allPoint[v]
			shiftGraph(temp,tempShift*-1)
			
			if(prevVdown == 0){ 
				result += packToString(temp)				
			}else{
				result += packToString(temp[1:])				
			}
		}else{
			// ซ้ายยาวกว่า (ซ้าย 0 แน่นอน)
			tempShift = allPoint[prevVdown]
			shiftGraph(temp,tempShift*-1)			
			result += packToString(temp[1:prevVup-prevVdown+1])
			
			// ดึกก้านข้าว
			tempShift = temp[len(temp)-1]
			shiftGraph(temp[prevVup-prevVdown:],tempShift*-1)			
			result += packToString(temp[prevVup-prevVdown+1:])
		}
	}
	fmt.Println(result)	
	return result
}

func shiftGraph(slice []int, numToAdd int) {
    for i := range slice {
        slice[i] += numToAdd
    }
}
func packToString(slice []int) string {
    result := ""
    for _, num := range slice {
        result += strconv.Itoa(num)
    }
    return result
}