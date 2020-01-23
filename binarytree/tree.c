#include <stdio.h>
#include <stdlib.h>

typedef struct _node{
    int data;
    struct _node *left;
    struct _node *right;
} node;

node *NewNode(int d) {
    node *p = (node*)malloc(sizeof(node));
    if(p == NULL) {
        printf("Out of memory!!\n");
    }
    p->data = d;
    p->left = p->right = NULL;
    return p;
}

void PreOrder(node *root) {
    if (!root) {
        return;
    }
    printf("%d ", root->data);
    PreOrder(root->left);
    PreOrder(root->right);
}

void InOrder(node *root) {
    if (!root) {
        return;
    }
    PreOrder(root->left);
    printf("%d ", root->data);
    PreOrder(root->right);
}

void PostOrder(node *root) {
    if (!root) {
        return;
    }
    PreOrder(root->left);
    PreOrder(root->right);
    printf("%d ", root->data);
}

int main() {
    node *root = NewNode(1);
    root->left = NewNode(2);
    root->right = NewNode(3);

    PreOrder(root);
    printf("\n");
    InOrder(root);
    printf("\n");
    PostOrder(root);
    printf("\n");
}