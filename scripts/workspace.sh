


# oms-v2 目录下

# 步骤 1
go work init

# 步骤 2
go work use ./internal/common/
go work use ./internal/order/
go work use ./internal/stock/
go work use ./internal/payment/
go work use ./internal/kitchen/

# 步骤 3
go work sync



curl -X GET \
  -H "Accept: application/json" \
  http://127.0.0.1:8282/api/v1/customer/12345/orders/67890

