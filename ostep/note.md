# OSTEP

## 进程调度
### 调度指标
turnaround time 周转时间 = 完成时间 - 到达时间<br />response time 响应时间 = 首次运行 - 到达时间
### 调度算法
**FIFO**<br />**SJF 最短任务优先**<br />**STCF 最短完成时间优先**<br />**Round-Robin RR 轮转**<br />**MLFQ 多级反馈队列**<br />game the scheduler 愚弄调度程序 产生 starvation 饥饿问题 通过 boost 周期性提升 解决。<br />周期时间：巫毒常量 voo-doo constant。避免巫毒常量 Ousterhout 定律。<br />FreeBSD 调度程序，基于当前进程使用了多少CPU，通过公式计算某个工作的当前优先级。Epema论文 使用量衰减算法 decay-usage。<br />规则：

| A 优先级 > B，运行A |
| --- |
| A 优先级 = B，轮转运行 A & B<br /> |
| 工作进入系统时，优先级最高 |
| 一旦工作用完了一层的时间配额（无论中间主动放弃多少次CPU），就降低优先级 |
| 经过一段时间，所有工作重新加入最高优先级队列 |

#### proportional-share 比例份额调度 || fair-share 公平份额调度
lottery scheduling 彩票调度    利用 随机性 randomness 。避免边角情况，方法轻量，随机方法快。<br />彩票机制：彩票货币 ticket currency，彩票转让 ticket transfer，彩票通胀 ticket inflation 。<br />实现：
```c
int count = 0;

int winner = get_random(0, total_ticket_sum);

node_ticket *current = head;

while(current) {
    count += current->tickets;
    if (count > winner) break;
    current = current->next;
}
```
不公平指标 unfairness metric：两个工作完成时刻相除。
#### 步长调度 stride scheduling
Waldspurger 提出，确定性的公平分配算法。
```c
current = remove_min(queue);
schedule(current);
current->pass += current->stride;
insert(queue, current);
```
彩票调度对比步长调度的优势：不需要全局状态。
### 多处理器调度 multiprocessor schelduling
缓存一致性 cache coherence 问题，总线窥探 bussnooping。<br />互斥锁 pthread_mutext_t<br />缓存亲和度 cache affinity
#### 单队列多处理器调度 Single Queue Multiprocessor Scheduling, SQMS
容易构建，负载均衡较好<br />缺乏可扩展性 scalability<br />缓存亲和性    来回迁移 migrating
#### 多队列多处理器调度 Multi-Queue Multiprocessor Scheduling, MQMS
负载不均 load imbalance<br />让工作移动，迁移 migration<br />工作窃取 work stealing
#### Linux 多处理器调度
O(1)调度程序    多队列，基于优先级，类似MLFQ

完全公平调度程序 CFS    多队列，确定的比例调度方法，类似步长调度<br />BF 调度程序 BFS    单队列，基于比例调度，最早最合适虚拟截止时间优先算法 EEVEF<br />
## 抽象：地址空间
提供一个易用的物理内存抽象：地址空间 address space 。<br />利用栈 stack 保存当前函数的调用信息，分配空间给局部变量，传递参数和函数返回值。<br />堆 heap 管理动态分配的、用户管理的内存，静态初始化的变量，等等。
### 虚拟化内存
虚拟内存 VM 系统的一个主要目标 透明 transparency 。实现虚拟内存的方式，运行的程序不可见。<br />另一个目标 效率 efficiency 。依赖硬件支持，包括 TLB这样的硬件功能。<br />第三个目标 保护 protection 。
### 内存操作 API
#### malloc()
```c
int main () {
    // sizeof() 被认为编译时操作符，而不是运行时函数调用。
    // 为10个整数的数组声明了空间。
    int *x = malloc(10 * sizeof(int));
    // 32位 返回 4，64位 返回 8，sizeof 认为在获取一个整数的指针有多大，而不是为x分配了多少内存。
    printf("%d\n", sizeof(x));
    
    int y[10];
    // 此时，编译器有足够的静态信息，知道分配了40个字节。
    printf("%d\n", sizeof(y));
}
```
为字符串声明空间：malloc(strlen(s) + 1) ，strlen() 获取字符串长度，+1 为字符串结束符留出空间。<br />malloc() 返回一个指向 void 类型的指针。通过强制类型转换 cast 处理。
#### free()
接受一个 malloc() 返回的指针作为参数，释放内存。分配区域的大小不传入，由内存分配库本身记录追踪。
#### 常见错误

- **忘记分配内存    段错误 segmentation fault**
```c
char *src = "hello";
char *dst = (char *) malloc(strlen(src) + 1);
strcpy(dst, src); // 或者使用strdup()
```

