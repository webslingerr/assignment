select
    ARRAY_AGG(product_name),
    ARRAY_AGG(list_price),
    ARRAY_AGG(quantity),
    store_name
from stocks
left join stores on stores.store_id = stocks.store_id
left join products on products.product_id = stocks.product_id
GROUP BY store_name