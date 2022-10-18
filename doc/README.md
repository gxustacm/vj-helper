## 介绍
这是一个基于codeforces题目和virtual judge的项目, 它可以帮你解放双手, 随机从codeforces中选取符合你要求的题目,然后在virtual judge中生成比赛.

### 使用方法
1. 下载config.yaml文件, 并且根据你的操作系统选择适合的压缩包下载, 你可以在release中下载, 我推荐你自行编译软件, 使用go build进行编译即可.
2. 填写config.yaml文件, 这里有一些参数列举在下方
    * username: virtual judge的账号
    * password: virtual judge的密码
    * groupId: 由于这是一个尝鲜版本, 我们只支持group模式创建比赛, 你可以先创建一个群组在virtual judge中, 并且找到组群的id, 这是确定组群唯一性的唯一方法
        - 首先上传一张组群头像, 确保它不是默认头像, 然后查看组群头像的url, 你会得到一个像这样的url, "https://vj.csgrandeur.cn/group/logo/8892?v=1606461042", 那么"8892"就是你的组群id
    * title: 比赛的标题
    * length: 比赛的时间长度 "5h12m"的格式
    * beginTime: 开始时间 2022-10-12 23:12:34
    * announcement: 会写在比赛的最上面哦, 支持markdown
    * problem: 填写一个题目的数组, 每道题目随机的最高难度和题目的最低难度是必须填写的, 每个题目的tag是可选项(不填写的时候随机)
3. 运行程序, 确保程序和config.yaml在一个文件夹中(最好在terminal中运行程序, 而不是直接运行, 否则你可能错过一些报错信息)

### 注意
1. 注意你的网络环境, 该软件的速度基于你连接codeforces的速度

## config.yaml format
you have to create config.yaml before run program.
```
problemset: https://codeforces.com/api/problemset.problems //codeforces api domain

username: $vj_account
password: $vj_password

groupId: 9109  //vj groupId, you can find out by above method in usage
title: etst // contest title
announcement: miaomiao
beginTime: 2022-10-12 23:12:34 
length: 5h12m //contest length

// you can add multiple problem below.
problems:
-   low: 800 // lowest rating of problem
    high: 1600 // highest rating of problem 
-   low: 1600
    high: 2000
-   low: 2000
    high: 2400
```