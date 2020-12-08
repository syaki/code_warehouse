# Designing Data-Intensive Application

## 可靠性，可伸缩性，可维护性

理想情况下，批量作业的运行时间是数据集的大小除以吞吐量。

中位数也被称为第 50 百分位点，有时缩写为 p50。95、99 和 99.9 百分位点（缩写为 p95，p99 和 p999）。

由于服务器只能并行处理少量的事务（如受其 CPU 核数的限制），所以只要有少量缓慢的请求就能阻碍后续请求的处理，这种效应有时被称为 头部阻塞（head-of-line blocking） 。

纵向伸缩（scaling up）（垂直伸缩（vertical scaling），转向更强大的机器）和横向伸缩（scaling out） （水平伸缩（horizontal scaling）
跨多台机器分配负载也称为“无共享（shared-nothing）”架构。

## 数据模型与查询语言

### 关系模型与文档模型

现在最著名的数据模型可能是 `SQL` 。它基于 `Edgar Codd` 在 1970 年提出的关系模型【1】：数据被组织成关系（ `SQL` 中称作表），其中每个关系是元组（ `SQL` 中称作行)的无序集合。

### `NoSQL`

不仅是 SQL（Not Only SQL）

混合持久化 `polyglot persistence`

模型之间的不连贯有时被称为阻抗不匹配（impedance mismatch）

对象关系映射（ORM object-relational mapping） 框架可以减少这个转换层所需的样板代码的数量，但是它们不能完全隐藏这两个模型之间的差异。

### 图数据模型

一个图由两种对象组成：顶点（`vertices`）（也称为节点（`nodes`） 或实体（`entities`）），和边（`edges`）（ 也称为关系（`relationships`）或弧 （`arcs`） ）。多种数据可以被建模为一个图形。

在属性图模型中，每个顶点（vertex） 包括：

唯一的标识符

-   一组 出边（outgoing edges）
-   一组 入边（ingoing edges）
-   一组属性（键值对）
-   每条 边（edge） 包括：

唯一标识符

-   边的起点/尾部顶点（tail vertex）
-   边的终点/头部顶点（head vertex）
-   描述两个顶点之间关系类型的标签
-   一组属性（键值对）

## 存储与检索
