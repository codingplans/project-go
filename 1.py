#
# amounts = [1] * 10 + [0] * 40
#
# print "12313";
# print('两数之和为 %.1f' %(float(input('输入第一个数字：'))-float(input('输入第二个数字：'))))

import cmath

num = int(input("请输入一个数字: "))
num_sqrt = cmath.sqrt(num)
print('{0} 的平方根为 {1:0.3f}+{2:0.3f}j'.format(num ,num_sqrt.real,num_sqrt.imag))
