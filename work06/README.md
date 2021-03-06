#### 什么是粘包
发送两条消息：消息1：abc,消息2：def

正常情况接受者收到：消息1：abc，消息2：def

![alt 属性文本](./images/socket1.png)

粘包：接受者收到消息1：abcdef

![alt 属性文本](./images/socket2.png)

半包：接受者收到消息1：ab，消息2：cdef

![alt 属性文本](./images/socket3.png)

半包：接受者收到消息1：abcd，消息2：ef

![alt 属性文本](./images/socket4.png)

#### 为什么会发生粘包
tcp协议为了高效传输付出的代价；对tcp来说，它处理的是底层的数据流，数据流本身没有任何开始和结束的边界。

发送数据过程：应用程序发送消息包，消息包以数据流的形式放入缓冲区，等待缓冲区的数据流达到一定阀值后，在发送到网络上。

接收数据过程：接收到网络传输的数据流，放入缓冲区，缓冲区的数据流达到一定的阀值后，通知应用程序进行读取数据。

在接受和发送数据的过程中，都是对数据流进行操作。

发送数据的时候，发送的数据长度大于缓冲区空间，就会发生数据流拆分，同一个数据包就会通过多次发送完成，表现为半包的情况。

发送数据的时候，发送的数据长度小于缓冲区空间，多个数据包填满了缓冲区才会进行发送，表现为粘包的情况。

接收数据的时候，应用程序没有持续读取缓冲区内容，导致缓冲区存放了多个数据包，在进行读取，也会表现粘包的情况。

#### 如何处理粘包
1.fix length

发送方每次发送固定长度的数据，且不超过缓存区大小；接收方每次按固定的长度接收数据。

2.delimiter based

发送方在发送数据中添加特殊分隔符，用来标记数据包边界；接收发按分隔符区别包数据。

3.length field based

封装请求协议；数据头+数据正文，在数据头中存储数据正文的大小，当读取的数据小于数据头中的大小时，继续读取数据，直到读取的数据长度等于数据头中的长度时才停止。