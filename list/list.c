#include <stdio.h>
#include <stdlib.h>

typedef struct _node{
    int data;
    struct _node *next;
} node;

node *insert(node *head, int data) {
    node *p = (node *)malloc(sizeof(node));
    if (!p) {
        return head;
    }
    p->data = data;
    p->next = NULL;

    if (!head) {
        return p;
    }

    node *p1 = head;
    while(p1->next) {
        p1 = p1->next;
    }

    p1->next = p;
    return head;
}

void foreach(node *head) {
    if (!head) {
        printf("Empty!");
    }
    node *p = head;
    while(p) {
        printf("%d ", p->data);
        p = p->next;
    }
    printf("\n");
}

//翻转
node* reverse(node *head) {
    if (!head || !head->next) {
        return head;
    }

    node *p1 = head;
    node *p2 = head->next;
    node *p3;
    while(p2 != NULL) {
        p3 = p2->next;
        p2->next = p1;
        p1 = p2;
        p2 = p3;
    }

    head->next = NULL;
    head = p1;
    return head;
}

//奇偶交换
node* doubleSwap(node *head) {
    if (!head || !head->next) {
        return head;
    }

    node *p1 = head;
    node *p2 = head->next;
    node *p3 = NULL;
    node *p4 = NULL;
    head = p2;
    while(p1 != NULL && p2 != NULL) {
        p3 = p2->next;
        p2->next = p1;
        p1->next = p3;
        if (p4 != NULL) {
            p4->next = p2;
        }

        if (p3 != NULL && p3->next != NULL) {
            p4 = p1;
            p1 = p3;
            p2 = p3->next;
        } else {
            p1 = p2 = NULL;
        }
    }

    return head;
}

//倒数第n个节点
node *findBackNode(node *head, int n) {
    if(!head) {
        printf("HEAD NULL\n");
        return NULL;
    }

    if(n<=0) {
        printf("Invalid N!\n");
        return NULL;
    }

    node *p = head;
    node *q = head;
    int m = n-1;
    while(p && m>0) {
        p = p->next;
        m --;
    }

    if (p != NULL) {
        while(p && p->next) {
            q = q->next;
            p = p->next;
        }
        printf("The back [%d] is %d!\n", n, q->data);
        return q;
    } else {
        printf("Invalid N!\n");
        return NULL;
    }
}

//链表倒着打印
void RevertPrint(node *head) {
    if (head == NULL) {
        return;
    }
    RevertPrint(head->next);
    printf("%d ", head->data);
}

//O(1)复杂度删除链表节点
node *DeleteNodeO1(node *head, node *del) {
    if (head == NULL || del == NULL) {
        return head;
    }

    node *tmp;
    node *pre;
    if(del->next != NULL) {
        del->data = del->next->data;
        tmp = del->next;
        del->next = del->next->next;
        free(tmp);
    } else {
        //常规删除
        if(head == del) { //第一个结点
            tmp = head;
            head = head->next;
            free(tmp);
        } else {    //最后一个节点
            pre = head;
            while(pre->next != NULL) {
                pre = pre->next;
            }

             pre = pre->next;
             free(tmp);
        }
    }
    return head;
}

//删除重复节点
node *DeleteDuplicateNode(node *head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }

    node *tmp;
    node *first = head;
    node *curr = first->next;
    while(first) {
        while(curr!=NULL && curr->data == first->data) {
            tmp = curr;
            first->next = curr->next;
            curr = curr->next;
            printf("Del %d\n", tmp->data);
            free(tmp);
        }
        if (curr == NULL) {
            break;
        }
        first = curr;
        curr = curr->next;
    }
    return head;
}

int main() {
    node *head = insert(NULL, 1);
    node *p = head;
    int n;

    head = insert(head, 2);
    head = insert(head, 3);
    head = insert(head, 3);
    head = insert(head, 4);
    head = insert(head, 5);
    head = insert(head, 5);
    head = insert(head, 5);
    head = insert(head, 6);
    foreach(head);

    head = DeleteDuplicateNode(head);
    foreach(head);
    //head = reverse(head);
    //foreach(head);

    //head = doubleSwap(head);
    //foreach(head);

//    findBackNode(head, 1);
//    findBackNode(head, 2);
//    findBackNode(head, 6);
//    findBackNode(head, 7);
    //RevertPrint(head);
}
