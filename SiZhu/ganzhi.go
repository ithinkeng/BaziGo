package SiZhu

// 干支

// 获得八字年的干支，0-59 对应 甲子到癸亥
func GetGanZhiFromYear(nYear int) int {
	if nYear > 0 {
		return (nYear - 4) % 60
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear - 3) % 60
	}
}

// 获得某公/农历年的天干，0-9 对应 甲到癸
func GetGanFromYear(nYear int) int {
	if nYear > 0 {
		return (nYear - 4) % 10
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear - 3) % 10
	}
}

// 获得某公/农历年的地支，0-11 对应 子到亥
func GetZhiFromYear(nYear int) int {
	if nYear > 0 {
		return (nYear - 4) % 12
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear - 3) % 12
	}
}

// 将干支拆分成天干地支，0-59 转换成 0-9 0-11
func ExtractGanZhi(GanZhi int) (int, int) {
	return (GanZhi % 10), (GanZhi % 12)
}

// 将天干地支组合成干支，0-9 0-11 转换成 0-59
func CombineGanZhi(nGan, nZhi int) int {
	if (nGan <= 9) && (nZhi <= 11) {
		for i := 0; i <= 6; i++ {
			if (i*10+nGan)%12 == nZhi {
				return i*10 + nGan
			}
		}
	}
	return -1
}