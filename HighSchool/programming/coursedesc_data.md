1. Goals
2. Classes
3. Projects
	1. Postgresql,mongodb,redis

The goals of the data class is to become proficient in interfacing with three key types of databases commonly found around the world. The first is an SQL database, the second is an object database, and the third is a Key:Value database. An sql database uses the Structured Query Language or SQL to interface with globs of data written to disk. These databases commonly write their data to the disk in binary in order to minimize the amount of write syscalls neccisary; this can commonly lead to databases that have over 30% faster write speeds then writing to the disk normally. The Structured Query Language has a heavy emphasis on normalization or the practice of writing identical data to the disk in as little places as possible, instead relying on JOINs and KEY's to link data together. The second type of database is the object database. This type of database values extremely high write and read speeds, and being easy to use and interface with. The database that was interfaced with during the course was MongoDB. The final type of database is the KEY:VALUE database, these are databases like Redis and Cassandra. The one we will use in this course is Redis. These databases emphasise speed and the specific format of their databases.

The class I took to gain knowledge about databases was the Data class by Skilstak Coding Arts done on Tuesdays. So far we have completed interfacing with a Postgres database for sql, and a mongodb database for interfacing with an object database. For a KV database, we will use Redis or Cassandra.

The projects I have done and completed in this class are the following: Setting up a fake bank database to retain practice with interfacing with a Postgresql-9.6 database, and interacting with a MongoDB database in the low level programming language C. 
