### 程序设计
整个程序主要由两大块组成，首先是获取输入的指令，调用getputin函数来实现，同时在这个函数中提取出输入的指令，放入一个结构体中储存。然后调用一个执行函数，将输入的指令进行分析执行。整个代码的输入规则都是由老师提供的链接文档中固定的。
代码的测试情况如下：

1.	将读入文档的第一页输出到屏幕上
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927125030945.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

2.	读取标准输入，而标准输入已被重定向为来自“text.txt”而不是显式命名的文件名参数
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927125302205.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

3.	将第1页至第二页写至标准输出，标准输出被重定向至“out.txt”
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019092712553663.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927125544279.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

4.	将1到2页写至标准输出，并将所有错误信息输出到error.txt文件中
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927140955835.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927141002329.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)
（由于没有错误信息，所以此文件为空）

5.	将第1至2页写至标准输出，并被重定向至out.txt文件中，左右错误信息被重定向至error.txt
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927141521426.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927141527737.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927141458142.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

6.	将1至2页写至标准输出，被重定向至out.txt，所有错误信息被重定向至空设备
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927141804610.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927141837649.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

7.	将标准输出丢弃，错误信息显示
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927142027262.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927142033326.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

8.	将页长设置为10行，并写入out.txt

![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927154325475.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927154339840.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

9.	由换页符定界
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927154619244.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)

10.	将第1页发送至命令cat，将前10行打印在屏幕上
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190927154940115.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NvZGlmZmVyZW50,size_16,color_FFFFFF,t_70)
