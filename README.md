# accumulator pipline worker

# Technical description EN
Need to implement the data processing model with a pipeline behavior that consists from these parameters:
1. Give to the input data packages. Data package = slice with len = 10 that consists of random integer elements. 
New packages should be created and sent to the pipeline every N ms (N sets as an environment variable).
2. Processing packages: finding three biggest figures in the package. Input: slice with len = 10, output slice with len = 3.
Data processing must use workers=M (M sets as an environment variable).
3. Accumulator: summarize three biggest figures received from the previous step and saving to the separate variable int (accumulator).
4. Publisher: publish to console current value of the accumulator every K seconds (K sets as an environment variable)

Example:
input: {1, 9, 6, 4, 4, 5, 7, 8, 0, 1}
processing: {9, 7, 8}
accumulator: 9+7+8=24
publisher: 24

# Technical description RU

Реализовать модель обработки данных в виде пайплайна, состоящего из следующих этапов
1. Подача на вход пакетов данных. Пакет данных = слайс случайных целых чисел из 10 элементов. 
Новый пакет подается каждые N мс (N задается в виде env переменной)   
2. Обработка пакетов: нахождение 3-х наибольших чисел в пакете. Вход: слайс int из 10 элементов, выход: слайс из 3-х элементов. 
Обработка пакетов должна производиться M воркерами (M задается в виде env переменной)
3. Аккумулятор: суммирование чисел обработанных пакетов, полученных на предыдущем этапе, и запись в единую переменную int
4. Публикатор: вывод на консоль текущего значения аккумулятора каждые K секунд (K задается в виде env переменной)

Пример:
вход: {1, 9, 6, 4, 4, 5, 7, 8, 0, 1}
обработка: {9, 7, 8}
аккумулятор: 9+7+8=24
публикатор: 24