
# oms-v2/internal/common 目录下
go mod init github.com/sjxiang/gorder-v2/common


# oms-v2/internal/order 目录下
go mod init github.com/sjxiang/gorder-v2/order
go mod edit -replace github.com/sjxiang/gorder-v2/common=../common

# oms-v2/internal/stock 目录下
go mod init github.com/sjxiang/gorder-v2/stock
go mod edit -replace github.com/sjxiang/gorder-v2/common=../common


# oms-v2/internal/payment 目录下
go mod init github.com/sjxiang/gorder-v2/payment
go mod edit -replace github.com/sjxiang/gorder-v2/common=../common


# oms-v2/internal/kitchen 目录下
go mod init github.com/sjxiang/gorder-v2/kitchen
go mod edit -replace github.com/sjxiang/gorder-v2/common=../common

