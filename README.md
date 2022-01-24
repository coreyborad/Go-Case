## Description

...Todo

## Table of Contents

- [Description](#description)
- [Table of Contents](#table-of-contents)
- [Case](#case)
    - [Simple Routine](#simple-routine)
    - [Channel](#channel)
    - [Channel Queue](#channel-queue)
    - [Worker Pool](#worker-pool)
    - [Graceful shutdown](#graceful-shutdown)

## Case

### Simple Routine

go routine最基本的用法，此案例會一次生成20組go routine，搭配waitGroup，等待所有routine做完

### Channel

透過channel的read跟write特性，去通知另外一個go routine跟channel去終止整個程式的運作

### Channel Queue

Channel是具有queue的特性，可以設定buff跟unbuff，達到上限時就無法再往裡面推進

### Worker Pool

當大量的job進來時，可能會開啟過多的routine，如果不想要花費過多的資源，就必須去限制routine數量，這邊稱worker。

當有新的job進來時，透過channel傳進去worker去做事情，worker都被佔用時，因為channel有queue的特性，因此需排隊等待worker被釋放，類似加油站排隊加油的案例。

ex 加油站有4個加油槽，這時候有20台車需要加油
1. 加油槽=worker數, 車子=job, 車道=job chan, 加油站站長辦公室報表機=result chan
2. 每台車子加油的花費時間都不同，加完油離開，加油槽就會空掉，新的車子再補上
3. 每台車都需要排隊等待加油，無法插隊

### Graceful Shutdown

讓服務關機之前，需等待剩餘的job做完後，才可以關閉。

原理相當於去攔截os層signal要關閉的訊號，使其先暫停關閉，設定多等待的時間，讓剩餘job做完，再繼續讓signal往下跑