- **没有分配足够的内存    缓冲区溢出 buffer overflow**
```c
char *src = "hello";
char *dst = (char *) malloc(strlen(src)); // too small
strcpy(dst, src); // 或者使用strdup()
```
常看起来运行正常，取决于如何实现 malloc 和许多其他细节。<br />字符串拷贝执行时，会在超过分配空间的末尾写入一个字节。**

- **忘记初始化分配的内存    未初始化的读取 uninitialized read**

会从堆中读取到未知的数据。

- **忘记释放内存    内存泄露 memory leak**
- **用完之前释放内存    悬挂指针 dangling pointer**
- **反复释放内存    重复释放 double free**
- **错误地调用 free()    无效的释放 invalid free**


<br />工具 purify & valgrind
#### 底层操作系统支持
malloc() 和 free() 不是系统调用，而是库调用。malloc 库管理虚拟地址空间内的空间，本身建立在一些系统调用之上。<br />一个这样的系统调用叫 brk，用于改变程序分断 break 的位置：堆结束的位置。需要一个参数（新分断的地址），根据新分断是大于还是小于当前分断，来增加或减少堆的大小。另一个调用 sbrk 要求传入一个增量，目的类似。<br />还可以通过 mmap() 调用获取内存。通过传入正确的参数，可以在程序中创建一个匿名 anonymous 内存区域，不与任何特定文件相关联，而是与交换空间 swap space 相关联，这种内存可以像堆一样对待并管理。<br />内存分配库还支持一些其他调用，calloc() 分配内存，并在返回前将其置零。如果你认为内存已归零并忘记自己初始化它，这可以防止一些错误。<br />当为某些东西分配空间，需要添加一些东西时，例程 realloc()：创建一个新的更大的内存区域，将旧区域复制到其中，并返回新区域的指针。
### 机制：地址转换
实现 CPU 虚拟化时，遵循的一般准则被称为受限直接访问 Limited Direct Execution LDE。<br />LDE 背后的想法：让程序运行的大部分指令直接访问硬件，只在一些关键点（如进程发起系统调用或发生时钟中断）由操作系统介入来确保在正确的时间，正确的地点，做正确的事。<br />为了实现高效的虚拟化，操作系统应该尽量让程序 自己运行，同时通过在关键点的及时介入 interposing，来保持对硬件的控制。高效和控制是现代操作系统的两个主要目标。<br />如何高效灵活地虚拟化内存？通用技术：有时称为基于硬件的地址转换 hardware-based address translation，简称为地址转换 address translation。将虚拟 virtual 地址转换为数据实际存储的物理 physical 地址。硬件只是提供了底层机制来提高效率。
#### 基址加界限机制 base and bound，动态重定位 dynamic relocation
每个 CPU 需要两个硬件寄存器：基址 base 寄存器和界限 bound 寄存器，有时称为限制 limit 寄存器。
```
physical address = virtual address + base
```
_基于软件的重定位：静态重定位 static relocation，一个名为加载程序 loader 的软件接手将要运行的可执行程序。不提供访问保护，一旦完成很难将内存空间重新定位到其他位置。_<br />地址转换发生在运行时，可以在进程开始运行后改变其地址空间，动态重定位 dynamic relocation。<br />CPU 负责地址转换的部分称为内存管理单元 Memory Management Unit MMU。<br />一个位，保存在处理器状态字 processor status word 中，说明当前 CPU 运行模式。<br />硬件必须提供基于基址和界限寄存器 base and bounds register，因此每个CPU的内存管理单元都需要这两个额外的寄存器。运行时，硬件转换每个地址，将程序产生的虚拟地址加上基址寄存器的内容。硬件也必须能检查地址是否有用，通过界限寄存器和CPU内的一些电路来实现。<br />硬件提供一些特殊指令，用于修改基址寄存器和界限寄存器，允许操作系统在切换进程时修改它们。这些指令是特权 privileged 指令，只能在内核模式下修改。<br />用户尝试非法访问内存（越界访问）时，CPU必须能够产生异常exception。阻止程序执行，安排异常处理程序 exception handler 去处理。

