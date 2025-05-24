

# oms-v2

## 订单管理系统


订单
    HTTP
        创建订单
        查询订单
    gRPC
        修改订单


库存
    gRPC
        查询库存
        扣减库存


支付
    

出餐



<!-- 	
	stockResponse = append(stockResponse, &pb.Item{
		Id:      "1",
		Name:    "瑞幸咖啡生椰拿铁",
		PriceId: "9.9",
		Quantity: 10,
	})
	stockResponse = append(stockResponse, &pb.Item{
		Id:      "2",
		Name:    "瑞幸咖啡冰镇杨梅瑞纳冰",
		PriceId: "14.9",
		Quantity: 5,
	})
	stockResponse = append(stockResponse, &pb.Item{
		Id:      "3",
		Name:    "瑞幸咖啡橙C美式",
		PriceId: "9.9",
		Quantity: 8,
	})
	 -->