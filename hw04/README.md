# Динамические массивы и очередь с приоритетом

## Цель

- Использование разных алгоритмов при реализации динамического массива и
  сравнение их производительности.
- Создание приоритетной очереди.

## Описание/Пошаговая инструкция выполнения домашнего задания

### УРОВЕНЬ JUNIOR

Выполнить все пункты. В каждом классе реализовать методы:

- `int size()` - вернуть размер
- `bool isEmpty()` - проверить массив на пустоту
- `private void resize()` - увеличить размер массива
- `T get(int index)` - получение элемента массива по индексу
- `void put(T item)` - размещение элемента в конец массива

---

- +1 байт. Реализовать динамический массив SingleArray.
- +1 байт. Реализовать динамический массив VectorArray.
- +1 байт. Реализовать динамический массив FactorArray.

Прислать отчёт о выполненном задании:

- какие пункты выполнены
- сколько байт собрано
- сколько времени ушло на выполнение домашнего задания
- ссылку на репозиторий

### УРОВЕНЬ MIDDLE

В каждом созданном классе дополнительно реализовать методы:

- `void put(T item, int index)` - вставка в указанный индекс со смещением вправо
- `T del(int index)` - удаление указанного элемента со смещением влево, вернуть
  удалённый элемент

---

- +1 байт. Реализовать эти методы в классах SingleArray, VectorArray,
  FactorArray

Создать динамический массив MatrixArray, использовать двойную адресацию:

- +1 байт. Реализовать в MatrixArray методы из уровня JUNIOR
- +1 байт. Реализовать в MatrixArray метода из уровня MIDDLE

Прислать отчёт о выполненном задании:

- какие пункты выполнены
- сколько байт собрано
- сколько времени ушло на выполнение домашнего задания
- ссылку на репозиторий

### УРОВЕНЬ SENIOR

Составить сравнительную таблицу по времени работы в созданных динамических
массивах на 10^3, 10^4, 10^5, 10^6, 10^7 элементов.

- +1 байт. Сделать обёртку над стандартным `ArrayList()` со всеми перечисленными
  методами для сравнительного тестирования
- +1 байт. Для метода `void put(T item)` для каждого класса
- +1 байт. Для метода `void T del(0)` для каждого класса, запускать и засекать
  время после того, как будут предварительно добавлены все элементы
- +1 байт. Для метода `void put(T item, 0)` и
  `void put(T item, random(0, size()-1))` для каждого класса

Прислать отчёт о выполненном задании:

- Приложить скриншот сравнительной таблицы
- Сделать выводы и сформулировать их в нескольких предложениях
- Какие пункты выполнены, сколько байт собрано
- Сколько времени ушло на выполнение домашнего задания
- Ссылку на сравнительную таблицу и текст вывода о сравнении результатов

### УРОВЕНЬ ARCHITECTOR

- +2 байта. Написать реализацию PriorityQueue - очередь с приоритетом

Варианты реализации - список списков, массив списков

Методы к реализации:

- `void enqueue(int priority, T item)` - поместить элемент в очередь
- `T dequeue()` - выбрать элемент из очереди

Прислать отчёт о выполненном задании:

- Какие пункты выполнены
- Сколько байт собрано
- Сколько времени ушло на выполнение домашнего задания
- Ссылку на сравнительную таблицу и текст вывода о сравнении результатов

При реализации следует использовать только стандартные массивы []. Стандартные
коллекции не использовать!