- 进程创建时，为进程的地址空间找到内存空间。
- 进程终止时，回收它的内存。
- 上下文切换时，保存和回复基础和界限寄存器。保存到内存中，或放在某种每个进程都有的结构中，如进程结构 process structure 或进程控制块 Process Control Block PCB 中。进程停止时，可改变其地址空间的物理位置。
- 操作系统必须提供 exception handler。
### 分段 segmentation
#### 泛化的基址/界限
在 MMU 中引入不止一个基址和界限寄存器对，而是地址空间内的每个逻辑段 segment 一对。<br />一个段只是地址空间里的一个连续定长的区域，在典型的地址空间里有3个逻辑不同的段：代码、栈和堆。<br />避免虚拟地址空间中的未使用部分占用物理内存。<br />只有已用的内存才在物理内存中分配空间，因此可以容纳巨大的地址空间，其中包含大量未使用的地址空间（有时又称为稀疏地址空间， sparse address spaces）。<br />**<br />**段异常 segmentation violation / 段错误 segmentation fault：支持分段的机器上发生了非法的内存访问。**
#### 如何知道段内偏移和地址引用的段
**显式 explicit 方法**。用虚拟地址开头几位标识不同的段。VAX/VMS 系统使用了这种技术。<br />_例：如果前2位是 00，硬件就知道这是属于代码段的地址，使用代码段的基址和界限来重定位到正确的物理地址。如果前2位是 01，则是堆地址，对应地，使用堆的基址和界限。一个虚拟地址，前两位决定使用哪个段寄存器，后12位作为段内偏移，偏移量和基址寄存器相加就得到物理地址。_<br />如果基址和界限放在数组中（每个段为一项）：
```c
// SEG_MASK = 0x3000
// SEG_SHIFT = 12
// OFFSET_MASK = 0xFFF

// get top 2 bits of 14-bit VA
Segment = (VirtualAddress & SEG_MASK) >> SEG_SHIFT
// now get offset
Offset = VirtualAddress & OFFSET_MASK
if (Offset >= Bounds[Segment])
    RaiseException(PROTECTION_FAULT)
else
    PhysAddr = Base[Segment] + Offset
    REgister = AccessMemory(PhysAddr)
```
有些系统会将堆和栈当作同一个段，因此只需要一位来标识。<br />**隐式 implicit 方式**<br />硬件通过地址产生的方式来确定段。<br />_例：地址如果由程序计数器产生（即它是指令获取），那么地址在代码段。如果基于栈或基址指针，它一定在栈段。其它地址则在堆段。_
#### 空闲列表管理算法
紧凑 compact 内存<br />最优匹配 best-fit，从空闲链表中找最接近需要分配空间的空闲块返回<br />最坏匹配 worst-fit<br />首次匹配 first-fit<br />下次匹配 next fit<br />伙伴算法 buddy algorithm<br />
<br />分段帮助实现更高效的虚拟内存，动态重定位，避免地址空间的逻辑段之间潜在的内存浪费，更好的支持稀疏地址空间。代码共享。
### 空闲空间管理 free-space management
malloc 库 管理进程中堆的页<br />操作系统  管理进程的地址空间<br />外部碎片 external fragmentation 问题：空闲空间被分割城不同大小的块。<br />
<br />malloc 库管理的空间由于历史原因称为堆，堆上管理空闲空间的数据结构称为空闲列表 free list，包含了管理内存区域中所有空闲块的引用。<br />分配程序给出的内存块超出了请求的大小，被认为是内部碎片 internal fragmentation。<br />
<br />分离空闲列表 segregated list<br />厚块分配程序 slab allocator<br />
<br />伙伴系统<br />二分伙伴分配程序 binary buddy allocator。有内部碎片 internal fragment<br />

### 分页
地址空间分割成固定大小的单元，页<br />物理内存看作定长槽块的阵列，页帧 page frame<br />每个页帧包含一个虚拟内存页<br />
<br />操作系统为每个进程保存一个数据结构，称为页表 page table。<br />页表为地址空间的每个虚拟页面保存地址转换 address translation。<br />页表是每个进程的数据结构，除倒排页表外 inverted page table。<br />
<br />从地址 virtual address 到寄存器 eax 的数据显式加载。<br />为了转换 translate 该过程生成的虚拟地址，分为两个组件：虚拟页面号 virtual page number 和页内的偏移量 offset。<br />
<br />分页特点：灵活性，高效的抽象地址空间。空闲空间管理的简单性。<br />操作系统保存一个所有空闲页的空闲列表 free list。<br />
<br />物理帧号 PFN    物理页号 physical page number PPN<br />页表格条目 PTE<br />线性页表 linear page table<br />物理帧号 PFN<br />页面替换 page replacement<br />
<br />有效位 valid bit：通常用于指示特定地址转换是否有效。<br />保护位 protection bit：表明页是否可以读写或执行。<br />存在位 present bit：表示该页是在物理存储器还是磁盘上（换出 swapped out）。<br />脏位 dirty bit：表明页面被带入内存后是否被修改过。<br />参考位 reference bit | 访问位 accessed bit：有时用于追踪页是否被访问，也用于确定哪些页很受欢迎，因该保留在内存中。<br />
<br />页表基址寄存器 page-table base register：包含页表的起始位置的物理地址。<br />

