1. Goals
2. Classes
3. Projects
	1. Postgresql,mongodb,redis

The goals of the data class is to become proficient in interfacing with three key types of databases commonly found around the world. The first is an SQL database, the second is an object database, and the third is a Key:Value database. An sql database uses the Structured Query Language or SQL to interface with globs of data written to disk. These databases commonly write thier data to the disk in binary in order to minimize the amount of write syscalls neccisary; this can commonly lead to databases that have over 30% faster write speeds then writing to the disk normally. The Structured Query Language has a heavy emphsis on normalization or the practice of writing identical data to the disk in as little places as possible, instead relying on JOINs and KEY's to link data together.
