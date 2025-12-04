# Binary Tree Traversal Implementations in Go

<img width="800" height="400" alt="image" src="https://github.com/user-attachments/assets/d5e6d734-694f-4cc8-8a07-cd591a52b99a" />

Bu repository, Go dilinde Binary Tree (İkili Ağaç) gezinme algoritmalarının implementasyonlarını içerir.

## İçerik

Binary Tree traversal için 4 temel yöntem implement edilmiştir:

### 1. Preorder Traversal (Ön Sıralı)
**Sıra:** Root → Left → Right

Önce kök node ziyaret edilir, sonra sol alt ağaç, en son sağ alt ağaç gezilir.

**Kullanım Alanları:**
- Ağaç kopyalama
- Prefix expression oluşturma

### 2. Inorder Traversal (İç Sıralı)
**Sıra:** Left → Root → Right

Önce sol alt ağaç, sonra kök node, en son sağ alt ağaç ziyaret edilir.

**Kullanım Alanları:**
- Binary Search Tree'de sıralı liste elde etme
- BST'de elemanları küçükten büyüğe yazdırma

### 3. Postorder Traversal (Son Sıralı)
**Sıra:** Left → Right → Root

Önce her iki alt ağaç gezilir, en son kök node ziyaret edilir.

**Kullanım Alanları:**
- Ağaç silme işlemleri
- Dosya sistemi boyutu hesaplama

### 4. Level Order Traversal (Seviye Sıralı - BFS)
**Sıra:** Seviye seviye, soldan sağa

Ağaç seviye seviye, her seviyede soldan sağa doğru gezilir.

**Kullanım Alanları:**
- En kısa yol bulma
- Minimum depth hesaplama
- Breadth-First Search (BFS)

## Örnek
```
        1
       / \
      2   3
     / \
    4   5
```

**Çıktılar:**
- Preorder: `[1, 2, 4, 5, 3]`
- Inorder: `[4, 2, 5, 1, 3]`
- Postorder: `[4, 5, 2, 3, 1]`
- Level Order: `[1, 2, 3, 4, 5]`

## Implementasyon Detayları

- **Preorder, Inorder, Postorder:** Stack kullanarak (DFS - Depth First Search)
- **Level Order:** Queue kullanarak (BFS - Breadth First Search)

## Komplekslik

- **Zaman Karmaşıklığı:** O(n) - Her node bir kez ziyaret edilir
- **Alan Karmaşıklığı:** O(h) - h: ağaç yüksekliği (recursive stack için)

## Kaynaklar

- [LeetCode - Binary Tree Traversal Problems](https://leetcode.com/)
- [GeeksforGeeks - Tree Traversals](https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/)

---
