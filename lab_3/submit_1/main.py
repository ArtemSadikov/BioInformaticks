# put your python code here
import sys                      #библиотеки

data = sys.stdin.readlines()    #считываем строки

temp1 = data[0]                 #записываем первую строку
temp2 = []                      #новый список для будущей строки

i=0

while i < len(temp1):           #цикл, в котором заменяем элементы на обратные
    if(temp1[i] == 'A'):        # и добавляем их в новый список
        temp2.append('T')
    elif(temp1[i] == 'T'):
        temp2.append('A')
    elif(temp1[i] == 'C'):
        temp2.append('G')
    elif(temp1[i] == 'G'):
        temp2.append('C')
    i = i+1
    
temp2 = ''.join(temp2)          #убираем пространство между элементами (запись в одно слово)
temp2 = temp2[::-1]             #отзеркаливаем слово

print(temp2)