# CypressAM
用于管理员对软件后台数据的可视化交互管理，允许多种角色用户自由设计菜单功能和授予权限的后台数据管理系统

###部署&使用
搭配完备的前端应用使用更佳

或者直接用于后端接口测试

1、修改config文件中连接本地数据库的配置信息

![image](https://github.com/user-attachments/assets/3364f6ee-1e0d-45b9-a987-c8d1a7166387)

2、gorm会根据model里的表结构自动在本地数据生成对应的表（但信息需要自行添加）

![image](https://github.com/user-attachments/assets/be2e5a6a-0c27-4039-a959-1bc1738cb783)

3、配置代码文件相关的modules（使用编译器设置中的go代理安装更快）

![image](https://github.com/user-attachments/assets/a9df011e-92e7-4ca8-acbf-80431ab81fea)

3、运行main.go文件

测试方法可以是简单的浏览器URL的GET/POST请求
但是这样或许难以提交表单
也可以借助api测试工具（如postman或者apifox），可以看见接口反馈的详细信息

![image](https://github.com/user-attachments/assets/cbee4998-d832-47e9-b786-1ed22a75662a)


###添加你需要的菜单功能接口
1、根据你的数据表结构使用gorm构建你的model（程序运行时会自动在数据库中添加新的表，当然手动在数据库中添加也是可行的）；


2、在service中添加你的代码，通过gorm的api与数据库产生联系；

![image](https://github.com/user-attachments/assets/d984b109-04a3-4f6d-ad84-28dc6423aa6e)


3、api文件夹中是你获取前端返回信息进行处理的接口，通过它调用到刚才你的service代码；

![image](https://github.com/user-attachments/assets/2eb54516-c991-42cf-b9d1-a95786b4f140)


4、在api文件夹下，为你的新的接口添加路由组，不仅能让你方便统一管理这些接口，在后面步骤的注册路由中也能为你省去许多行代码；

![image](https://github.com/user-attachments/assets/57e05cb0-5067-41dd-b411-5c802e5d069a)


5、在initilization中初始化你的路由组，记住路由组分为公开和私有，根据你的情况决定它，错误的作用域会导致你的接口无法被访问！；

![image](https://github.com/user-attachments/assets/a77a0ef6-8b83-4d60-a028-ebbcbb6adb7b)

添加新的菜单接口主要步骤如上所述，如果想优化你的接口逻辑，可以仿照其他现成的接口代码看看，他们大部分是很不错的例子
