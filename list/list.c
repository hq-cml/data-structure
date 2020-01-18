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

int main() {
    node *head = insert(NULL, 1);
    head = insert(head, 2);
    head = insert(head, 3);
    head = insert(head, 4);
    foreach(head);

    head = reverse(head);
    foreach(head);

    //head = doubleSwap(head);
    //foreach(head);
}
