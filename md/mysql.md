## DB基础-MySQL安装  
+ DMBMS分类   
    1. 基于共享文件系统的DMBMS(Access)
    2. 基于C-S架构的DMBMS(MySQL、Oracle、SqlServer)  
+ 手动管理mysql启动和关闭  
    1. net stop "name"(mysql的名字)  
    2. net start "name"  
+ 连接与退出mysql  
    + cmd连接  
        1. mysql [-h localhost -P 3306] -u root -p //输入后回车即可  
        2. exit  
    + mysql自带的客户端    
---  
+ mysql常见的命令  
    + 参看当前所有的数据库
        `show databases;`
    + 打开指定的库  
        `use 库名;`
    + 查看当前库的所有表   
        `show tables;`  
    + 查看其他库的所有表  
        `show tables from 库名;`
    + 创建表  
    ```   
        create table 表名(  
            列名 列类型，
            列名 列类型，
                。。。
        )  
    ``` 
    + 查看表结构  
        `desc 表名;`  
    + 服务器的版本  
        1. 登录到mysql的服务端  
            `select version(); `
        2. 在cmd界面  
            `mysql --version / mysql --v`  
---  
+ MySQL的语法规范  
    1. 不区分大小写  
    2. 每条命令用分号结尾或者\g  
    3. 每条命令可以根据需要进行缩进或换行  
    4. 注释  
        1. 单行注释：# 或 --空格  
        2. /*  */    

+ 基础查询  
    + select 查询列表 from 表名;
        1. select last_name from employees;
        2. select last_name,salary from employees;
        3. select * from employees;  
        4. select 100('join');//查询常量  int型或字符型   使用``分开 关键字和字段名  
        5. select 100%99  //表达式
        6. select VERSION() //函数
        7. select 100%98 as 结果; selcet 100 结果; //起别名,起的别名中有特殊符号时要加引号。   
        8. select distinct ... from ...  去重  
        9. mysql中的加号：  
            1. select 90 + 10  
            2. select "90" + 10 尝试转化字符型为数字型；失败的化字符型为0  
            3. select null + 10  结果一定为null  
        10. 字符型的拼接  
            1. select CONCAT('a','b')
            2. select ifnull(commission_pct,0) from employees;
            3. select concat(first_name,'+',ifnull(commission_pct,0)) from employees;  
        11. 条件查询  
            + 语法：   
                ```  
                select   
                    查询列表  
                from      
                    表名  
                where  
                    筛选条件   
                1. 条件运算符  
                    < > = <>//不等于  >= <=  
                2.  逻辑运算符  
                    && || ！  
                    and or not  
                3. 模糊查询  
                    like 查询员工中保存字符a的员工信息,一般与通配符一起使用：  
                        % 任意多个字符，包括0; '%a'
                        _ 任意单个字符  
                        ESCAPE 定义转义字符  
                        select   
	                        last_name,email,salary
                        from   
                            employees
                        where  
	                        last_name like "_$_%" escape "$";
                    between and    
                    in 判断某字段的值是否属于in后的列表中
                        select      
                            last_name,job_id
                        from   
                            employees
                        where  
                            job_id in ('IT_PROT','AD_VP')
                    is null     
                        select   
                            last_name,commission_pct
                        from   
                            employees
                        where  
                            commission_pct is null;# = 不支持 null的比较     
                            commission_pct is not null //只能判断null
                            commission_pct <=> null 可以判断其他的等于
                example:    
                    select     
                        email,salary
                    from   
                        employees
                    where  
                        salary >= 100 and salary <= 120>; //使用between and  salary between 100 and 120  
            ```  
        12.   排序查询  
            ```  
                语法：
                    select 查询列表  
                    from 表  
                    [where 查询条件]  
                    order by 查询列表 [asc|desc] ;//asc是升序，desc是降序
                example：  
                    select * from employees order by salary desc;     

                    select   
                        length(first_name) as 长度,first_name
                    from   
                        employees
                    order by 长度;   
                //选择工资不在8000到17000的员工的姓名和工资，按工资降序排列
                select   
                    last_name,salary
                from   
                    employees
                where   
                    salary not between 8000 and 17000
                order by salary desc;
            ```  
        14. 常见函数  
            ```  
                  
            ```

