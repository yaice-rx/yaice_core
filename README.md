##YaIce是什么?
YaIce是基于Golang开发的一款游戏框架，在开发过程中采用kcp作为外部网络层，grpc作为内部连接，使用etcd服务。



##YaIce特征
* 动态增长服务
* 将各个模块分开，单独成立一个服务
* 登陆auth服务，采用nginx，利用负载均衡，动态分配auth，分配对应的game服务器

##有问题反馈
在使用中有任何问题，欢迎反馈给我，可以用以下联系方式跟我交流

* 邮件(yaice.rx@qq.com)



##感激
感谢以下的项目,排名不分先后

* [kcp](https://github.com/xtaci/kcp-go)
* [logrus](github.com/sirupsen/logrus)
* [etcd](github.com/coreos/etcd)

##关于作者

```javascript
  var ihubo = {
    nickName  : "YaIce",
    email : "yaice.rx@qq.com"
  }
```