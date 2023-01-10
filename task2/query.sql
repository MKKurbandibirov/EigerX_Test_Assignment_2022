select name as dept_name, (select count(id) from employee e where e.dept_id = department.id) as empl_count from department
order by empl_count desc, dept_name;

