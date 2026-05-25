select u.name,COALESCE(SUM(r.distance),0) as travelled_distance
from rides r
right join users u on u.id = r.user_id 
group by u.name,r.user_id
order by travelled_distance desc

  