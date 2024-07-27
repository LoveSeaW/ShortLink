# ShortLink
Short link project based on golang

## 搭建项目骨架
1. 建库见表
新建发号器表
新建长链接短链接映射表
2. 搭建go-zero框架的骨架
编写go-zero api文件，使用框架指令生成代码
```bash
goctl api go -api shortener.api -dir .

goctl model mysql datasource -url="root:20212021Ynn.@tcp(82.156.206.148:3306)/shortLink" -table="short_url_map" -dir="./model"

goctl model mysql datasource -url="root:20212021Ynn.@tcp(82.156.206.148:3306)/shortLink" -table="sequence" -dir="./model"

go mod tidy
```