#### 快速地址转换 TLB
地址转换旁路缓冲存储器 translation-lookasidebuffer TLB，频繁发生的虚拟到物理地址转换的硬件缓存。更好的名称：地址转换缓存 address-translation cache。<br />
<br />线性页表 linear page table，页表是一个数组，和硬件管理的 TLB hardware-managed TLB。
```c
VPN = (VirtualAddress & VPN_MASK) >> SHIFT
(Success, TlbEntry) = TLB_Lookup(VPN)
if (Success == True) // TLB Hit
    if (CanAccess(TlbEntry.ProtectBits) == True)
        Offset = VirtualAddress & OFFSET_MASK
        PhysAddr = (TlbEntry.PFN << SHIFT) | Offset
        AccessMemory(PhysAddr)
    else
        RaiseException(PROTECTION_FAULT)
else // TLB Miss
    PTEAddr = PTBR + (VPN * sizeof(PTE))
    PTE = AccessMemory(PTEAddr)
    if (PTE.Valid == False)
        RaiseException(SEGMENTATION_FAULT)
    else if (CanAccess(PTE.ProtectBIts) == False)
        RaiseException(PROTECTION_FAULT)
    else
        TLB_Insert(VPN, PTE.PFN, PTE.ProtectBits)
        RetryInstruction()
```

<br />_硬件缓存背后的思想是利用指令和数据引用的局部性 locality。_<br />_时间局部性 temporal locality：最近访问过的指令或数据项可能很快会再次访问。_<br />_空间局部性 spatial locality：当访问内存地址 x 时，可能很快会访问邻近的 y 内存。_<br />_
#### 如何处理 TLB 未命中
硬件或软件（操作系统）。<br />
<br />以前的硬件有复杂的指令集（也称为复杂指令集计算机，Complex-Instruction Set Computer，CISC），硬件全权处理 TLB 未命中。硬件必须知道页表在内存中的位置，通过页表基址寄存器，page-table base register，以及页表的确切格式。<br />未命中时，遍历页表，找到正确的页表项，取出想要的转换映射，用它更新 TLB。<br />一个例子是 x86 架构，采用固定的多级页表 multi-level page table，当前页表由 CR3 寄存器指出。<br />
<br />精简指令集计算机 Reduced-Instruction Set Computer，RISC，有所谓的软件管理 TLB software-managed TLB。<br />未命中时，硬件系统抛出一个异常，暂停当前的指令流，将特权级提升至内核模式，跳转至陷阱处理程序 trap handler。运行操作系统一段代码，查找页表中的转换映射，用特权指令更新 TLB，从陷阱返回。硬件重试该指令，命中 TLB。
```c
VPN = (VirtualAddress & VPN_MASK) >> SHIFT
(Success, TlbEntity) = TLB_Lookup(VPN)
if (Success == True) // TLB Hit
    if (CanAccess(TlbEntry.ProtectBits) == True)
        Offset = VirtualAddress & OFFSET_MASK
        PhysAddr = (TlbEntry.PFN << SHIFT) | Offset
        AccessMemory(PhysAddr)
    else
        RaiseException(PROTECTION_FAULT)
else // TLB Miss
    RaiseException(TLB_MASK)
```
系统陷入内核时保存不同的程序计数器。<br />避免无限递归。把 TLB 未命中陷阱处理程序直接放到物理内存中（它们没被映射过 unmapped，不经过地址转换）。或者在 TLB 中保留一些项，记录永久有效的地址转换，并将一些永久地址转换槽块留给处理代码本身，这些被监听的 wired 地址转换总会命中 TLB。<br />
<br />典型的 TLB 有 32，64，128 项，并且是全相联的的 fully associative 缓存。<br />一条 TLB 项内容：VPN | PFN | 其他位<br />VPN 和 PFN 同时存在于 TLB 中，因为一条地址映射可能出现在任意位置。<br />
<br />TLB 有一个有效位标识该项是否有效的地址转换。还有一些保护位，标识该页是否有访问权限。其他位包括地址空间标识符 address-space identifier、脏位 dirty bit 等。<br />TLB 的有效位 != 页表的有效位：<br />页表项 PTE 被标记为无效，就意味着该页并没有被进程申请使用，正常运行的程序不应该访问该地址。程序试图访问时，会陷入操作系统，操作系统会杀掉该进程。<br />TLB 的有效位指出 TLB 项是否有效的地址映射。在系统上下文切换时起重要作用。<br />
<br />上下文切换时清空 flush TLB，开销大。<br />增加硬件支持，实现跨上下文切换的 TLB 共享。在 TLB 中添加一个地址空间标识符 Address Space Identifier，ASID。可以把 ASID 看作是进程标识符 Process Identifier PID，但比 PID 位数少。<br />
<br />共享代码页减少内存开销。<br />
<br />访问 TLB 容易成为 CPU 流水线瓶颈，尤其是所谓的物理地址索引缓存 physically-idnexed cache。有了这种缓存，地址转换必须发生在访问该缓存之前，会让操作变慢。用虚拟地址直接访问缓存，避免地址转换。虚拟地址索引缓存 virtually-indexed cache。<br />

