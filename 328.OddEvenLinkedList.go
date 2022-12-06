/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    //if list is empty
    if(head==nil){
        return nil
    }

    even,evenHead:=head.Next,head.Next
    odd:=head

    for even!=nil && even.Next!=nil {
        
        //for odd list
        odd.Next=even.Next
        odd=odd.Next

        //for even list
        even.Next=odd.Next
        even=even.Next

    }
    //join both the lists
    odd.Next=evenHead
    return head

}
