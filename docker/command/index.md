# docker基本命令

## run命令

## 删除悬空的镜像

在 Docker 中，<none> 镜像通常指的是没有标签的镜像或“悬空”（dangling）镜像。这些镜像往往是由于构建过程中的中间层或删除的标签导致的。删除这些镜像可以释放存储空间。

以下是删除所有 <none> 镜像的几种方法：

### 1. 使用 Docker CLI 删除 <none> 镜像

可以使用以下命令查找并删除所有 <none> 镜像：

docker image prune


此命令将删除所有悬空（dangling）镜像，悬空镜像是指没有与任何标签关联的镜像。

如果您想要更深入的控制，可以使用以下命令：

docker rmi $(docker images -f "dangling=true" -q)


- docker images -f "dangling=true"：过滤出所有悬空镜像。
- -q：仅显示镜像 ID。
- $(...)：将命令的输出作为参数传递给 docker rmi，从而删除这些镜像。

### 2. 删除所有无标签的镜像

如果您想删除所有无标签的镜像（即 <none> 镜像），包括不再使用的及中间层，可以使用以下命令：

docker rmi $(docker images | grep '<none>' | awk '{print $3}')


- 这条命令会搜寻所有没有标签的镜像并删除它们。
- awk '{print $3}' 用于提取到 <none> 镜像的镜像 ID。

### 3. 直接删除所有未使用的镜像

如果您想删除所有未使用的镜像，包括 <none> 镜像，可以使用：

docker system prune


这会删除所有未使用的容器、网络、悬空的镜像和构建缓存。运行这个命令时，Docker 可能会要求确认操作。

> 注意：在运行此命令之前，确保您了解所有未使用的项，因为这会删除 Docker 系统中所有不再使用的容器和镜像。

### 4. 强制删除悬空镜像

如果某些镜像无法删除（例如它们可能被正在运行的容器所占用），您可能需要强制删除这些镜像。可以使用：

docker rmi -f $(docker images -f "dangling=true" -q)


- -f 选项强制删除镜像。

### 总结

以上几种方法可以帮助您清理 Docker 中的 <none> 镜像。根据您的需求选择合适的方法进行清理。如果您只是想删除悬空镜像，使用 docker image prune 会是最安全和高效的选择。
