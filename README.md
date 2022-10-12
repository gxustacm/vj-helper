### Usage
1. download config.yaml and virtual_judge.exe/virtual_judge.darwin/virtual_judge.linux in same dictionary
2. fill in config.yaml, the parameters are listed below:
	* username: username of virtual judge
	* password: password of virtual judge
	* groupId: There is a pre-release edition, so we only support group mode, you need create a group in virtual judge, and find the groupId, which can be found by follow the step below:
		- Change group logo, confirm logo of group is not default logo. The address of default logo comes from gravatar, that's not what we need.
		- copy the logo's address of your group, then paste it in any place, like "https://vj.csgrandeur.cn/group/logo/8892?v=1606461042", the 8892 is your groupId.
	* title: the title of contest
	* announcement: as you think
	* problem: It's a problem array, the low to high is the problem rating scope, tag is optional.
3. run the program(It's best to run the program in terminal instead of opening it directly. If you open it directly, you may miss some important massage,such as error massage or other).

### warning
1. focus on your network, don't attempt to use this software before a contest.
2. you can change problemset: https://codeforces.com/api/problemset.problems to problemset: https://codeforces.ml/api/problemset.problems, this is a mirror.

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