
import  datetime 
import time
import os
pre_sec=0

#1   $$$$$$||    # ví dụ minh họa là số 0 và có kích thước 10x8, các hàng sẽ  
#2  $$     $||   # đường chuyển thành 1 hàng ngang có chiều dài 80
#3  $$     $$|
#4  $$     $$|
#5  $$     $$|
#6  $$     $$|
#7  | $$$$$$ |
#8   |______| ''',
#   0123456789

ky_tu=["","","","","","","","","",""]
ky_tu[0]=" $$$$$$|| $$     $||$$     $$|$$     $$|$$     $$|$$     $$|| $$$$$$ | |______| "
ky_tu[1]="  $$||    $$$$ |    |_$$ |      $$ |      $$ |      $$ |    $$$$$$||  |______|  "
ky_tu[2]="  $$$$$$||$$  __$$|||__|  $$ | $$$$$$  |$$  ____| $$ |      $$$$$$$$|||________|"
ky_tu[3]=" $$$$$$|| $$ ___$$        $$ |  $$$$$ |   |___$$||      $$ ||$$$$$$  | |______| "
ky_tu[4]="$$|   $$||$$ |  $$ |$$ |  $$ |$$$$$$$$ ||_____$$ |      $$ |      $$ |      |__|"
ky_tu[5]="$$$$$$$|| $$  ____| $$ |      $$$$$$$|| |_____$$||       $$||$$$$$$  | |______| "
ky_tu[6]=" $$$$$$|| $$    $$||$$        $$$$$$$|| $$````$$||$$    $$ |$$$$$$$  | |______| "
ky_tu[7]="$$$$$$$$|||____$$  |    $$  |    $$  |    $$  |    $$  |    $$  |     |__|      "
ky_tu[8]=" $$$$$$|| $$    $$||$$     $$|$$$$$$$$ |$$     $$|$$     $$||$$$$$$  | |______| "
ky_tu[9]=" $$$$$$|| $$    $$||$$    $$ ||$$$$$$$ |      $$ |      $$ ||$$$$$$  | |______| "
hai_cham="             $$||      |__|                $$||      |__|                       "
def xuat_ky_tu(mang) :
    i=0
    for x in range(8) :
        for y in range(8) :
            b=mang[y]
            lines[x]+=b[i:i+10]+"  "
        i=i+10
    for z in range(8) :
        print(lines[z])
#Xuat_hour:minute:second
def print_ky_tu(a):
        for i in range(len(a)) :
            if a[i].isdigit() :
                mang_phu[i]=ky_tu[int(a[i])]
            else :
                mang_phu[i]=hai_cham
        xuat_ky_tu(mang_phu)
#ham_main
while True:
    mang_phu= [[],[],[],[],[],[],[],[],[]]
    lines=["","","","","","","","",""]
    time_now= datetime.datetime.now()
    t=str(time_now)[11:19]
    if time_now.second != pre_sec:
        os.system("clear")
        print(time_now)
        print(t)
        print_ky_tu(t)
        pre_sec=time_now.second
    time.sleep(0.001)
