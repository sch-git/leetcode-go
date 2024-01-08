# [环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/description/?envType=study-plan-v2&envId=top-100-liked)

给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

> 快慢指针的方式<br>
> 设 fast、slow，fast 每次走两步，slow 每次走一步<br>
> 如果存在环，当指针相遇时，fast 走了 2s，slow 走了s<br>
> 设链表到环入口为 a，环长度为b，则 s=a+nb<br>
> **将 fast 移动到 head，再次相遇即为环入口**，走了 a步