select distinct orders.customer_id,customers.customer_name
from orders
join customers on customers.customer_id = orders.customer_id
where 
orders.customer_id not in (select customer_id from orders where product_name = 'C') 
and orders.customer_id in (select customer_id from orders where product_name = 'A')
and orders.customer_id in (select customer_id from orders where product_name = 'B')
group by orders.customer_id,customers.customer_name
order by customers.customer_name


