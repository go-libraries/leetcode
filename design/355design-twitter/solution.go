package _355design

import (
	"sort"
)

type Twitter struct {
	Twitters map[int][]int
	Follows  map[int][]int
	Sub      map[int][]int
	Feed     map[int][]int
	AllTw    map[int]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Twitters: make(map[int][]int),
		Follows:  make(map[int][]int),
		Feed:     make(map[int][]int),
		Sub:      make(map[int][]int),
		AllTw:    make(map[int]int, 0),
	}
}

func (tw *Twitter) initUser(userId int) {
	if _, ok := tw.Twitters[userId]; !ok {
		tw.Twitters[userId] = make([]int, 0)
	}

	if _, ok := tw.Follows[userId]; !ok {
		tw.Follows[userId] = []int{userId}
	}

	if _, ok := tw.Sub[userId]; !ok {
		tw.Sub[userId] = []int{userId}
	}

	if _, ok := tw.Feed[userId]; !ok {
		tw.Feed[userId] = []int{}
	}
}

/** Compose a new tweet. */
func (tw *Twitter) PostTweet(userId int, tweetId int) {
	tw.initUser(userId)
	tw.Twitters[userId] = append(tw.Twitters[userId], tweetId)
	tw.AllTw[tweetId] = len(tw.AllTw)

	for _, feedUserId := range tw.Sub[userId] {
		//fmt.Println(userId, feedUserId, tweetId)

		tw.Feed[feedUserId] = append([]int{tweetId}, tw.Feed[feedUserId][0:]...)
		if len(tw.Feed[feedUserId]) > 10 {
			tw.Feed[feedUserId] = tw.Feed[feedUserId][0:10]
		}
		//tw.changeFeed(feedUserId)
	}
	// tw.Feed[userId] = append([]int{tweetId}, tw.Feed[userId]...)
	//	if len(tw.Feed[userId]) > 10 {
	//		tw.Feed[userId] = tw.Feed[userId][0:10]
	//	}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (tw *Twitter) GetNewsFeed(userId int) []int {
	if v, ok := tw.Feed[userId]; ok {
		return v
	}
	return nil
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (tw *Twitter) Follow(followerId int, followeeId int) {
	tw.initUser(followerId)
	tw.initUser(followeeId)
	if followerId == followeeId {
		return
	}
	exists := false
	for _, id := range tw.Follows[followerId] {
		if id == followeeId {
			exists = true
			break
		}
	}
	if !exists {
		tw.Follows[followerId] = append(tw.Follows[followerId], followeeId)
	}
	if len(tw.Twitters[followeeId]) > 0 && !exists {
		tw.addUserFeed(followerId, followeeId)
	}

	exists = false
	for _, id := range tw.Sub[followeeId] {
		if id == followerId {
			exists = true
			break
		}
	}
	if !exists {
		tw.Sub[followeeId] = append(tw.Sub[followeeId], followerId)
	}

	//tw.changeFeed(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	tw.initUser(followerId)
	tw.initUser(followeeId)
	if followerId == followeeId {
		return
	}

	follows := make([]int, 0)
	for _, v := range tw.Follows[followerId] {
		if v == followeeId {
			continue
		}
		follows = append(follows, v)
	}
	tw.Follows[followerId] = follows

	subs := make([]int, 0)
	for _, v := range tw.Sub[followeeId] {
		if v == followerId {
			continue
		}
		subs = append(subs, v)
	}
	tw.Sub[followeeId] = subs
	if len(tw.Twitters[followeeId]) > 0 {
		tw.resetUserFeed(followerId)
	}
}

func (tw *Twitter) addUserFeed(userId int, SubId int) {
	temp := make([]int,0)
	for i := 0; i < len(tw.Twitters[SubId]); i++ {
		fid := tw.Twitters[SubId][i]
		temp = append(temp, tw.AllTw[fid])
	}
	lf  := len(tw.Feed[userId])
	for j := 0; j < lf; j++ {
		id := tw.Feed[userId][j]
		temp = append(temp, tw.AllTw[id])
	}

	sort.Ints(temp)
	l := len(temp)
	if l > 10 {
		temp = temp[l-10:]
	}

	tw.Feed[userId] = []int{}
	for i := 0; i < len(temp); i++ {
		for id, idPos := range tw.AllTw {
			if temp[i] == idPos {
				tw.Feed[userId] = append([]int{id}, tw.Feed[userId]...)
				break
			}
		}
	}
}

func (tw *Twitter) resetUserFeed(userId int) {
	//clear
	tw.Feed[userId] = []int{}
	temp := make([]int, 0)

	for _, uid := range tw.Follows[userId] {
		for i := 0; i < len(tw.Twitters[uid]); i++ {
			fid := tw.Twitters[uid][i]
			temp = append(temp, tw.AllTw[fid])
		}
	}

	sort.Ints(temp)
	l := len(temp)
	if l > 10 {
		temp = temp[l-10:]
	}

	for i := 0; i < len(temp); i++ {
		for id, idPos := range tw.AllTw {
			if temp[i] == idPos {
				tw.Feed[userId] = append([]int{id}, tw.Feed[userId]...)
				break
			}
		}
	}

	//if len(tw.Feed[userId]) > 10 {
	//	tw.Feed[userId] = tw.Feed[userId][0:10]
	//}
}
