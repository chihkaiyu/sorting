def insertion_sort(nums):
    for i in range(1, len(nums)):
        for j in range(i, 0, -1):
            if nums[j] < nums[j-1]:
                nums[j], nums[j-1] = nums[j-1], nums[j]
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
        sorted_list = insertion_sort(i)
        print('After: {}'.format(sorted_list))
        print()

        assert sorted_list == expected


if __name__ == '__main__':
    main()
