def BubbleSort(Iterator):
    """ 冒泡排序 """
    a = list(Iterator)
    for i in range(len(a)):
        for j in range(len(a)-1-i):
            if a[j] > a[j+1]:
                a[j], a[j+1] = a[j+1], a[j]
    return a

def InsertSort(Iterator):
    """ 插入排序 """
    a = list(Iterator)
    for i in range(1, len(a)):
        for j in range(i, 0, -1):
            if a[j] < a[j-1]:
                a[j], a[j-1] = a[j-1], a[j]
    return a

def main():
    data = [74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586]
    print(BubbleSort(data))
    print(InsertSort(data))

if __name__ == '__main__':
    main()