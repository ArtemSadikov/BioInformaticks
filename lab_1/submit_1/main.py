# put your python code here
import sys

data = sys.stdin.readlines()    #чтение строки
i,a = 0,0                       #счетчики

temp1 = data[0]                 #строка первая (слово, которое ищем)
temp2 = data[1]                 #строка вторая (текст)

temp1 = temp1.split('\n')       #разбиваем на элементы
temp1 = temp1[0]                

temp1_len = len(temp1)          #запоминаем длинны строк
temp2_len = len(temp2)

while i < temp2_len:            #цикл, проходим через строку в поиске нашего слова и запоминаем количество совпадений
    if (temp2[i:i+temp1_len] == temp1):
        a = a+1
    i = i+1
print(a)                        #вывод