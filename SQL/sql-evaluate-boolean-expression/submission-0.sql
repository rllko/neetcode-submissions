select
e.*, case 
  when e.operator = '>' then vl.value > vr.value
  when e.operator = '=' then vl.value = vr.value
  when e.operator = '<' then vl.value < vr.value end as value
from expressions e
left join variables vl
on e.left_operand = vl.name
left join variables vr
on e.right_operand = vr.name