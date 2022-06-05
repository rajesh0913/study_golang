package main

//安装redisgo mod(interface)
/*
	key-val(CRUD)

	1.添加key-val:				[set]->set key1 hello
	2.获取key-val:			   	[get]->get key1	<-"hello"
	3.切换数据库:			   	[select]-select [0-15]
	4.查看当前数据库：		   	 [dbsize]
	5.清空当前/所有数据库：		 [flushdb flushall]
	6.删除del：					[del]->del key
	7.setex(set with expire): 	键秒值，equal with : set key value;expire key seconds;
	8.meset:					[mset]->mset key1 value1...
	9.mget:						[mget]->mget key1 key2...
*/

/*
	hash(CRUD)
	1.hset/hmset
	2.hget/hgetall/hmget
	3.hdel
	4.hlen
	5.hexists key field
*/

/*
	list(CRUD)
	1.lpush/rpush
	2.lrange
	3.lpop/rpop
	4.del
	5.lindex->lindex
	6.llen key:key == null <-0
*/

/*
	set(CRUD)
	元素无序，元素值不能重复
	1.sadd
	2.smembers
	3.sismember
	4.srem
*/
