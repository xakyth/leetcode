package main

import "container/heap"

type tweet struct {
	id  int
	idx int
}
type tweetHeap []tweet

func (h tweetHeap) Len() int           { return len(h) }
func (h tweetHeap) Less(i, j int) bool { return h[i].idx < h[j].idx }
func (h tweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h tweetHeap) Peek() tweet        { return h[0] }
func (h *tweetHeap) Push(x any)        { *h = append(*h, x.(tweet)) }
func (h *tweetHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

type Twitter struct {
	userTweetsMap map[int][]tweet
	userFollowMap map[int]map[int]struct{}
	cnt           int
}

func Constructor() Twitter {
	userTweetsMap := map[int][]tweet{}
	userFollowMap := map[int]map[int]struct{}{}
	cnt := 0
	return Twitter{userTweetsMap: userTweetsMap, userFollowMap: userFollowMap, cnt: cnt}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userTweetsMap[userId] = append(this.userTweetsMap[userId], tweet{id: tweetId, idx: this.cnt})
	this.cnt++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := tweetHeap{}
	follows, initialized := this.userFollowMap[userId]
	if initialized {
		for followeeId := range follows {
			tweets := this.userTweetsMap[followeeId]
			for i, cnt := len(tweets)-1, 0; i >= 0 && cnt < 10; i, cnt = i-1, cnt-1 {
				if h.Len() >= 10 && tweets[i].idx < h.Peek().idx {
					break
				}
				heap.Push(&h, tweets[i])
			}
		}
	}
	tweets := this.userTweetsMap[userId]
	for i, cnt := len(tweets)-1, 0; i >= 0 && cnt < 10; i, cnt = i-1, cnt-1 {
		if h.Len() >= 10 && tweets[i].idx < h.Peek().idx {
			break
		}
		heap.Push(&h, tweets[i])
	}
	for h.Len() > 10 {
		heap.Pop(&h)
	}
	feed := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		feed[i] = (heap.Pop(&h).(tweet)).id
	}
	return feed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	_, initialized := this.userFollowMap[followerId]
	if !initialized {
		this.userFollowMap[followerId] = map[int]struct{}{}
	}
	this.userFollowMap[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.userFollowMap[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
