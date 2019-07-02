
num_0 = [" _ _ ",
         "|   |",
         "|_ _|"]
num_1 = ["     ",
         " /|  ",
         "  |  "]
num_2 = [" _ _ ",
         " _ _|",
         "|_ _ "]
num_3 = ["_ _  ",
         "_ _| ",
         "_ _| "]
num_4 = ["     ",
         "|_ _|",
         "    |"]
num_5 = [" _ _ ",
         "|_ _ ",
         " _ _|"]
num_6 = [" _ _ ",
         "|_ _ ",
         "|_ _|"]
num_7 = ["_ _  ",
         "  /  ",
         " /   "]
num_8 = [" _ _ ",
         "|_ _|",
         "|_ _|"]
num_9 = [" _ _ ",
         "|_ _|",
         " _ _|"]
symbol =["     ",
         "  %  ",
         "  %  "]

font = [num_0, num_1, num_2, num_3, num_4, num_5, num_6, num_7, num_8, num_9, symbol]

#input: clock of system
#output:string clock
def get_string_clock_of_system
    time_now=Time.now.to_s
    time_string=time_now[11..18]
    return time_string
end

#input: 1 character and font
#output: arr_font of character
def convert_1_character_to_1_font(character,font)
    numbers="0123456789" #array to check number
    linh_canh=0                 # flag
    for i in 0..9 
        if character==numbers[i]
            linh_canh=1
        end
    end
    if linh_canh==1
        return font[character.to_i]
    else
        return font[-1]
    end 
    
end

#input: string 12:03:07 and font
#output: array[[],[],..[]] of font
def convert_time_string_to_array_font (time_string,font) 
    array=[]
    for i in 0..(time_string.length-1) 
        convert=convert_1_character_to_1_font(time_string[i],font)
        array.push(convert)
    end
    return array
    
end

#input: arr[[],[],..,[]]
#output: merge_ar["..","..",..]
def merge_arr (arr)
    merge=["","","","","","","","","",""]
    for i in 0..(arr[0].length-1) 
        for j in 0..(arr.length-1) 
            font=arr[j]
            merge[i]+=font[i]
        end
    end
    return merge
end

#display_merge_array
def display_merge (merge)
    puts merge
end

pre_sec=0
while true
    time_string=get_string_clock_of_system
    if time_string[6..7].to_i  != pre_sec
        system("clear")
        puts Time.now
        convert=convert_time_string_to_array_font(time_string,font)
        merge=merge_arr(convert)
        display_merge(merge)
        pre_sec = time_string[6..7].to_i
    end 
    sleep(0.01)    
end

