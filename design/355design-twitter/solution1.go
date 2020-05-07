package _355design

import (
	"sort"
	"strconv"
)

type Twitter1 struct {
	globalId int
	follower map[int][]int
	checkFollowed map[string]bool
	//这里twitter的value存globalId
	twitter map[int][]int
	findTwitter []int
}


/** Initialize your data structure here. */
func Constructor1() Twitter1 {
	return Twitter1{
		globalId: 0,
		follower: make(map[int][]int),
		checkFollowed: make(map[string]bool),
		twitter: make(map[int][]int),
		findTwitter: make([]int,0),
	}
}


/** Compose a new tweet. */
func (this *Twitter1) PostTweet(userId int, tweetId int)  {
	this.twitter[userId]=append(this.twitter[userId],this.globalId)
	this.findTwitter=append(this.findTwitter,tweetId)
	this.globalId++
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter1) GetNewsFeed(userId int) []int {
	//可以维护一个大小为10的堆,但这里不这么写了(考虑到可能没有10个推特)
	var globalId []int
	cpy:=make([]int,len(this.follower[userId]))
	copy(cpy,this.follower[userId])
	key:=strconv.Itoa(userId)+" "+strconv.Itoa(userId)
	//因为测试用例里有自己关注自己这种东西，所以加上这一句
	if !this.checkFollowed[key]{
		cpy=append(cpy,userId)
	}
	for _,v:=range cpy{
		globalId=append(globalId,this.twitter[v]...)
	}
	sort.Ints(globalId)
	var res []int
	for i:=len(globalId)-1;i>=0&&i>=len(globalId)-10;i--{
		res=append(res,this.findTwitter[globalId[i]])
	}
	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter1) Follow(followerId int, followeeId int)  {
	key:=strconv.Itoa(followerId)+" "+strconv.Itoa(followeeId)
	if !this.checkFollowed[key]{
		//关注关系更新
		this.checkFollowed[key]=true
		this.follower[followerId]=append(this.follower[followerId],followeeId)
	}

}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter1) Unfollow(followerId int, followeeId int)  {
	key:=strconv.Itoa(followerId)+" "+strconv.Itoa(followeeId)
	if this.checkFollowed[key]{
		this.checkFollowed[key]=false
		for i:=0;i<len(this.follower[followerId]);i++{
			if this.follower[followerId][i]==followeeId{
				this.follower[followerId]=append(this.follower[followerId][:i],this.follower[followerId][i+1:]...)
			}
		}
	}
}



