#!/bin/bash
# $1 project path
# $2 kssj like 2020-03-10T20:47:51
# $3 jssj like 2020-03-10T20:47:51
cd $1 && svn log -r {$2}:{$3} -v

# svn log -r {2020-03-10T20:47:51}:{2020-03-10T20:49:51} -v
# svn log -r {2020-03-10}:{2020-03-11} -v
