[toc]

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

  然后在`Settings/Secrets/Actions`中`New repository secret` ,  其中docker信息是在[阿里云容器仓库](https://cr.console.aliyun.com/cn-hongkong/instances)中创建了一个个人实例(创建之后需要去`/个人实例/访问凭证`设置固定密码)，当然，你也可以使用其他的docker仓库，不过第2步的`Login to Aliyun Container Registry (ACR)`需要进行修改

  ### 2. 设置Secrets
<a id="Anchortable">表格</a>
   |  Name  |   Value   |   说明   |   举例   |
   | ---- | ---- | ---- | ---- |
   |   DOCKER_USERNAME   |  your docker username    |  docker仓库登陆用户名    | spatxos    |
   |   DOCKER_PASSWORD   |  your docker pwd    |  docker仓库固定密码    | spatxospwd    |
|   HOST  |  your server ip    |  服务器IP    | 101.10.11.121    |
|   HOST_USERNAME   |  your server username    |  服务器ssh登陆账户名    | spatxosdocker   |
|   HOST_PASSWORD   |  your server pwd    |  服务器ssh登陆密码    |  spatxosdockerpwd   |
|   HOST_PORT   |  your server ssh port    |  服务器ssh端口    |  22    |
|   DOCKER_REGISTRY   |  docker registry    |  docker仓库地址    |  registry.cn-hongkong.aliyuncs.com    |
|   DOCKER_REGISTRY_REGION   |  docker registry region id    |  docker仓库区域id    |  cn-hongkong    |
|   CNBLOGS_ISDOWN   |  Whether to pull blogs from cnblogs    |  本次执行是否从cnblogs拉取博客    |  true或false   |
|   CNBLOGS_COOKIE   |  cnblogs of cookie    |  cnblogs的cookie    |  __gads=ID=bbfxxxxxxxxxx    |
|   BLOG_NAME   |  blog of name    |  博客的名称，发布到docker或者服务器上创建的文件夹都将使用这个    |  spatxos    |
  ### 3.   在`/.github/workflows/`下添加一个yml文件，可以进行自定义
  我写好了一个[cicd.yml](https://github.com/spatxos/spatxos-blog/blob/master/.github/workflows/cicd.yml)，是将hexo发布到阿里云的docker仓库，然后进行服务器部署，后期将会添加发布到github pages的yml

# 另一种办法，直接fork
  现在我已经把仓库创建好并且上传到了github，仓库地址[https://github.com/spatxos/spatxos-blog](https://github.com/spatxos/spatxos-blog)，可以直接进行fork，然后去[阿里云容器仓库](https://cr.console.aliyun.com/cn-hongkong/instances)中创建一个个人实例，购买或者使用一个云服务器，在github仓库中填写一下<a id="#Anchortable">Secrets表格</a>中Secrets即可

# 从博客园拉取之间创建的博客到新建的hexo
  在Secrets表格中设置好`CNBLOGS_ISDOWN`是`true`，然后去登陆博客园，F12，随便找一个xhr类型的请求，查找对应的cookie，然后到`/source/_posts`下执行一下`go run convertcnblogbookie.go -cookie "替换成你的cookie"`将会获得一个输出的新cookie，将新的cookie作为`CNBLOGS_COOKIE`的value写入进去（github action中使用secrets时，不能包含某些特殊字符，否则会被截断，所以这次执行其实是对特殊字符的替换，后面拉取博客时再替换回来）
  在每次执行提交时，将会把docker里面的hexo静态页面映射到服务器的`/root/${{BLOG_NAME}}-blog/html`文件夹下，首次拉取博客园的文章之后，如何博客园文章未进行更新，我们可以更改一下`CNBLOGS_ISDOWN`为`false`，那么就不会从博客园拉取了，原本的文件还会存在，不会进行覆盖
# 遇到的问题
  ### 1.阿里云登陆和docker push 时tag错误，问题描述参见[docker tag error](https://github.com/actions/starter-workflows/issues/1635)，解决办法参见[GitHub Actions持续集成阿里云容器镜像服务（ACR）](https://mincong.io/cn/github-actions-acr/)，我后面这么写的
  `docker build -t "${{ secrets.DOCKER_REGISTRY }}/${{secrets.BLOG_NAME}}/${{secrets.BLOG_NAME}}-blog:${{env.TAG_NAME}}" -f Dockerfile .`
  ### 2.go传入参数和secrets截断问题
      go传入参数使用conf来做，secrets截断问题通过先替换掉会截断的字符，然后使用时替换回去

# [博客github地址](https://github.com/spatxos/spatxos-blog)