--[[Многострочный
комментарий]]

-- одинарный комментарий

function summ(a, b)
    return a + b
end

function rate (a, b)
    return a ^ b
end

print(" --- file.lua ---")
print("summ =", summ(20, 40))
print("rate =", rate(2, 10))

local st1 = 'text' --присвоение строки
st2 = st1           --приравнивание другой переменной
st3 = st1 .. st2    --результат сложения двух строк
print("st3 = ", st3)

i = 10     --number
st = "11"  --string
konc1 = 1 .. st -- конкатенация
konc2 = 100 .. ""  --сцепление с пустой строкой переводит переменную в string
print(konc1, type(konc1), konc2, type(konc2))  -->>111  100  string string

i1 = 1
str = "11"
i2 = tonumber("555") --перевод string -> number
s2 = tostring(i1) -- перевод number -> string
print(i1, str, i2, type(i2), s2, type(s2))

local Text = "Test Text Example string"
lenText = string.len(Text)
print(string.format("%s *** lenText = %d", Text, lenText))

print(string.lower('Hello World!')) -->> hello world!
print(string.upper('Hello World!')) -->>HELLO WORLD!
local s = 'Corona '
print(s:rep(3))-->> Corona Corona Corona
print(string.reverse('9876543210')) -->0123456789

local st = 'To be or not to be'
print(st:sub(7))-->>or not to be
print(st:sub(7,12))-->>or not
print(st:sub(-5))-->>to be
print(st:sub(4,-3))-->>be or not to
print(st:sub(-8,-4))-->>ot to

local s2 = 'dust ist fantastisch'
print(s2:gsub('s','S'))-->>duSt iSt fantaStiSch	4

local s3 = 'lalalalala'
print(s3:gsub('a','A',3))-->>lAlAlAlala	3

--[[
https://smart-lab.ru/blog/291665.php

a = tostring (10) — a равно «10»
a = tostring (true) — a равно «true»
a = tostring (nil) — a равно «nil»
a = tostring ({[1] = «это поле 1»}) — a равно «table: 06DB1058»

a = render (10) — a равно «10»
a = render (true) — a равно «true»
a = render (nil) — a равно «nil»
a = render ({[1] = «это поле 1»}) — a равно "{[1] = «это поле 1»}"

a = tonumber («10») — a равно «10»
a = tonumber («10»..".5") — a равно 10.5
a = tonumber (true) — a равно «nil»
a = tonumber (nil) — a равно «nil»

a = «строка»
len = #a — len равно 6
len = #«ещё строка» — len равно 10

if a < 0 then
   a = 0 — если a меньше 0, присвоить a значение 0
else
   a = 1
end

i = 10; t = {}
while i > 0 do    -- цикл от 10 до 1
   t[i] = «поле » .. i  -- .. склейка строк
   i = i — 1
end

a = {3, 2, 5, 7, 9}
i = 0; sum = 0
repeat
   i = i + 1
   sum = sum + a[i]
until sum > 10


for i = 1, 10 do  -— цикл от 1 до 10 с шагом 1
  MsgBox («i равно » .. i)
end

for i = 10, 1, -1 do   —- цикл от 10 до 1 с шагом -1
   MsgBox («i равно » .. i)
end

a = {3, 5, 8, -6, 5}   -- #a = 5 - 'это длина массива'
for i = 1,#a do        -— ищем в массиве отрицательное значение
   if a[i] < 0 then   -— если найдено...
      index = i       —- сохраняем индекс найденного значения...
      break           -— и прерываем цикл
   end
end

https://habr.com/ru/post/344562

--считаем от 2 до 6 прибавляя по 2
for i = 2,6,2 do
	print(i)
end

--считаем от 9 до 3 отнимая по 3
for i = 9,3,-3 do
	print(i)
end

local reg = " (%w+):(%d+)"--регулярное выражение с захватом 2 параметров
local s = "masha:30, lena:25, olia:26, kira:29"  --строка для поиска и захвата

for key, value in s:gmatch(reg) do   --цикл с захватом всех вхождений
	print(key, value)
end
-->> masha  30
-->> lena  25
-->> olia  26
-->> kira  29

for key, value in ipairs{'vasia','petia','kolia'} do
	print(key, value)
end
-->> 1  vasia
-->> 2  petia
-->> 3  kolia

for key, value in pairs{v1 = 'OK',v2 = 123, v3 = 12.3} do
	print(key, value)
end
-->> v1  OK
-->> v2  123
-->> v3  12.3

Работа с файлами
Создание/Запись/Чтение файла целиком

local write_data = "Hello file!"  --строка для записи в файл
local fileName = 'saves.dat'      --имя файла
local filePath = system.pathForFile(fileName, system.DocumentsDirectory )--путь к файлу
local file = io.open(filePath, "w")    --открываем файл с его созданием для записи
file:write( write_data )   --записываем
io.close( file )           --закрываем файл

    «r» — режим чтения (по умолчанию); указатель файла помещается в начале файла
    «w» — режим только для записи; перезаписывает файл, если файл существует. Если файл не существует, создается новый файл для записи.
    «a» — режим добавления (только для записи); указатель файла находится в конце файла, если файл существует (файл находится в режиме добавления). Если файл не существует, он создает новый файл для записи.
    «r+» — режим обновления (чтение / запись); все предыдущие данные сохранены. Указатель файла будет в начале файла. Если файл существует, он будет перезаписан, только если вы явно напишите ему.
    «w+» — режим обновления (чтение / запись); все предыдущие данные стираются. Перезаписывает существующий файл, если файл существует. Если файл не существует, создается новый файл для чтения и записи.
    «a+» — добавить режим обновления (чтение / запись); предыдущие данные сохраняются, и запись разрешается только в конце файла. Указатель файла находится в конце файла, если файл существует (файл открывается в режиме добавления). Если файл не существует, он создает новый файл для чтения и записи.

local fileName = 'log.txt'  --имя файла
local filePath = system.pathForFile(fileName, system.CachesDirectory )--путь к файлу
local file = io.open( filePath, "r" )  --открываем файл для чтения
for line in file:lines() do            --в цикле читаем все строки
	print( line )  --печать строки в консоль
end
io.close( file )   --закрываем файл

]]