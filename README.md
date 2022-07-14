# 搭建Hexo博客
  环境
  ```
  > node -v
  v16.15.0
  > npm -v
  8.5.5
  ```
  安装Hexo CLI
  ```
  # 在你的博客文件夹根目录下下执行
  npm install -g hexo-cli
  hexo init blogname
  cd blogname
  hexo s
  # 打开http://localhost:4000 即可浏览
  ```
# 安装主题hexo-theme-bamboo
  [hexo-theme-bamboo](https://github.com/yuang01/hexo-theme-bamboo)
  我们不使用npm安装，直接将主题download到themes文件夹下
  github安装
  ```
  git clone https://github.com/yuang01/hexo-theme-bamboo.git
  ```
  gitee安装
  ```
  git clone https://gitee.com/yuang01/hexo-theme-bamboo.git
  ```
  修改hexo根目录下的站点配置文件_config.yml，把主题改为hexo-theme-bamboo，通过主题文件夹下的config.yml配置主题即可，然后在`\themes\hexo-theme-bamboo\`删除`.git`文件夹
  然后根据[https://yuang01.github.io/](https://yuang01.github.io/)或者[作者博客地址](http://120.48.121.186/)来配置对应的样式或者widget
# 修改博客名称等信息
  在`_config.yml`和`\themes\hexo-theme-bamboo\_config.yml`中将自己博客名称、网址信息等进行替换
# 添加github action发布
  ### 1. 在github中创建自己的博客仓库

  然后在`Settings/Secrets/Actions`中`New repository secret` ,  其中docker信息是在[阿里云容器仓库](https://cr.console.aliyun.com/cn-hongkong/instances)中创建了一个个人实例，当然，你也可以使用其他的docker仓库，不过第2步的`Login to Aliyun Container Registry (ACR)`需要进行修改
<a id="Anchortable">表格</a>
   |  Name  |   Value   |   说明   |
   | ---- | ---- | ---- |
   |   DOCKER_USERNAME   |  your docker username    |  docker仓库登陆用户名    |
   |   DOCKER_PASSWORD   |  your docker pwd    |  docker仓库密码    |
|   HOST  |  your server ip    |  服务器IP    |
|   HOST_USERNAME   |  your server username    |  服务器ssh登陆账户名    |
|   HOST_PASSWORD   |  your server pwd    |  服务器ssh登陆密码    |
|   HOST_PORT   |  your server ssh port    |  服务器ssh端口    |

  ### 2.   在`/.github/workflows/`下添加一个yml文件，可以进行自定义
  我写好了一个yml，是将hexo发布到阿里云的docker仓库，然后进行服务器部署，后期将会添加发布到github pages的yml
  ```
  name: Build Docker Image

on:
  push:
    branches:
      - main
      - master
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout source
        uses: actions/checkout@v2
        with:
          ref: master

      - name: Setup Node.js
        uses: actions/setup-node@v1
        with:
          node-version: '16.15.0'

      - name: Setup Hexo
        run: |
          npm install hexo-cli -g
          npm install hexo-generator-search --save
          npm install hexo-generator-feed --save
          npm i hexo-wordcount
          npm install
      - name: Login to Aliyun Container Registry (ACR)
        uses: aliyun/acr-login@v1
        with:
          login-server: registry.cn-hongkong.aliyuncs.com
          region-id: cn-hongkong  # 3
          username: "${{ secrets.DOCKER_USERNAME }}"
          password: "${{ secrets.DOCKER_PASSWORD }}"
      - name: Deploy and Build Image
        run: |
          hexo clean
          hexo deploy
          ls -la
          pwd
          docker build -t spatxos/spatxos-blog:latest -f Dockerfile .
      - name: Push Image
        run: |
          docker pull registry.cn-hongkong.aliyuncs.com/spatxos/spatxos-blog:latest

  # Docker 自动部署
  deploy-docker: 
    needs: [build]
    name: Deploy Docker
    runs-on: ubuntu-latest
    steps:
      - name: Deploy
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.HOST }} # 服务器ip
          username: ${{ secrets.HOST_USERNAME }} # 服务器登录用户名
          password: ${{ secrets.HOST_PASSWORD }} # 服务器登录密码
          port: ${{ secrets.HOST_PORT }} # 服务器ssh端口
          script: |
            docker info
            echo $(docker ps -aqf "name=spatxos-blog")
            docker stop $(docker ps -aqf "name=spatxos-blog")
            docker container rm spatxos-blog
            docker rmi spatxos-blog
            echo 查看是否成功删除spatxos-blog
            docker ps -a
            echo 从harbor拉取docker镜像
            docker pull registry.cn-hongkong.aliyuncs.com/spatxos/spatxos-blog:latest
            docker tag registry.cn-hongkong.aliyuncs.com/spatxos/spatxos-blog:latest wangpengzong/spatxos-blog:latest
            docker run -it --rm -d -p 80:80 -v /root/soft/docker:/root/soft/docker --name spatxos-blog wangpengzong/spatxos-blog
            docker system prune -f
            echo docker容器启动成功
  ```
# 另一种办法，直接fork
  现在我已经把仓库创建好并且上传到了github，仓库地址[https://github.com/spatxos/spatxos-blog](https://github.com/spatxos/spatxos-blog)，可以直接进行fork，然后去[阿里云容器仓库](https://cr.console.aliyun.com/cn-hongkong/instances)中创建一个个人实例，购买或者使用一个云服务器，在github仓库中填写一下<a id="#Anchortable">Secrets表格</a>中Secrets即可