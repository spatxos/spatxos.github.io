title: 使用hexo admin写博客
author: spatxos
date: 2022-07-07 12:17:04
swiper: true # 将改文章放入轮播图中
categories: ['前端']
tags: ['hexo', 'hexo-admin']
---
## 安装hexo admin
1. 安装hexo并初始化一个博客
```
npm install -g hexo
cd ~/
hexo init my-blog
cd my-blog
npm install 
```
2. 安装hexo admin并运行
```
npm install --save hexo-admin
hexo server -d
open http://localhost:4000/admin/
```
3.实际上安装之后的hexo admin在图片粘贴进去时，图片的路径出现了问题，我们找到`\node_modules\hexo-admin\api.js`，进行替换
```
var imagePath = '/images' 替换成 var imagePath = 'images'
在
hexo.source.process().then(function () { 
这一行前添加一行
var imageSrc = path.join(hexo.config.root + filename).replace(/\\/g, '/').replace('.png\/', '.png')
然后将
res.done({
          src: path.join(hexo.config.root + filename),
          msg: msg
})
替换成
res.done({
          src: imageSrc,
          msg: msg
})
```
这样发不出来的图片就不会出现裂图了