# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result = ListNode(0)
        temp = result
        while l1 or l2:
            if l1:
                temp.val += l1.val
                l1 = l1.next
            if l2:
                temp.val += l2.val
                l2 = l2.next

            if temp.val > 9:
                temp.next = ListNode(val=(temp.val // 10))
                temp.val %= 10
            elif l1 or l2:
                temp.next = ListNode()
            temp = temp.next

        return result

