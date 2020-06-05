# mps-common
mps-common

# protogenerator.sh使用方法

## 1. bash protogenerator.sh

不带参执行，表示对当前所有proto进行代码生成，当前proto列表为：

agv

area

cameramonitor

config

devices

ledscreen

log

lpr

material

opc

order

point

rack

racklot

racktype

role

route

script

segment

station

system

task

user

## 2. bash protogenerator.sh xxx

带参执行，表示对指定的proto进行代码生成.

比如根据agv.proto生成代码

bash protogenerator.sh agv
