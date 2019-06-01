def bubble_sort(nums):
    for i in range(len(nums)):
        for j in range(len(nums)-i-1):
            if nums[j] > nums[j+1]:
                nums[j], nums[j+1] = nums[j+1], nums[j]
    return nums


def main():
    input = [
        [1, 2, 3, 4, 5],
        [5, 4, 3, 2, 1],
        [2, 4, 1, 3, 5],
    ]

    expected = [1, 2, 3, 4, 5]

    for i in input:
        print('Before: {}'.format(i))
        sorted_list = bubble_sort(i)
        print('After: {}'.format(sorted_list))
        print()

        assert expected == [1, 2, 3, 4, 5]


if __name__ == '__main__':
    main()
