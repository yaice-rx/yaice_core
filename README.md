###########YaIce
基于Golang开发的一款游戏框架
支持kcp协议

1、启动的时候需要根据[服务分组]和[服务类型]，决定所需要连接的服务     
2、逻辑处理通过channel，在同一线程中处理    
3、服务各个内部的连接都是单向，通过grpc连接处理

###########有问题反馈
在使用中有任何问题，欢迎反馈给我，可以用以下联系方式跟我交流

* 邮件(yaice.rx@qq.com)



###########感激
感谢以下的项目,排名不分先后

* [kcp](https://github.com/xtaci/kcp-go)
* [logrus](github.com/sirupsen/logrus)
* [etcd](github.com/coreos/etcd)
