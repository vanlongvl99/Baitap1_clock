package main
import (
	"fmt"
	"time"
   "strconv"
 	"os"
	"os/exec"
	"strings"
)
var font =[][]string {
	[]string {" _ _ ",
			  "|   |",
			  "|_ _|"},
	[]string {"     ",
			  " /|  ",
			  "  |  "},
	[]string {" _ _ ",
         	  " _ _|",
         	  "|_ _ "},
	[]string {"_ _  ",
			  "_ _| ",
			  "_ _| "},		  
	[]string {"     ",
			  "|_ _|",
			  "    |"},
	[]string {" _ _ ",
			  "|_ _ ",
			  " _ _|"},
	[]string {" _ _ ",
			  "|_ _ ",
			  "|_ _|"},
	[]string {"_ _  ",
			  "  /  ",
			  " /   "},
	[]string {" _ _ ",
			  "|_ _|",
			  "|_ _|"},
	[]string {" _ _ ",
			  "|_ _|",
			  " _ _|"},
	[]string {"     ",
			  "  %  ",
			  "  %  "},

}
//input: clock of system
//output:string clock
func get_time_clock_of_os()string{
	var time_now = time.Now().String()
	return time_now[11:19]
}
//input: 1 character and font
//output: arr_font of character
func covert_1_character_to_1_font(character string,font [][]string) ([]string){
	var font_convert []string
	if strings.Contains("0123456789",character){
		int_chr,err :=strconv.Atoi(character)
		if err == nil{
			font_convert= font[int_chr]
		}
	} else{
		font_convert= font[10]
	}
	return font_convert
}
//input: string 12:03:07 and font
//output: array[[],[],..[]] of font
func convert_time_strings_to_fonts (time_str string,font [][]string) ([][]string){
	var array [][]string
	characters := strings.Split(time_str, "") 
	for _, character := range characters{
		arr:=covert_1_character_to_1_font(character,font)
		array=append(array,arr)		
	}
	return array
}
//input: arr[[],[],..,[]]
//output: merge_ar["..","..",..]
func merge_arr_of_fonts (fonts [][]string ) ([]string){
	merge :=[]string{"","",""}
	for _, font := range fonts {
		for row_index, row := range font {
			merge[row_index] += row
		}
	}
	return merge
}
//display
func display_merge_font (merge []string){
	for _, item := range merge{
		fmt.Println(item)
	}
}
// clear terminal
func clear_terminal(){
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
func main(){
	var pre_sec=0
	for {
		time_now:=time.Now()
		time_second:=time_now.Second()		
		if pre_sec != time_second{
			clear_terminal()
			fmt.Println(time_now)
			time_string :=get_time_clock_of_os()
			fmt.Println(time_string)
			convert_font := convert_time_strings_to_fonts(time_string,font)
			merge:=merge_arr_of_fonts(convert_font)
			display_merge_font(merge)
			pre_sec = time_second
		}
		time.Sleep(1*time.Millisecond)
	}
}