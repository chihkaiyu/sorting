def selection_sort(nums):
    for i in range(len(nums)):
        minimum_index = i
        for j in range(i, len(nums)):
            if nums[j] < nums[minimum_index]:
                minimum_index = j
        nums[i], nums[minimum_index] = nums[minimum_index], nums[i]
    return nums


def main():
    test = [
        [1, 2, 3, 4, 5],
        [5, 4, 3, 2, 1],
        [2, 4, 1, 3, 5],
    ]

    expected = [1, 2, 3, 4, 5]

    for i in test:
        print('Before: {}'.format(i))
        sorted_list = selection_sort(i)
        print('After: {}'.format(sorted_list))
        print()

        assert sorted_list == expected


if __name__ == '__main__':
    main()
