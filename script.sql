SELECT 
    product_name,
    list_price,
    store_name
FROM products
JOIN stocks USING(product_id)
LEFT JOIN stores ON stores.store_id = stocks.store_id