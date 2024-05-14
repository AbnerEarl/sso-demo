/**
 * @author: yangchangjia
 * @email 1320259466@qq.com
 * @date: 2024/5/13 09:30
 * @desc: about the role of class.
 */

package main

import (
	"github.com/AbnerEarl/sso/provider"
)

func main() {
	provider.StartProvider("provider/config.yaml")
}
