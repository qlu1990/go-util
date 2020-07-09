# go-util
golang util-package


`Q: 为什么数据库中的btree(b+tree) 分支存储是左闭右开的`

A： 因为常用key值数据的递增性质(反过来也可以说明为啥数据库中索引建议使用递增数据)

`Q: btree(b+tree) 持久化如何实现`

A： 按照io page 的整数倍分block，block 内设置id
```
    树结构
              root
                |
            ------------
            |      |
            (1,a)  (5,d)
             |
          ------
          |
        (2,b)

   持久化数据结构

    block1  4K   root(id = 0) 
    block2  4k   1 -> a(id = 0),5-> d(id=1)
    block3  4k   2 -> b(id = 0)


    查找 key 2
    以root 入口 获得 查找获得 1->a 的（blockid，offsetid）
    在 1 -> a 中查找获得他的子叶（blockid，offsetid）
    最终获得 2 -> b


```
# 联系方式 
 mail: 876392131@qq.com
