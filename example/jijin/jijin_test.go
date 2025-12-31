package jijin

import (
	"fmt"
	"testing"
)

func (f *Fund) ExecCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentMoney float64) {
	title := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
		"成本单价", "初买数量", "补仓净值", "补仓金额", "补仓数量", "最终成本", "最终数量", "百分4净值")
	fmt.Println(title)
	cover := newCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentMoney)
	cover.ToString()
	for i := 0; i < (f.ReplenishmentTimes - 1); i++ {
		cover = newCover(cover.finalCost, cover.finalQuantity, cover.netWorth4, replenishmentMoney)
		cover.ToString()
	}
}

// 兴全合泰A 007802
func TestXingQuan(t *testing.T) {
	var costPrice = 1.4352                          //成本单价
	var initialPurchaseQuantity = float64(16654.96) // 初买入数量
	var netReplenishmentValue = 1.2721              //补仓净值
	var replenishmentMoney = float64(1000)          // 补仓金额
	fund := NewFund("兴全合泰", 40)
	fund.ExecCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentMoney)
}

// 中欧时代先锋 001938
func TestZhongO(t *testing.T) {
	var costPrice = 1.5234                          //成本单价
	var initialPurchaseQuantity = float64(34162.09) // 初买入数量
	var netReplenishmentValue = 1.2618              //补仓净值
	var replenishmentMoney = float64(1000)          // 补仓数量
	fund := NewFund("中欧时代先锋", 40)
	fund.ExecCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentMoney)
}

// 交银新成长混合 519736
func TestJiaoYin(t *testing.T) {
	var costPrice = 3.9368                          //成本单价
	var initialPurchaseQuantity = float64(19533.50) // 初买入数量
	//var netReplenishmentValue = 3.3760              //补仓净值
	var netReplenishmentValue = 3.1840        //补仓净值
	var replenishmentQuantity = float64(5000) // 补仓数量
	fund := NewFund("交银新成长混合", 40)
	fund.ExecCover(costPrice, initialPurchaseQuantity, netReplenishmentValue, replenishmentQuantity)
}

func TestDecimal(t *testing.T) {
	fmt.Println(Decimal(1.3451407128861068))
}
