## Description

...Todo

## Table of Contents

- [Description](#description)
- [Table of Contents](#table-of-contents)
- [Case](#case)
    - [Simple Routine](#simple-routine)
    - [Channel](#channel)
    - [Channel Queue](#channel-queue)
    - [Channel 3](#channel-3)
    - [Channel 4](#channel-4)
    - [Channel 5](#channel-5)

## Case

### Simple Routine

go routine最基本的用法，此案例會一次生成20組go routine，搭配waitGroup，等待所有routine做完

### Channel

透過channel的read跟write特性，去通知另外一個go routine跟channel去終止整個程式的運作

### Channel Queue

Channel是具有queue的特性，可以設定buff跟unbuff，達到上限時就無法再往裡面推進

