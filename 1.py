import cmath

num = int(input("请输入一个数字: "))
num_sqrt = cmath.sqrt(num)
：print('{0} 的平方根为 {1:0.3f}+{2:0.3f}j'.format(num ,num_sqrt.real,num_sqrt.imag))




def eat():
	return '我在吃饭'

def work():
        return '我做工作、。。。'

def play():
        return '我在玩！！'

print(eat())
print(work())
print(play())














