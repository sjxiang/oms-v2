


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


