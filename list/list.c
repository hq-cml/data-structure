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

//特殊的插入，用于构造环
node *insertTmp(node *head, int data) {
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
    return p;      //返回的是实际的新增的那个节点
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
    head = p2;      //最终头节点是原来第二个节点
    while(p1 != NULL && p2 != NULL) {
        p3 = p2->next;
        p2->next = p1;
        p1->next = p3;
        if (p4 != NULL) {
            p4->next = p2;   //这个地方必须要更换前面已调整的最后一个节点的后继指针，否则就断了
        }

        if (p3 != NULL && p3->next != NULL) {
            p4 = p1;
            p1 = p3;
            p2 = p3->next;
        } else {
            p1 = p2 = NULL; //结束，只剩下0个或者1个节点
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

//判断链表是否存在环，如果存在，返回环长度
int CheckLIstRing(node *head) {
    if (head == NULL || head->next == NULL) {
        return 0;
    }

    node *p1, *p2;
    p1 = p2 = head;
    if (p2->next->next == p1) {
        return 2; //两个节点，自成环的特例
    }
    int find = 0;
    while(p2->next && p2->next->next) {
        p2 = p2->next->next;
        p1 = p1->next;
        if(p2 == p1) {
            find = 1;
            break;
        }
    }
    if (!find) {
        return 0;//不存在环
    }

    int cnt = 0;
    do {
        cnt ++;
        p2 = p2->next->next;
        p1 = p1->next;
    } while(p1 != p2);
    return cnt;
}

//求链表存在环的话，环的入口
node* FindRingEntry(node *head) {
    int cnt = CheckLIstRing(head);
    if (cnt == 0) {
        return NULL;
    }

    int i;
    node *p1, *p2;
    p1 = p2 = head;
    for(i=0; i<cnt; i++) {
        p2 = p2->next;
    }
    while(p1 != p2) {
        p1 = p1->next;
        p2 = p2->next;
    }
    return p1;
}

//合并有序链表
node *Merge(node *head1, node *head2) {
    if (head1==NULL) return head2;
    if (head2==NULL) return head1;

    node *head = NULL;
    node *p1 = head1;
    node *p2 = head2;
    node *p = NULL;

    //合并
    while(p1 && p2) {
        if (p1->data <= p2->data) {
            if (head == NULL) {
                head = p1;
                p = p1;
            } else {
                p->next = p1;
                p = p->next;
            }
            p1 = p1->next;
        } else {
            if (head == NULL) {
                head = p2;
                p = p2;
            } else {
                p->next = p2;
                p = p->next;
            }
            p2 = p2->next;
        }
    }

    //扫尾
    if(p1) {
        p->next = p1;
    } else if(p2) {
        p->next = p2;
    }

    return head;
}

//链表的归并排序
node *MergeSort(node *head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }
    node *p1 = head;
    node *p2 = head->next->next;

    while(p1 && p2) {
        p1 = p1->next;
        p2 = p2->next;
        if (p2) {
            p2 = p2->next;
        }
    }
    //截断成两根
    p2 = p1->next;
    p1->next = NULL;
    return Merge(MergeSort(head), MergeSort(p2));
}

node *NewNode(int data) {
    node *p = (node *)malloc(sizeof(node));
    p->data = data;
    p->next = NULL;
    return p;
}

//另类链表合并，实现类似于加法的操作
//例如输入：head1 = [1,3,5,7] head2 = [2,4,6,8]
//输出： head = [3, 7, 1, 6, 1]
//例如输入：head1 = [1,3,5,7] head2 = [2,4,6,8,9]
//输出： head = [3, 7, 1, 6, 0, 1]
node *AddMerge(node *head1, node *head2){
    if (head1 == NULL) {
        return head2;
    }

    if (head2 == NULL) {
        return head1;
    }

    node *head = NULL;
    int flag = 0;
    node *p1 = head1;
    node *p2 = head2;
    node *tail;
    while(p1 && p2) {
        int tmp = p1->data + p2->data;
        if (flag == 1) {
            tmp ++;
        }
        if (tmp>=10) {
            flag = 1;
            tmp -= 10;
        } else {
            flag = 0;
        }
        node *p = NewNode(tmp);
        if(head == NULL) {
            head = p;
            tail = p;
        } else {
            tail->next = p;
            tail = p;
        }
        p1 = p1->next;
        p2 = p2->next;
    }

    if (p1 == NULL && p2 == NULL) {
        if (flag == 1) {
            node *p = NewNode(1);
            tail->next = p;
            tail = p;
            flag = 0;
        }
    } else {
        node *final = NULL;
        if (p1 != NULL) {
            final = p1;
        } else {
            final = p2;
        }

        while(final) {
            int tmp = final->data;
            if (flag == 1) {
                tmp ++;
            }
            if (tmp>=10) {
                flag = 1;
                tmp -= 10;
            } else {
                flag = 0;
            }
            node *p = NewNode(tmp);
            tail->next = p;
            tail = p;
            final = final->next;
        }
    }

    //还有可能有最后一个进位
    if (flag == 1) {
        node *p = NewNode(1);
        tail->next = p;
        tail = p;
    }

    return head;
}

int main() {
    //构造环
//    node *head = insert(NULL, 1);
//    node *p1, *p2;
//    int n;
//    head = insert(head, 2);
//    p1 = insertTmp(head, 3);
//    head = insert(head, 4);
//    head = insert(head, 5);
//    p2 = insertTmp(head, 6);
//    p2->next = p1;
//
//    printf("%d\n", CheckLIstRing(head));
//    printf("%d\n", FindRingEntry(head)->data);
    node *head1 = insert(NULL, 1);
    head1 = insert(head1, 3);
    head1 = insert(head1, 5);
    head1 = insert(head1, 7);

    node *head2 = insert(NULL, 2);
    head2 = insert(head2, 4);
    head2 = insert(head2, 6);
    head2 = insert(head2, 8);
    head2 = insert(head2, 9);
    foreach(AddMerge(head1, head2));
}
