package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GAQFBANZkACIgycDMXlienpqUX6UFovqzg/LzSElYHRKU4h8U3YwbzLDgJ7v2uvaU2e6mL77JhgZ7CgesuFa/Msbh/bffsmr6qTXQDLKgGT6duEmBRDOa5Ed1k8aQioZ9v9k2dLWdCrdW/dtmoKfFDsSp+fbvv+rfT5yarHC2UWbVvxnnnSNucwKd1Nc9wk33NcKgvyESpkvX7lqpqE2sUCrqkh5aFit99oH0jz/9GTLha0J+BW1zJtjT37uNVu1mabOgZoVp9ZOovzXkx4kzD3t4Sb51avjJ/0ueS9xj3J4+/j3PX5n59e/svRoTBBw7r2s+mvbf5FDpOUeB2YxfJzgw12d7CfclXzXZg2+2SGTP29vrbbH/8VnktL7y+UajjIfN2H/9SETqMU34qlGi1n7Zw/eNy+eKNdaEnRboELMbHGO2ZzyG7+3LuMXfbYzJUBpswPHdcXem+yfnZmzaFlukfecBWHv8mYV8DhxbXOQXSfemTmVq/uDVMsb3bf27c/ZmbWn4CTUsFZIaVZmybOL3/4yZXbW0T2GvPtI3YiFT/Eq3I3Wfe1sWd2nXwR/OfNz++TD/94qXG+tn7tpZXC7zQPuLy/vYw/fvEDDlmWq7Ef1sm+7XncnKxr1qB/5eHP2Fymg+le6s5HlORnvzwc37dIfP2ju/0XPxz+z19/VHj2tt27G/pncK++nPxFv/7/Vwc1J63ent9XMzVMCj/lmx2eP1nF+mS6VdfTdzdPd20tLfja+qn3sOqCNNYS4QodyVubj978lRrSUvnOv+rX1WCxjaXrnt7cnHV2j8/emqtCRdGTgqKZ9nJFPldTO3663sZ+zk5p/tf+1XO3pKsfmyWZXjzlwA42+QK7tFfPd/Fxdp+I86o/LM/IwPD/f4A3O0eViOSOGCYGBiNBBgZYMmVgEEJLpuyIZApOmS5xCokg3chqArwZmUSYEckc2WRQMoeB/40gEm+iRxiF3SkQIMDw33E5EwMWh7GygeSZGJgYOhkYGJ4ygXiAAAAA//9r6tY9hAMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
