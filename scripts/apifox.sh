
# 测试脚本

curl --location --request POST 'http://127.0.0.1:8282/api/v1/customer/20002/orders' \
--header 'User-Agent: Apifox/1.0.0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customer_id": "20002",
    "items": [
        {
            "id": "瑞幸咖啡",
            "quantity": 10
        },
        {
            "id": "库迪咖啡",
            "quantity": 20
        }
    ]
}'


curl --location --request GET 'http://127.0.0.1:8282/api/v1/customer/20002/orders/1747131735' \
--header 'User-Agent: Apifox/1.0.0'