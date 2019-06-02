def heap_sort(heap):
    sorted_heap = []

    while len(heap) > 1:
        heap[0], heap[-1] = heap[-1], heap[0]
        sorted_heap.append(heap.pop())

        min_heapify(heap, 0)

    sorted_heap.append(heap.pop())
    return sorted_heap


def build_min_heap(heap):
    last_parent = int((len(heap) - 2) / 2)
    for i in range(last_parent, -1, -1):
        min_heapify(heap, i)


def min_heapify(heap, parent):
    minimum_child = parent * 2 + 1
    if minimum_child >= len(heap):
        return

    right = minimum_child + 1

    # compare left with right child if we have right child
    if right < len(heap) and heap[right] < heap[minimum_child]:
        minimum_child += 1

    # compare the largest child with parent
    if heap[minimum_child] < heap[parent]:
        heap[minimum_child], heap[parent] = heap[parent], heap[minimum_child]
        # recursively min heapify the child we swap
        min_heapify(heap, minimum_child)


def main():
    test = [
        [1, 2, 3, 4, 5],
        [5, 4, 3, 2, 1],
        [2, 4, 1, 3, 5],
    ]

    expected = [1, 2, 3, 4, 5]

    for i in test:
        print('Before: {}'.format(i))
        build_min_heap(i)
        sorted_heap = heap_sort(i)
        print('After: {}'.format(sorted_heap))
        print()

        assert sorted_heap == expected


if __name__ == '__main__':
    main()
