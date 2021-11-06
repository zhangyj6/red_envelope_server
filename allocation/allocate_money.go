/*上界upper用于确定锦鲤的范围，那就先取出一部分钱作为锦鲤，
锦鲤个数n = 0.1%*num，金额在范围[avg*20,upper]中随机
一次性随机一个长度为n的数组表示锦鲤的红包id
total_money -= 锦鲤总钱数，num -= 0.1%num
下界lower用于保底，先减去下界total_money -= num * lower
普通红包的策略就是在[1,2倍均值]之间随机

取钱函数：判断红包id是否为锦鲤，若不是，avg = total_money/num，
money = rand(1,2*avg)  total_money -= money  num--
return money + lower
*/
package allocation

import "fmt"

type Allocation struct{
	total_money int
	num int
	upper, lower int
}
/*func AllocateMoney(total_money int, num int, max_money int, lower int) {
	avg := float64(total_money) / float64(num)

	for ; num > 0; num-- {
	}
}*/
func main(){
	var a Allocation
	//a := Allocation(500000000000, 200000000000, 5, 66600)
	fmt.Printf("%v\n",a)
}