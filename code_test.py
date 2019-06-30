import  datetime 
import time
import os
pre_sec=0


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
#input: clock_of_sourse
#output: time_string
def get_string_of_clock() :
    time_now= datetime.datetime.now()
    time_string=str(time_now)[11:19]
    return time_string


#input: 1 ky tu
#output: font mong muon
def covert_1_character_to_1_font (character,font):
        if character.isdigit():
            return font[int(character)]
        else:
            return symbol
#

#int: time_string, font
#output: array_font
def convert_time_string_to_array_font (time_string,font):
    array=[]
    for i in range(len(time_string)) :
        convert=covert_1_character_to_1_font (time_string[i],font)
        array.append(convert)
    return array


# input: array 1 chieu phan tu rieng le
#output: array 1 chieu cac phan tu da duoc hop nhat
def merged_array (array) :
    merged=["","",""]
    for i in range(len(array[0])) :
        for j in range(len(array)) :
            b=array[j]
            merged[i-1]+=b[i-1]
    return merged


def display_merged (merged) :
    for i in range(len(merged)) :
        print(merged[i])


while True :
    time_now= datetime.datetime.now()
    if time_now.second != pre_sec:
        os.system("clear")
        print(time_now)
        time_string=get_string_of_clock()
        converted = convert_time_string_to_array_font (time_string,font)
        merged = merged_array (converted)
        display_merged (merged)
        pre_sec=time_now.second
    time.sleep(0.001)
