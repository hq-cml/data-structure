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
    InOrder(root->left);
    printf("%d ", root->data);
    InOrder(root->right);
}

void PostOrder(node *root) {
    if (!root) {
        return;
    }
    PostOrder(root->left);
    PostOrder(root->right);
    printf("%d ", root->data);
}

node *BuildTree(int *preOrder, int *inOrder, int len) {
    if(len<=0) {
        return NULL;
    }

    if(len == 1) {
        return NewNode(preOrder[0]);
    }

    node *root = NewNode(preOrder[0]);
    int idx, i;
    for(i=0; i<len; i++) {
        if(inOrder[i] == preOrder[0]) {
            idx = i;
            printf("A------%d, %d\n", idx, preOrder[0]);
            break;
        }
    }

    if(idx == 0) {
         root->left = NULL;
    }

    if(idx == len-1) {
        root->right = NULL;
    }

    if(idx > 0) {
        root->left = BuildTree(preOrder+1, inOrder, idx);
    }

    if(idx < len-1) {
        root->right = BuildTree(preOrder + idx +1, inOrder + idx + 1, len - idx -1);
    }

    return root;
}

int main() {
//    node *root = NewNode(1);
//    root->left = NewNode(2);
//    root->right = NewNode(3);
//
//    PreOrder(root);
//    printf("\n");
//    InOrder(root);
//    printf("\n");
//    PostOrder(root);
//    printf("\n");

    int pre[] = {1, 2, 4, 7, 3, 5, 6, 8};
    int in[] = {4, 7, 2, 1, 5, 3, 8, 6};
    node *root = BuildTree(pre, in, 8);
    PreOrder(root);
    printf("\n");
    InOrder(root);
    printf("\n");
    PostOrder(root);
}