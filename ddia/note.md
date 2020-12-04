# Designing Data-Intensive Application

## 可靠性，可伸缩性，可维护性

理想情况下，批量作业的运行时间是数据集的大小除以吞吐量。

中位数也被称为第 50 百分位点，有时缩写为 p50。95、99 和 99.9 百分位点（缩写为 p95，p99 和 p999）。

由于服务器只能并行处理少量的事务（如受其 CPU 核数的限制），所以只要有少量缓慢的请求就能阻碍后续请求的处理，这种效应有时被称为 头部阻塞（head-of-line blocking） 。

纵向伸缩（scaling up）（垂直伸缩（vertical scaling），转向更强大的机器）和横向伸缩（scaling out） （水平伸缩（horizontal scaling）
跨多台机器分配负载也称为“无共享（shared-nothing）”架构。
