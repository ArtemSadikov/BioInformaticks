# put your python code here
import sys                              #подключаем библиотеки
from collections import Counter

data = sys.stdin.readlines()            #считываем строку

temp = data[0]                          #записываем первую строку
k = int(data[1])                        #записываем вторую строку (число)

temp = temp.split('\n')                 #разделаем строку на элементы
temp = temp[0]

temp_len = len(temp)                    #запоминаем длинну

i = 0                                   #создаем переменную цикла и новый массив(список)
words = []

while i < temp_len:                     #проходим по строке и записываем в новый список слова длинной k
    if(len(temp[i:i+k]) == k):
        words.append(temp[i:i+k])
    i = i+1

c = Counter(words)                      #сортируем словарь, для нахождения количество повторений

max_val = max(c.values())               #запоминаем max кол.повторений

out = []                                #создаем список вывода, записываем в него слова с значением max
for k,v in c.items():
    if v==max_val:
        out.append(k)
j=0
while j < len(out):                     #вывод каждого слова
    print(out[j])
    j = j+1