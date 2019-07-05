package main
import "fmt"
import "time"
import   "strconv"
import 	"os"
import     "os/exec"



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
	int_chr,err :=strconv.Atoi(character)
	if err == nil{
		font_convert= font[int_chr]
	}
	return font_convert
}


//input: string 12:03:07 and font
//output: array[[],[],..[]] of font
func convert_time_strings_to_fonts (time_str string,font [][]string) ([][]string){
	var array [][]string
	for i:=0;i<len(time_str);i++{
		str:=time_str[i]
		str1:=strconv.Itoa(int(str)-48)
		arr:=covert_1_character_to_1_font(str1,font)
		array=append(array,arr)
		
	}
	return array
}


//input: arr[[],[],..,[]]
//output: merge_ar["..","..",..]
func merge_array (array [][]string ) ([]string){
	merge :=[]string{"","",""}
	for i:=0; i< len(array[0]);i++ {
		for j:=0;j<len(array);j++ {
			font_j :=array[j]
			merge[i]+=font_j[i]
		}

	}
	return merge
}

func display_merge_font (merge []string){
	for i:=0; i<len(merge);i++{
		fmt.Println(merge[i])
	}
}

// clear terminal
func clear_terminal(){
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
var pre_sec=0
func main(){
	for true{
		time_string :=get_time_clock_of_os()
		time_second:=int(time_string[6]-48)*10+int(time_string[7]-48)
		if pre_sec != time_second{
			clear_terminal()
			fmt.Println(time_string)
			convert_font := convert_time_strings_to_fonts(time_string,font)
			merge:=merge_array(convert_font)
			display_merge_font(merge)
			pre_sec = time_second
		}
		time.Sleep(1*time.Millisecond)
	}
}