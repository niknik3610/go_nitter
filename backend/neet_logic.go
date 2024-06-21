package main;

type neet struct {
    Creator *nitterUser;
    Contents string;
    Likes uint32;
    Replies []*neetReply;
}

type neetReply struct {
    Creator *nitterUser;
    Contents string;
    Likes uint32;
}

type nitterUser struct {
    UserName string;
    Neets []*neet;
    Followers []*nitterUser;
    Following []*nitterUser;
    LikedNeets []*neet;
}

type NitterState struct{
    users []nitterUser;
}

func (state NitterState) FetchNeetFeed(userName string) []*neet{
    user := nitterUser{
        UserName: "Test User",
        Neets: make([]*neet, 0),
        Followers: make([]*nitterUser, 0),
        Following: make([]*nitterUser, 0),
        LikedNeets: make([]*neet, 0),
    };

    neetOne := neet{
        Creator: &user,
        Contents: "Hello World",
        Likes: 0,
        Replies: make([]*neetReply, 0),
    };

    neetTwo := neet{
        Creator: &user,
        Contents: "Bye World",
        Likes: 0,
        Replies: make([]*neetReply, 0),
    };

    return []*neet{&neetOne, &neetTwo};
}
